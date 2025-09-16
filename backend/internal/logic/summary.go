package logic

import (
	"context"
	"time"
	"volunteer-team/backend/internal/infrastructure/global"
	"volunteer-team/backend/internal/infrastructure/types"
	"volunteer-team/backend/internal/repo"
)

type ISummaryLogic interface {
	CreateSummary(ctx context.Context, userID int64, req types.CreateSummaryReq) (string, error)
}
type SummaryLogic struct {
	summaryRepo *repo.SummaryRepo
	orderRepo   *repo.OrderRepo
}

func NewSummaryLogic() *SummaryLogic {
	return &SummaryLogic{
		summaryRepo: repo.NewSummaryRepo(),
		orderRepo:   repo.NewOrderRepo(),
	}
}

func (sl *SummaryLogic) CreateSummary(ctx context.Context, userID int64, req types.CreateSummaryReq) (string, error) {
	err := sl.summaryRepo.CreateSummary(ctx, userID, req)
	if err != nil {
		global.Log.Error(err)
		return "", DEFAULT_ERROR
	}
	//异步更新订单状态（独立 ctx，防止父 ctx 被取消）
	//这里是为了实现写了修机总结就表示这一单完成了
	go func(orderID int) {
		defer func() {
			if p := recover(); p != nil {
				global.Log.Error("UpdateOrderState panic:", p)
			}
		}()

		// 新建一个与请求无关的 context，最长 3s
		c, cancel := context.WithTimeout(context.Background(), 3*time.Second)
		defer cancel()

		if err := sl.orderRepo.UpdateOrderState(c, orderID); err != nil {
			global.Log.Error("UpdateOrderState failed:", err)
			// TODO: 这里可以落一张“补偿任务表”，或者发 MQ,防止修机总结写失败，而订单完成，这样就永远写不了修机总结了
			return
		}
	}(req.OrderID)
	return "创建修机总结成功", nil
}

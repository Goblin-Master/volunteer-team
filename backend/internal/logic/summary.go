package logic

import (
	"volunteer-team/backend/internal/infrastructure/global"
	"volunteer-team/backend/internal/infrastructure/types"
	"volunteer-team/backend/internal/repo"
)

type ISummaryLogic interface {
	CreateSummary(userID int64, req types.CreateSummaryReq) (string, error)
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

func (sl *SummaryLogic) CreateSummary(userID int64, req types.CreateSummaryReq) (string, error) {
	err := sl.summaryRepo.CreateSummary(userID, req)
	if err != nil {
		global.Log.Error(err)
		return "", DEFAULT_ERROR
	}
	//异步更新订单状态,考虑消息队列进行维护
	//这里是为了实现写了修机总结就表示这一单完成了
	go func(id int) {
		e := sl.orderRepo.UpdateOrderState(id)
		if e != nil {
			global.Log.Error(e)
			return
		}
	}(req.OrderID)
	return "创建修机总结成功", nil
}

package logic

import (
	"context"
	"errors"
	"strconv"
	"time"
	"volunteer-team/backend/internal/infrastructure/global"
	"volunteer-team/backend/internal/infrastructure/types"
	"volunteer-team/backend/internal/infrastructure/utils/snowflake"
	"volunteer-team/backend/internal/repo"
)

type ISummaryLogic interface {
	CreateSummary(ctx context.Context, userID int64, req types.CreateSummaryReq) (string, error)
	GetSummaryList(ctx context.Context) (types.SummaryListResp, error)
	GetSummaryDetail(ctx context.Context, req types.SummaryDetailReq) (types.SummaryDetailResp, error)
	UpdateSummary(ctx context.Context, req types.UpdateSummaryReq) (string, error)
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

var _ ISummaryLogic = (*SummaryLogic)(nil)

func (sl *SummaryLogic) CreateSummary(ctx context.Context, userID int64, req types.CreateSummaryReq) (string, error) {
	//先校验id的合法性，避免浪费时间查询
	orderID, err := strconv.ParseInt(req.OrderID, 10, 64)
	if err != nil {
		return "", PARAMS_TYPE_ERROR
	}

	err = sl.summaryRepo.CreateSummary(ctx, userID, orderID, snowflake.GetIntID(global.Node), req)
	if err != nil {
		global.Log.Error(err)
		return "", DEFAULT_ERROR
	}
	//异步更新订单状态（独立 ctx，防止父 ctx 被取消）
	//这里是为了实现写了修机总结就表示这一单完成了
	go func(orderID int64) {
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
	}(orderID)
	return "创建修机总结成功", nil
}

func (sl *SummaryLogic) GetSummaryList(ctx context.Context) (types.SummaryListResp, error) {
	var resp types.SummaryListResp
	// 内部人员：查全部修机总结
	list, err := sl.summaryRepo.GetSummaryList(ctx)
	if err != nil {
		return resp, DEFAULT_ERROR
	}
	var summaries []types.SummaryItem
	for _, v := range list {
		summary := types.SummaryItem{
			OrderID:            v.OrderID,
			SummaryID:          v.SummaryID,
			ProblemDescription: v.ProblemDescription,
			Utime:              v.Utime,
			ProblemType:        v.ProblemType,
			RepairSummary:      v.RepairSummary,
			ReceiverName:       v.ReceiverName,
		}
		summaries = append(summaries, summary)
	}
	resp.Summaries = summaries
	return resp, nil
}

func (sl *SummaryLogic) GetSummaryDetail(ctx context.Context, req types.SummaryDetailReq) (types.SummaryDetailResp, error) {
	var resp types.SummaryDetailResp
	summaryID, err := strconv.ParseInt(req.SummaryID, 10, 64)
	if err != nil {
		return resp, PARAMS_TYPE_ERROR
	}
	data, err := sl.summaryRepo.GetSummaryDetail(ctx, summaryID)
	if err != nil {
		if errors.Is(err, repo.SUMMARY_NOT_EXIST) {
			return resp, SUMMARY_NOT_EXIST
		}
		global.Log.Error(err)
		return resp, DEFAULT_ERROR
	}
	//组装数据
	resp.RepairSummary = data.RepairSummary
	resp.ReceiverName = data.ReceiverName
	resp.ProblemDescription = data.ProblemDescription
	resp.ProblemType = data.ProblemType
	resp.Utime = data.Utime
	resp.OrderID = data.OrderID
	resp.SummaryID = data.SummaryID
	return resp, nil
}

func (sl *SummaryLogic) UpdateSummary(ctx context.Context, req types.UpdateSummaryReq) (string, error) {
	summaryID, err := strconv.ParseInt(req.SummaryID, 10, 64)
	if err != nil {
		return "", PARAMS_TYPE_ERROR
	}
	err = sl.summaryRepo.UpdateSummary(ctx, summaryID, req)
	if err != nil {
		if errors.Is(err, repo.SUMMARY_NOT_EXIST) {
			return "", SUMMARY_NOT_EXIST
		}
		global.Log.Error(err)
		return "", DEFAULT_ERROR
	}
	return "更新修机总结成功", nil
}

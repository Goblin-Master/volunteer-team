package repo

import (
	"context"
	"errors"
	"volunteer-team/backend/internal/infrastructure/global"
	"volunteer-team/backend/internal/infrastructure/model"
	"volunteer-team/backend/internal/infrastructure/types"
	"volunteer-team/backend/internal/repo/dto"
)

type ISummaryRepo interface {
	CreateSummary(ctx context.Context, userID, orderID, summaryID int64, req types.CreateSummaryReq) error
	GetSummaryList(ctx context.Context) ([]model.Summary, error)
	GetSummaryDetail(ctx context.Context, summaryID int64) (model.Summary, error)
	UpdateSummary(ctx context.Context, summaryID int64, req types.UpdateSummaryReq) error
}
type SummaryRepo struct {
	summaryDto *dto.SummaryDto
}

func NewSummaryRepo() *SummaryRepo {
	return &SummaryRepo{
		summaryDto: dto.NewSummaryDto(),
	}
}

var _ ISummaryRepo = (*SummaryRepo)(nil)

func (sr *SummaryRepo) CreateSummary(ctx context.Context, userID, orderID, summaryID int64, req types.CreateSummaryReq) error {
	summary := model.Summary{
		SummaryID:          summaryID,
		UserID:             userID,
		OrderID:            orderID,
		RepairSummary:      req.RepairSummary,
		ReceiverName:       req.ReceiverName,
		ProblemDescription: req.ProblemDescription,
		ProblemType:        req.ProblemType,
	}
	return sr.summaryDto.AddSummary(ctx, summary)
}

func (sr *SummaryRepo) GetSummaryList(ctx context.Context) ([]model.Summary, error) {
	return sr.summaryDto.GetSummaryList(ctx)
}

func (sr *SummaryRepo) GetSummaryDetail(ctx context.Context, summaryID int64) (model.Summary, error) {
	summary, err := sr.summaryDto.GetSummaryDetailByID(ctx, summaryID)
	if err != nil {
		if errors.Is(err, dto.SUMMARY_NOT_EXIST) {
			return summary, SUMMARY_NOT_EXIST
		}
		global.Log.Error(err)
		return summary, DEFAULT_ERROR
	}
	return summary, nil
}

func (sr *SummaryRepo) UpdateSummary(ctx context.Context, summaryID int64, req types.UpdateSummaryReq) error {
	var summary = model.Summary{
		ProblemType:        req.ProblemType,
		ProblemDescription: req.ProblemDescription,
		RepairSummary:      req.RepairSummary,
		ReceiverName:       req.ReceiverName,
	}
	err := sr.summaryDto.UpdateSummaryByID(ctx, summaryID, summary)
	if err != nil {
		if errors.Is(err, dto.SUMMARY_NOT_EXIST) {
			return SUMMARY_NOT_EXIST
		}
		global.Log.Error(err)
		return DEFAULT_ERROR
	}
	return nil
}

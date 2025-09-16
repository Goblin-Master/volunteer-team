package repo

import (
	"context"
	"volunteer-team/backend/internal/infrastructure/model"
	"volunteer-team/backend/internal/infrastructure/types"
	"volunteer-team/backend/internal/repo/dto"
)

type ISummaryRepo interface {
	CreateSummary(ctx context.Context, userID int64, req types.CreateSummaryReq) error
}
type SummaryRepo struct {
	summaryDto *dto.SummaryDto
}

func NewSummaryRepo() *SummaryRepo {
	return &SummaryRepo{
		summaryDto: dto.NewSummaryDto(),
	}
}

func (sr *SummaryRepo) CreateSummary(ctx context.Context, userID int64, req types.CreateSummaryReq) error {
	summary := model.Summary{
		UserID:             userID,
		RepairSummary:      req.RepairSummary,
		ReceiverName:       req.ReceiverName,
		ProblemDescription: req.ProblemDescription,
		ProblemType:        req.ProblemType,
		OrderID:            req.OrderID,
	}
	return sr.summaryDto.AddSummary(ctx, summary)
}

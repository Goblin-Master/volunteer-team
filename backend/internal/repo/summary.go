package repo

import "volunteer-team/backend/internal/repo/dto"

type ISummaryRepo interface{}
type SummaryRepo struct {
	summaryDto *dto.SummaryDto
}

func NewSummaryRepo() *SummaryRepo {
	return &SummaryRepo{
		summaryDto: dto.NewSummaryDto(),
	}
}

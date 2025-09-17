package dto

import (
	"context"
	"time"
	"volunteer-team/backend/internal/infrastructure/global"
	"volunteer-team/backend/internal/infrastructure/model"
)

type SummaryDto struct{}

func NewSummaryDto() *SummaryDto {
	return &SummaryDto{}
}
func (sd *SummaryDto) AddSummary(ctx context.Context, summary model.Summary) error {
	summary.Ctime = time.Now().UnixMilli()
	summary.Utime = time.Now().UnixMilli()
	return global.DB.WithContext(ctx).Create(&summary).Error
}

func (sd *SummaryDto) GetSummaryList(ctx context.Context) ([]model.Summary, error) {
	var summary []model.Summary
	err := global.DB.WithContext(ctx).Find(&summary).Order("utime DESC").Error
	if err != nil {
		global.Log.Error(err)
		return summary, DEFAULT_ERROR
	}
	return summary, nil
}

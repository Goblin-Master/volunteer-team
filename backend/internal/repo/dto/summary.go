package dto

import (
	"time"
	"volunteer-team/backend/internal/infrastructure/global"
	"volunteer-team/backend/internal/infrastructure/model"
)

type SummaryDto struct{}

func NewSummaryDto() *SummaryDto {
	return &SummaryDto{}
}
func (sd *SummaryDto) AddSummary(summary model.Summary) error {
	summary.Ctime = time.Now().UnixMilli()
	summary.Utime = time.Now().UnixMilli()
	return global.DB.Create(&summary).Error
}

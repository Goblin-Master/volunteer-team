package dto

import (
	"context"
	"errors"
	"time"
	"volunteer-team/backend/internal/infrastructure/global"
	"volunteer-team/backend/internal/infrastructure/model"

	"gorm.io/gorm"
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
		return summary, ErrDefault
	}
	return summary, nil
}

func (sd *SummaryDto) GetSummaryDetailByID(ctx context.Context, summaryID int64) (model.Summary, error) {
	var summary model.Summary
	err := global.DB.WithContext(ctx).Model(&model.Summary{}).Where("summary_id = ?", summaryID).Take(&summary).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return summary, ErrSummaryNotExist
		}
		global.Log.Error(err)
		return summary, ErrDefault
	}
	return summary, nil
}

func (sd *SummaryDto) UpdateSummaryByID(ctx context.Context, summaryID int64, data model.Summary) error {
	data.Utime = time.Now().UnixMilli() //更新修改时间
	result := global.DB.WithContext(ctx).Model(&model.Summary{}).Where("summary_id = ?", summaryID).Updates(data)
	if result.Error != nil {
		global.Log.Error(result.Error)
		return ErrDefault
	}
	if result.RowsAffected == 0 {
		return ErrSummaryNotExist // 没有匹配行
	}
	return nil
}

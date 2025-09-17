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
		return summary, DEFAULT_ERROR
	}
	return summary, nil
}

func (sd *SummaryDto) GetSummaryDetailByID(ctx context.Context, id int) (model.Summary, error) {
	var summary model.Summary
	err := global.DB.WithContext(ctx).Model(&model.Summary{}).Where("id = ?", id).Take(&summary).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return summary, SUMMARY_NOT_EXIST
		}
		global.Log.Error(err)
		return summary, DEFAULT_ERROR
	}
	return summary, nil
}

func (sd *SummaryDto) UpdateSummaryByID(ctx context.Context, id int, data model.Summary) error {
	data.ID = 0                         // 禁止结构体里的 ID 参与任何条件/更新
	data.Utime = time.Now().UnixMilli() //更新修改时间
	result := global.DB.WithContext(ctx).Model(&model.Summary{}).Where("id = ?", id).Updates(data)
	if result.Error != nil {
		global.Log.Error(result.Error)
		return DEFAULT_ERROR
	}
	if result.RowsAffected == 0 {
		return SUMMARY_NOT_EXIST // 没有匹配行
	}
	return nil
}

package repositories

import (
	"github.com/crowdint/gopher-spree-api/domain"
)

type AdjustmentRepository struct {
	DbRepository
}

func NewAdjustmentRepository() *AdjustmentRepository {
	return &AdjustmentRepository{
		DbRepository{Spree_db},
	}
}

func (this *AdjustmentRepository) AllByAdjustable(adjustableId int64, adjustableType string) []domain.Adjustment {
	adjustments := []domain.Adjustment{}
	this.All(&adjustments, nil, "adjustable_id = ? AND adjustable_type = ?", adjustableId, adjustableType)
	return adjustments
}

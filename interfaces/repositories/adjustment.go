package repositories

import (
	"github.com/crowdint/gopher-spree-api/domain/json"
)

type AdjustmentRepository struct {
	DbRepository
}

func NewAdjustmentRepository() *AdjustmentRepository {
	return &AdjustmentRepository{
		DbRepository{Spree_db},
	}
}

func (this *AdjustmentRepository) AllByAdjustable(adjustableId int64, adjustableType string) []json.Adjustment {
	adjustments := []json.Adjustment{}
	this.All(&adjustments, nil, "adjustable_id = ? AND adjustable_type = ?", adjustableId, adjustableType)
	return adjustments
}

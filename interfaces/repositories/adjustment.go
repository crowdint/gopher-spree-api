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

func (this *AdjustmentRepository) AllByAdjustable(adjustable domain.Adjustable) []domain.Adjustment {
	adjustments := []domain.Adjustment{}

	this.All(&adjustments, nil, "adjustable_id = ? AND adjustable_type = ?", adjustable.AdjustableId(), adjustable.SpreeClass())

	for _, adjustment := range adjustments {
		adjustment.Adjustable = adjustable
	}

	return adjustments
}

package repositories

import (
	"github.com/crowdint/gopher-spree-api/domain/models"
)

type AssetRepository struct {
	DbRepository
}

func NewAssetRepo() *AssetRepository {
	return &AssetRepository{
		DbRepository{dbHandler: Spree_db},
	}
}

func (this *AssetRepository) FindByViewableIds(viewableIds []int64) ([]*models.Asset, error) {
	var assets []*models.Asset

	if len(viewableIds) == 0 {
		return assets, nil
	}

	query := this.dbHandler.
		Where("viewable_id in (?)", viewableIds).
		Find(&assets)

	return assets, query.Error
}

func (this *AssetRepository) AllImagesByVariantId(variantId int64) ([]*models.Asset, error) {
	modelImages := []*models.Asset{}
	err := this.All(&modelImages, map[string]interface{}{
		"order": "position ASC",
	}, "type IN ('Spree::Image') AND viewable_id = ? AND viewable_type = ?", variantId, "Spree::Variant")
	return modelImages, err
}

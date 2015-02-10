package repositories

import "github.com/crowdint/gopher-spree-api/domain/models"

type AssetRepo DbRepo

func NewAssetRepo() *AssetRepo {
	return &AssetRepo{
		dbHandler: Spree_db,
	}
}

func (this *AssetRepo) FindByViewableIds(viewableIds []int64) ([]*models.Asset, error) {
	var assets []*models.Asset

	if len(viewableIds) == 0 {
		return assets, nil
	}

	query := this.dbHandler.
		Where("viewable_id in (?)", viewableIds).
		Find(&assets)

	return assets, query.Error
}

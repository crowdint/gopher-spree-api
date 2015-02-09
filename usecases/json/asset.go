package json

import (
	"github.com/crowdint/gopher-spree-api/domain/json"
	"github.com/crowdint/gopher-spree-api/domain/models"
	"github.com/crowdint/gopher-spree-api/interfaces/repositories"
	"github.com/jinzhu/copier"
)

type AssetInteractor struct {
	Repo *repositories.AssetRepo
}

func NewAssetInteractor() *AssetInteractor {
	return &AssetInteractor{
		Repo: repositories.NewAssetRepo(),
	}
}

type JsonAssetsMap map[int64][]*json.Asset

func (this *AssetInteractor) GetJsonAssetsMap(viewableIds []int64) (JsonAssetsMap, error) {

	assets, err := this.Repo.FindByViewableIds(viewableIds)
	if err != nil {
		return JsonAssetsMap{}, err
	}

	assetsJson := this.modelsToJsonAssetsMap(assets)

	return assetsJson, nil
}

func (this *AssetInteractor) modelsToJsonAssetsMap(assetSlice []*models.Asset) JsonAssetsMap {
	jsonAssetsMap := JsonAssetsMap{}

	for _, asset := range assetSlice {
		assetJson := json.Asset{}
		copier.Copy(&assetJson, asset)

		if _, exists := jsonAssetsMap[asset.ViewableId]; !exists {
			jsonAssetsMap[asset.ViewableId] = []*json.Asset{}
		}

		jsonAssetsMap[asset.ViewableId] = append(jsonAssetsMap[asset.ViewableId], &assetJson)

	}

	return jsonAssetsMap
}

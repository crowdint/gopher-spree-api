package json

import (
	"strings"

	"github.com/crowdint/gopher-spree-api/configs"
	"github.com/crowdint/gopher-spree-api/domain"
	"github.com/crowdint/gopher-spree-api/interfaces/repositories"
	"github.com/crowdint/gopher-spree-api/utils"
)

type AssetInteractor struct {
	Repository *repositories.AssetRepository
}

func NewAssetInteractor() *AssetInteractor {
	return &AssetInteractor{
		Repository: repositories.NewAssetRepo(),
	}
}

type JsonAssetsMap map[int64][]*domain.Asset

func (this *AssetInteractor) GetJsonAssetsMap(viewableIds []int64) (JsonAssetsMap, error) {

	assets, err := this.Repository.FindByViewableIds(viewableIds)
	if err != nil {
		utils.LogrusError(utils.FuncName(), err)
		return JsonAssetsMap{}, err
	}

	assetsJson := this.modelsToJsonAssetsMap(assets)

	return assetsJson, nil
}

func (this *AssetInteractor) modelsToJsonAssetsMap(assetSlice []*domain.AssetModel) JsonAssetsMap {
	jsonAssetsMap := JsonAssetsMap{}

	for _, asset := range assetSlice {
		assetJson := this.toJson(asset)

		if _, exists := jsonAssetsMap[asset.ViewableId]; !exists {
			jsonAssetsMap[asset.ViewableId] = []*domain.Asset{}
		}

		jsonAssetsMap[asset.ViewableId] = append(jsonAssetsMap[asset.ViewableId], assetJson)

	}

	return jsonAssetsMap
}

func (this *AssetInteractor) toJsonAssets(modelAssets []*domain.AssetModel) []*domain.Asset {
	jsonAssets := []*domain.Asset{}
	for _, modelAsset := range modelAssets {
		jsonAssets = append(jsonAssets, this.toJson(modelAsset))
	}
	return jsonAssets
}

func (this *AssetInteractor) toJson(asset *domain.AssetModel) *domain.Asset {
	assetJson := domain.Asset{
		"id":                      asset.Id,
		"viewable_id":             asset.ViewableId,
		"viewable_type":           asset.ViewableType,
		"attachment_width":        asset.AttachmentWidth,
		"attachment_height":       asset.AttachmentHeight,
		"attachment_file_size":    asset.AttachmentFileSize,
		"position":                asset.Position,
		"attachment_content_type": asset.AttachmentContentType,
		"attachment_file_name":    asset.AttachmentFileName,
		"attachment_updated_at":   asset.AttachmentUpdatedAt,
		"type":                    asset.Type,
		"alt":                     asset.Alt,
	}

	defaultStyles := configs.Get(configs.SPREE_DEFAULT_STYLES)
	stylesSlice := strings.Split(defaultStyles, ",")

	for _, style := range stylesSlice {
		assetUrl := style + "_url"
		assetJson[assetUrl] = asset.AssetUrl(style)
	}

	return &assetJson
}

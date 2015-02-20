package json

import (
	"strings"

	"github.com/crowdint/gopher-spree-api/configs"
	"github.com/crowdint/gopher-spree-api/domain/json"
	"github.com/crowdint/gopher-spree-api/domain/models"
	"github.com/crowdint/gopher-spree-api/interfaces/repositories"
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
		assetJson := this.toJson(asset)

		if _, exists := jsonAssetsMap[asset.ViewableId]; !exists {
			jsonAssetsMap[asset.ViewableId] = []*json.Asset{}
		}

		jsonAssetsMap[asset.ViewableId] = append(jsonAssetsMap[asset.ViewableId], assetJson)

	}

	return jsonAssetsMap
}

func (this *AssetInteractor) toJson(asset *models.Asset) *json.Asset {
	assetJson := json.Asset{
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

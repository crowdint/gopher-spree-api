package json

import (
	"github.com/crowdint/gopher-spree-api/domain/models"
	"testing"
)

func TestAssetInteractor_modelsToJsonAssetsMap(t *testing.T) {
	assetSlice := []*models.Asset{
		&models.Asset{
			Id:                 1,
			ViewableId:         1,
			AttachmentFileName: "asset1.jpg",
		},
		&models.Asset{
			Id:                 2,
			ViewableId:         3,
			AttachmentFileName: "asset2.jpg",
		},
	}

	assetInteractor := NewAssetInteractor()

	jsonAssetMap := assetInteractor.modelsToJsonAssetsMap(assetSlice)

	a1 := *jsonAssetMap[1][0]

	if a1["mini_url"] != "/spree/products/1/mini/asset1.jpg" {
		t.Error("Wrong assignment of values")
	}
}

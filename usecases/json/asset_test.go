package json

import (
	"github.com/crowdint/gopher-spree-api/domain/models"
	"testing"
)

func TestAssetInteractor_modelsToJsonAssetsMap(t *testing.T) {
  assetSlice := []*models.Asset{
    &models.Asset{
      Id:        1,
      ViewableId: 1,
      AttachmentFileName: "asset1.jpg",
    },
    &models.Asset{
      Id:        2,
      ViewableId: 3,
      AttachmentFileName: "asset2.jpg",
    },
  }

  assetInteractor := NewAssetInteractor()

  jsonAssetMap := assetInteractor.modelsToJsonAssetsMap(assetSlice)

  a1 := jsonAssetMap[1][0]
  a2 := jsonAssetMap[3][0]

  if a1.Id != 1 || a1.MiniUrl != "/mini/asset1.jpg" {
    t.Error("Wrong assignment of values", a1.MiniUrl, a1.Id)
  }

  if a2.Id != 2 || a2.MiniUrl != "/mini/asset2.jpg" {
    t.Error("Wrong assignment of values", a2.MiniUrl, a2.Id)
  }

}


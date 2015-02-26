package json

import (
	"testing"

	"github.com/crowdint/gopher-spree-api/domain/json"
)

func TestAssetInteractor_modelsToJsonAssetsMap(t *testing.T) {
	assetSlice := []*json.AssetModel{
		&json.AssetModel{
			Id:                 1,
			ViewableId:         1,
			AttachmentFileName: "asset1.jpg",
		},
		&json.AssetModel{
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

func TestAssetInteractor_toJsonAssets(t *testing.T) {
	modelAssets := []*json.AssetModel{
		&json.AssetModel{
			Id:                 1,
			ViewableId:         1,
			AttachmentFileName: "asset1.jpg",
		},
		&json.AssetModel{
			Id:                 2,
			ViewableId:         3,
			AttachmentFileName: "asset2.jpg",
		},
	}

	assetInteractor := NewAssetInteractor()
	jsonAsstes := assetInteractor.toJsonAssets(modelAssets)

	if len(jsonAsstes) == 0 {
		t.Error("Json Assets len should be 2, but was 0")
	}

	for i, jsonAsset := range jsonAsstes {
		modelAsset := modelAssets[i]
		if (*jsonAsset)["id"].(int) != modelAsset.Id {
			t.Errorf("Json Asset Id should be %d, but was %d", modelAsset.Id, (*jsonAsset)["id"])
		}

		if (*jsonAsset)["viewable_id"].(int64) != modelAsset.ViewableId {
			t.Errorf("Json Asset Id should be %d, but was %d", modelAsset.ViewableId, (*jsonAsset)["viewable_id"])
		}

		if (*jsonAsset)["attachment_file_name"].(string) != modelAsset.AttachmentFileName {
			t.Errorf("Json Asset Id should be %d, but was %d", modelAsset.AttachmentFileName, (*jsonAsset)["attachment_file_name"])
		}
	}
}

package json

import "testing"

func TestAsset_assetUrl(t *testing.T) {
	asset := AssetModel{
		Id:                 1,
		AttachmentFileName: "image.jpg",
	}

	assetUrl := asset.AssetUrl("waka")

	if assetUrl != "/spree/products/1/waka/image.jpg" {
		t.Error("Wrong asset url string", assetUrl)
	}
}

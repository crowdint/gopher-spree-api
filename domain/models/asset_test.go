package models

import "testing"

func TestAsset_assetUrl(t *testing.T) {
  asset := Asset{
    AttachmentFileName: "image.jpg",
  }

  assetUrl := asset.assetUrl("waka")

  if assetUrl != "/waka/image.jpg" {
    t.Error("Wrong asset url string")
  }
}

func TestAsset_MiniUrl(t *testing.T) {
  asset := Asset{
    AttachmentFileName: "image.jpg",
  }

  assetUrl := asset.MiniUrl()

  if assetUrl != "/mini/image.jpg" {
    t.Error("Wrong asset url string: ", assetUrl)
  }
}


func TestAsset_SmallUrl(t *testing.T) {
  asset := Asset{
    AttachmentFileName: "image.jpg",
  }

  assetUrl := asset.SmallUrl()

  if assetUrl != "/small/image.jpg" {
    t.Error("Wrong small url: ", assetUrl)
  }
}


func TestAsset_ProductUrl(t *testing.T) {
  asset := Asset{
    AttachmentFileName: "image.jpg",
  }

  assetUrl := asset.ProductUrl()

  if assetUrl != "/product/image.jpg" {
    t.Error("Wrong product url: ", assetUrl)
  }
}


func TestAsset_LargeUrl(t *testing.T) {
  asset := Asset{
    AttachmentFileName: "image.jpg",
  }

  assetUrl := asset.LargeUrl()

  if assetUrl != "/large/image.jpg" {
    t.Error("Wrong large url: ", assetUrl)
  }
}

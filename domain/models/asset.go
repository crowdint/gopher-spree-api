package models

import (
	"github.com/crowdint/gopher-spree-api/configs"
	"regexp"
	"strconv"
	"time"
)

type Asset struct {
	Id                    int
	ViewableId            int64
	ViewableType          string
	AttachmentWidth       int64
	AttachmentHeight      int64
	AttachmentFileSize    int64
	Position              int64
	AttachmentContentType string
	AttachmentFileName    string
	Type                  string
	AttachmentUpdatedAt   time.Time
	Alt                   string
	CreatedAt             time.Time
	UpdatedAt             time.Time
}

func (this Asset) TableName() string {
	return "spree_assets"
}

func (this Asset) AssetConfig(style string) (string, map[string]string) {
	assetUrl := configs.Get(configs.SPREE_ASSET_PATH)

	assetConfig := map[string]string{
		":host":     configs.Get(configs.SPREE_ASSET_HOST),
		":id":       strconv.Itoa(this.Id),
		":style":    style,
		":filename": this.AttachmentFileName,
	}
	return assetUrl, assetConfig
}

func (this Asset) AssetUrl(style string) string {
	url, config := this.AssetConfig(style)
	regExp := regexp.MustCompile(":([a-zA-Z]*)")

	assetUrl := regExp.ReplaceAllStringFunc(url, func(m string) string {
		return config[m]
	})

	return assetUrl
}

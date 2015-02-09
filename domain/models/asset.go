package models

import (
	"github.com/crowdint/gopher-spree-api/configs"
	"strconv"
	"strings"
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

func (this Asset) AssetUrl(style string) string {
	assetHost := configs.Get(configs.SPREE_ASSET_HOST)
	assetUrl := configs.Get(configs.SPREE_ASSET_PATH)
	//Default :host/spree/products/:asset_id/:style/:filename
	assetUrl = strings.Replace(assetUrl, ":host", assetHost, 1)
	assetUrl = strings.Replace(assetUrl, ":asset_id", strconv.Itoa(this.Id), 1)
	assetUrl = strings.Replace(assetUrl, ":style", style, 1)
	assetUrl = strings.Replace(assetUrl, ":filename", this.AttachmentFileName, 1)

	return assetUrl
}

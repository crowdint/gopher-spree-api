package models

import "time"

type Asset struct {
	Id                    int64
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

func (this Asset) assetUrl(style string) string {
  return "/" + style + "/" + this.AttachmentFileName
}

func (this Asset) MiniUrl() string {
  return this.assetUrl("mini")
}

func (this Asset) SmallUrl() string {
  return this.assetUrl("small")
}

func (this Asset) ProductUrl() string {
  return this.assetUrl("product")
}

func (this *Asset) LargeUrl() string {
  return this.assetUrl("large")
}

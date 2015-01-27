package models

import "time"

type Asset struct {
  ID                    int64
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

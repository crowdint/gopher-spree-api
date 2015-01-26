package domain

import (
	"time"
)

type Asset struct {
	ID                    int64     `json:"id"`
	ViewableID            int64     `json:"viewable_id"`
	ViewableType          string    `json:"viewable_type"`
	AttachmentWidth       int64     `json:"attachment_width"`
	AttachmentHeight      int64     `json:"attachment_height"`
	AttachmentFileSize    int64     `json:"-"`
	Position              int64     `json:"position"`
	AttachmentContentType string    `json:"attachment_content_type"`
	AttachmentFileName    string    `json:"attachment_file_name"`
	AttachmentUpdatedAt   time.Time `json:"attachment_updated_at"`
	Alt                   string    `json:"alt"`
	DeletedAt             time.Time `json:"-"`
	UpdatedAt             time.Time `json:"-"`
}

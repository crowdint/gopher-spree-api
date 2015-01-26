package json

import (
	"testing"
	"time"
)

func TestAssetStructure(t *testing.T) {
	expected := `{"id":1,"viewable_id":2,"viewable_type"` +
		`:"type","attachment_width":12,"attachment_height":10` +
		`,"position":10,"attachment_content_type":"text",` +
		`"attachment_file_name":"file.txt","attachment_updated_at":` +
		`"2015-01-01T00:00:00Z","alt":"image..."}`

	someTime := time.Date(2015, 1, 1, 0, 0, 0, 0, time.UTC)

	asset := Asset{
		ID:                    1,
		ViewableID:            2,
		ViewableType:          "type",
		AttachmentWidth:       12,
		AttachmentHeight:      10,
		AttachmentFileSize:    56,
		Position:              10,
		AttachmentContentType: "text",
		AttachmentFileName:    "file.txt",
		AttachmentUpdatedAt:   someTime,
		Alt:                   "image...",
		DeletedAt:             someTime,
		UpdatedAt:             someTime,
	}

	AssertEqualJson(t, asset, expected)
}

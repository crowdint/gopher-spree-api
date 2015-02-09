package json

import (
	"testing"
	"time"
)

func TestAssetStructure(t *testing.T) {
	expected := `{"id":24,"position":1,"attachment_content_type":"image/jpeg",` +
		`"attachment_file_name":"image.jpeg","type":"Spree::Image",` +
		`"attachment_updated_at":"2015-01-01T00:00:00Z","attachment_width":360,` +
		`"attachment_height":360,"alt":"Alt","viewable_type":"Spree::Variant","viewable_id":3,` +
		`"mini_url":"mini/image.jpeg","small_url":"small/image.jpeg",` +
		`"product_url":"product/image.jpeg","large_url":"large/image.jpeg"}`

	someTime := time.Date(2015, 1, 1, 0, 0, 0, 0, time.UTC)

	asset := Asset{
		Id:                    24,
		Position:              1,
		AttachmentContentType: "image/jpeg",
		AttachmentFileName:    "image.jpeg",
		Type:                  "Spree::Image",
		AttachmentUpdatedAt:   someTime,
		AttachmentWidth:       360,
		AttachmentHeight:      360,
		Alt:                   "Alt",
		ViewableType:          "Spree::Variant",
		ViewableID:            3,
		MiniUrl:               "mini/image.jpeg",
		SmallUrl:              "small/image.jpeg",
		ProductUrl:            "product/image.jpeg",
		LargeUrl:              "large/image.jpeg",
	}

	AssertEqualJson(t, asset, expected)
}

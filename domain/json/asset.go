package json

import "time"

type Asset struct {
  ID                      int64     `json:"id"`
  Position                int64     `json:"position"`
  AttachmentContentType   string    `json:"attachment_content_type"`
  AttachmentFileName      string    `json:"attachment_file_name"`
  Type                    string    `json:"type"`
  AttachmentUpdatedAt     time.Time `json:"attachment_updated_at"`
  AttachmentWidth         int64     `json:"attachment_width"`
  AttachmentHeight        int64     `json:"attachment_height"`
  Alt                     string    `json:"alt"`
  ViewableType            string    `json:"viewable_type"`
  ViewableId              int64     `json:"viewable_id"`
  MiniUrl                 string    `json:"mini_url"`
  SmallUrl                string    `json:"small_url"`
  ProductUrl              string    `json:"product_url"`
  LargeUrl                string    `json:"large_url"`
}

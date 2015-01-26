package models

import "time"

type OptionType struct {
  ID            int64
  Name          string
  Presentation  string
  Position      int64
  CreatedAt     time.Time
  UpdatedAt     time.Time
}

func (this OptionType) TableName() string {
  return "spree_option_types"
}

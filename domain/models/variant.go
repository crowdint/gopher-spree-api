package models

import "time"

type Variant struct {
  ID              int64
  Sku             string
  Weight          float64
  Height          float64
  Width           float64
  Depth           float64
  DeletedAt       time.Time
  IsMaster        bool
  ProductId       int64
  CostPrice       float64
  Position        int64
  CostCurrency    string
  TrackInventory  bool
  TaxCategoryId   int64
  UpdatedAt       time.Time
  StockItemsCount int64
}

func (this Variant) TableName() string {
  return "spree_variants"
}

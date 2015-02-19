package json

type StockLocation struct {
	Id     int64
	Active bool
}

func (this StockLocation) TableName() string {
	return "spree_stock_locations"
}

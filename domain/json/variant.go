package json

type Variant struct {
	Id              int64         `json:"id"`
	Name            string        `json:"name"`
	Sku             string        `json:"sku"`
	Price           string        `json:"price"`
	Weight          float64       `json:"weight,string"`
	Height          float64       `json:"height,string"`
	Width           float64       `json:"width,string"`
	Depth           float64       `json:"depth,string"`
	IsMaster        bool          `json:"is_master"`
	Slug            string        `json:"slug"`
	Description     string        `json:"description"`
	TrackInventory  bool          `json:"track_inventory"`
	CostPrice       string        `json:"cost_price"`
	DisplayPrice    string        `json:"display_price"`
	OptionsText     string        `json:"options_text"`
	InStock         bool          `json:"in_stock"`
	IsBackorderable bool          `json:"is_backorderable"`
	TotalOnHand     int64         `json:"total_on_hand"`
	IsDestroyed     bool          `json:"is_destroyed"`
	OptionValues    []*OptionValue `json:"option_values"`
	Images          []*Asset      `json:"images"`
}

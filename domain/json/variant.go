package json

type Variant struct {
	ID              int64         `json:"id"`
	Name            string        `json:"name"`
	Sku             string        `json:"sku"`
	Price           string        `json:"price"`
	Weight          string        `json:"weight"`
	Height          string        `json:"height"`
	Width           string        `json:"width"`
	Depth           string        `json:"depth"`
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
	OptionValues    []OptionValue `json:"option_values"`
	Images          []*Asset      `json:"images"`
}

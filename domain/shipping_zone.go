package domain

type ShippingZone struct {
	Id          int64  `json:"id"`
	Description string `json:"description"`
	Name        string `json:"name"`
}

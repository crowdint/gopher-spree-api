package domain

type Country struct {
	Id      int64  `json:"id"`
	Iso     string `json:"iso"`
	IsoName string `json:"iso_name"`
	Iso3    string `json:"iso3"`
	Name    string `json:"name"`
	Numcode int64  `json:"numcode"`
}

func (this Country) TableName() string {
	return "spree_countries"
}

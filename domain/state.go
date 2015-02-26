package domain

type State struct {
	Id        int64  `json:"id"`
	Abbr      string `json:"abbr"`
	CountryId int64  `json:"country_id"`
	Name      string `json:"name"`
}

func (this State) TableName() string {
	return "spree_states"
}

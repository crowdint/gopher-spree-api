package models

type Classification struct {
	ProductId int64
	TaxonId   int64
	Id        int64
	Position  int64
}

func (this Classification) TableName() string {
	return "spree_products_taxons"
}

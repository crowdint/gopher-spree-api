package models

import "time"

type TaxonomyRepository interface {
	FindById(id int64)
}

type Taxonomy struct {
	Id        int64
	Name      string
	Position  int64
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (this Taxonomy) TableName() string {
	return "spree_taxonomies"
}

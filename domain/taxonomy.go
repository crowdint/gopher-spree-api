package domain

import "time"

type Taxonomy struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
	Root Taxon  `json:"root"`

	Position  int64     `json:"-"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

func (this Taxonomy) TableName() string {
	return "spree_taxonomies"
}

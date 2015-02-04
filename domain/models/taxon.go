package models

import "time"

type Taxon struct {
	Id                     int64
	ParentId               int64
	Position               int64
	Name                   string
	Permalink              string
	TaxonomyId             int64
	Lft                    int64
	Rgt                    int64
	IconFileName           string
	IconContentType        string
	IconFileSize           int64
	IconUpdatedAt          time.Time
	Description            string
	CreatedAt              time.Time
	UpdatedAt              time.Time
	MetaTitle              string
	MetaDescription        string
	MetaKeywords           string
	Depth                  int64
	PrettyName             string
	ClassificationPosition int64
	ProductId              int64
}

func (this Taxon) TableName() string {
	return "spree_taxons"
}

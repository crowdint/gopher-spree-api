package domain

import "time"

type Taxon struct {
	Id         int64    `json:"id"`
	Name       string   `json:"name"`
	PrettyName string   `json:"pretty_name"`
	Permalink  string   `json:"permalink"`
	ParentId   int64    `json:"parent_id"`
	TaxonomyId int64    `json:"taxonomy_id"`
	Lft        int64    `json:"-"`
	Rgt        int64    `json:"-"`
	Depth      int64    `json:"-"`
	Taxons     []*Taxon `json:"taxons"`

	Position               int64     `json:"-"`
	IconFileName           string    `json:"-"`
	IconContentType        string    `json:"-"`
	IconFileSize           int64     `json:"-"`
	IconUpdatedAt          time.Time `json:"-"`
	Description            string    `json:"-"`
	CreatedAt              time.Time `json:"-"`
	UpdatedAt              time.Time `json:"-"`
	MetaTitle              string    `json:"-"`
	MetaDescription        string    `json:"-"`
	MetaKeywords           string    `json:"-"`
	ClassificationPosition int64     `json:"-"`
	ProductId              int64     `json:"-"`
}

func (this Taxon) TableName() string {
	return "spree_taxons"
}

func (this Taxon) AfterFind() (err error) {
	this.PrettyName = this.Name
	return
}

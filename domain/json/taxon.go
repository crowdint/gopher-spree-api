package domain

type Taxon struct {
	ID         int64   `json:"id"`
	ParentID   int64   `json:"parent_id"`
	TaxonomyID int64   `json:"taxonomy_id"`
	Position   int64   `json:"position"`
	Name       string  `json:"name"`
	PrettyName string  `json:"pretty_name"`
	Permalink  string  `json:"permalink"`
	Taxons     []Taxon `json:"taxons"`
}

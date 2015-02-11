package json

type Taxon struct {
	Id         int64   `json:"id"`
	Name       string  `json:"name"`
	PrettyName string  `json:"pretty_name"`
	Permalink  string  `json:"permalink"`
	ParentID   int64   `json:"parent_id"`
	TaxonomyID int64   `json:"taxonomy_id"`
	Taxons     []Taxon `json:"taxons"`
}

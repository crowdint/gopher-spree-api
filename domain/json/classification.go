package json

type Classification struct {
	TaxonId  int64 `json:"taxon_id"`
	Position int64 `json:"position"`
	Taxon    Taxon `json:"taxon"`
}

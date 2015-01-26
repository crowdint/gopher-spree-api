package json

import "testing"

func TestClassificationStructure(t *testing.T) {
	expected := `{"taxon_id":1,"position":20,"taxon":` +
		`{"id":1,"parent_id":1,"taxonomy_id":10,"position"` +
		`:20,"name":"taxon","pretty_name":"a name","permalink"` +
		`:"http://someurl.com","taxons":null}}`

	classification := Classification{
		TaxonId:  1,
		Position: 20,
		Taxon: Taxon{
			ID:         1,
			ParentID:   1,
			TaxonomyID: 10,
			Position:   20,
			Name:       "taxon",
			PrettyName: "a name",
			Permalink:  "http://someurl.com",
		},
	}
	AssertEqualJson(t, classification, expected)
}

package json

import "testing"

func TestClassificationStructure(t *testing.T) {
	expected := `{"taxon_id":1,"position":20,"taxon":` +
		`{"id":1,"name":"taxon","pretty_name":"a name",` +
		`"permalink":"http://someurl.com","parent_id":1,"taxonomy_id"` +
		`:10,"taxons":null}}`

	classification := Classification{
		TaxonId:  1,
		Position: 20,
		Taxon: Taxon{
			ID:         1,
			Name:       "taxon",
			PrettyName: "a name",
			Permalink:  "http://someurl.com",
			ParentID:   1,
			TaxonomyID: 10,
		},
	}
	AssertEqualJson(t, classification, expected)
}

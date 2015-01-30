package json

import "testing"

func TestTaxonStructure(t *testing.T) {
	expected := `{"id":1,"name":"taxon1","pretty_name":"a name","permalink":"http://someurl.com",` +
		`"parent_id":1,"taxonomy_id":10,"taxons":[{"id":2,"name":"taxon2","pretty_name":"a name",` +
		`"permalink":"http://someurl.com","parent_id":2,"taxonomy_id":10,"taxons":null}]}`

	taxon := Taxon{
		ID:         1,
		Name:       "taxon1",
		PrettyName: "a name",
		Permalink:  "http://someurl.com",
		ParentID:   1,
		TaxonomyID: 10,
		Taxons: []Taxon{
			{
				ID:         2,
				Name:       "taxon2",
				PrettyName: "a name",
				Permalink:  "http://someurl.com",
				ParentID:   2,
				TaxonomyID: 10,
			},
		},
	}

	AssertEqualJson(t, taxon, expected)
}

package json

import "testing"

func TestTaxonStructure(t *testing.T) {
	expected := `{"id":1,"parent_id":1,"taxonomy_id":10,"position"` +
		`:20,"name":"taxon1","pretty_name":"a name","permalink":` +
		`"http://someurl.com","taxons":[{"id":2,"parent_id":2,` +
		`"taxonomy_id":10,"position":20,"name":"taxon2","pretty_name":` +
		`"a name","permalink":"http://someurl.com","taxons":null},{"id":0,` +
		`"parent_id":0,"taxonomy_id":0,"position":0,"name":"",` +
		`"pretty_name":"","permalink":"","taxons":null}]}`

	taxon := Taxon{
		ID:         1,
		ParentID:   1,
		TaxonomyID: 10,
		Position:   20,
		Name:       "taxon1",
		PrettyName: "a name",
		Permalink:  "http://someurl.com",
		Taxons: []Taxon{
			{
				ID:         2,
				ParentID:   2,
				TaxonomyID: 10,
				Position:   20,
				Name:       "taxon2",
				PrettyName: "a name",
				Permalink:  "http://someurl.com",
			},
			{},
		},
	}

	AssertEqualJson(t, taxon, expected)
}

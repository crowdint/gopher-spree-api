package json

import "testing"

func TestTaxonStructure(t *testing.T) {
	expected := `{"id":1,"name":"taxon1","pretty_name":"a name","permalink":"http://someurl.com",` +
		`"parent_id":1,"taxonomy_id":10,"taxons":[{"id":2,"name":"taxon2","pretty_name":"a name",` +
		`"permalink":"http://someurl.com","parent_id":2,"taxonomy_id":10,"taxons":null}]}`

	taxon := Taxon{
		Id:         1,
		Name:       "taxon1",
		PrettyName: "a name",
		Permalink:  "http://someurl.com",
		ParentId:   1,
		TaxonomyId: 10,
		Taxons: []Taxon{
			{
				Id:         2,
				Name:       "taxon2",
				PrettyName: "a name",
				Permalink:  "http://someurl.com",
				ParentId:   2,
				TaxonomyId: 10,
			},
		},
	}

	AssertEqualJson(t, taxon, expected)
}

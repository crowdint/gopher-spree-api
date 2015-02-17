package json

import (
	"testing"

	"github.com/crowdint/gopher-spree-api/domain/json"
)

func TestTaxonStateMachine(t *testing.T) {
	root := &json.Taxon{Lft: 1, Rgt: 12, Depth: 0, Taxons: []*json.Taxon{}}
	root2 := &json.Taxon{Lft: 13, Rgt: 22, Depth: 0, Taxons: []*json.Taxon{}}

	taxons := []*json.Taxon{
		root,
		root2,
		&json.Taxon{Lft: 2, Rgt: 3, Depth: 1, Taxons: []*json.Taxon{}},
		&json.Taxon{Lft: 4, Rgt: 5, Depth: 1, Taxons: []*json.Taxon{}},
		&json.Taxon{Lft: 6, Rgt: 11, Depth: 1, Taxons: []*json.Taxon{}},
		&json.Taxon{Lft: 7, Rgt: 8, Depth: 2, Taxons: []*json.Taxon{}},
		&json.Taxon{Lft: 9, Rgt: 10, Depth: 2, Taxons: []*json.Taxon{}},
		&json.Taxon{Lft: 14, Rgt: 15, Depth: 1, Taxons: []*json.Taxon{}},
		&json.Taxon{Lft: 16, Rgt: 17, Depth: 1, Taxons: []*json.Taxon{}},
		&json.Taxon{Lft: 18, Rgt: 19, Depth: 1, Taxons: []*json.Taxon{}},
		&json.Taxon{Lft: 20, Rgt: 21, Depth: 1, Taxons: []*json.Taxon{}},
	}

	toTaxonTree(taxons)

	if len(root.Taxons) != 3 {
		t.Errorf("Root 1 should have 3 children, it has %d", len(root.Taxons))
	}

	if len(root2.Taxons) != 4 {
		t.Errorf("Root 2 should have 4 children, it has %d", len(root2.Taxons))
	}
}

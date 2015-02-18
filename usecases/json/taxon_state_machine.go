package json

import (
	"fmt"

	"github.com/crowdint/gopher-spree-api/domain/json"
)

const (
	LEFT  = "left"
	RIGHT = "right"
)

var (
	root        *json.Taxon
	current     *json.Taxon
	childTaxons = map[int64][]*json.Taxon{}
	toEvaluate  = []*json.Taxon{}
)

type StateFunc func() StateFunc

func initMachine(paramRoot *json.Taxon, taxons []*json.Taxon) {
	root = paramRoot

	current = root

	toEvaluate = taxons

	childTaxons = map[int64][]*json.Taxon{}
}

func walkingDown() StateFunc {
	nextTaxon := findBy(LEFT, current.Lft+1)

	if nextTaxon != nil {
		current = nextTaxon

		if isLeaf(nextTaxon) {
			addToChildMap(nextTaxon)

			return bottomReached
		}
	} else {
		return walkingUp
	}

	return walkingDown
}

func bottomReached() StateFunc {
	nextTaxon := findBy(LEFT, current.Rgt+1)

	if nextTaxon != nil {
		current = nextTaxon

		if isLeaf(current) {
			addToChildMap(nextTaxon)
		}

		return walkingDown
	}

	return walkingUp
}

func walkingUp() StateFunc {
	nextTaxon := findBy(RIGHT, current.Rgt+1)

	if nextTaxon != nil {
		nextTaxon.Taxons = getFromChildMap(nextTaxon.Depth + 1)

		addToChildMap(nextTaxon)

		current = nextTaxon

		return walkingUp
	} else if current.Rgt < root.Rgt {

		current = findBy(LEFT, current.Rgt+1)

		if current != nil {

			if isLeaf(current) {
				addToChildMap(current)

				return bottomReached
			}
			return walkingDown
		}
	}

	return nil
}

func findBy(direction string, number int64) *json.Taxon {
	for _, taxon := range toEvaluate {
		if direction == LEFT {
			if taxon.Lft == number {
				return taxon
			}
		}
		if direction == RIGHT {
			if taxon.Rgt == number {
				return taxon
			}
		}
	}
	return nil
}

func isLeaf(taxon *json.Taxon) bool {
	return (taxon.Rgt - taxon.Lft) == 1
}

func addToChildMap(taxon *json.Taxon) {
	if _, exists := childTaxons[taxon.Depth]; !exists {

		childTaxons[taxon.Depth] = []*json.Taxon{}

	}
	childTaxons[taxon.Depth] = append(childTaxons[taxon.Depth], taxon)
}

func getFromChildMap(depth int64) []*json.Taxon {
	if _, exists := childTaxons[depth]; exists {

		arr := childTaxons[depth]
		childTaxons[depth] = []*json.Taxon{}

		return arr
	}
	return []*json.Taxon{}
}

func getRoots(taxons []*json.Taxon) []*json.Taxon {
	roots := []*json.Taxon{}

	for _, taxon := range taxons {
		if taxon.Depth == 0 {
			roots = append(roots, taxon)
		}
	}

	return roots
}

func printTaxonTree(taxon *json.Taxon) {
	for _, childTaxon := range taxon.Taxons {
		printTaxonTree(childTaxon)
	}
	fmt.Println(taxon)
}

func toTaxonTree(taxons []*json.Taxon) {

	rts := getRoots(taxons)

	for _, r := range rts {
		initMachine(r, taxons)

		for state := walkingDown; state != nil; {

			state = state()

		}
	}

}

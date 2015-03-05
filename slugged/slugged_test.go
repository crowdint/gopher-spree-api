package slugged

import (
	"strings"
	"testing"
)

type FakeSlugged struct{}

func (this *FakeSlugged) SlugCandidates() []interface{} {
	return []interface{}{"Spree Tote", []interface{}{"Spree Tote", "SPR-00015"}}
}

func (this *FakeSlugged) SetSlug(slug string) {

}

func TestGenerateSlugs(t *testing.T) {
	slugs := GenerateSlugs(&FakeSlugged{}, "-")
	if len(slugs) != 3 {
		t.Errorf("Lenght of slugs should be 3, but was %d", len(slugs))
		return
	}

	if slugs[0] != "spree-tote" {
		t.Errorf("%s should be spree-tote", slugs[0])
	}

	if slugs[1] != "spree-tote-spr-00015" {
		t.Errorf("%s should be spree-tote-spr-00015", slugs[1])
	}

	if !strings.Contains(slugs[2], "spree-tote") {
		t.Errorf("%s should contain spree-tote", slugs[2])
	}
}

func TestTransformCandidate(t *testing.T) {
	candidate := transformCandidate("Spree Tote", "&")
	if candidate != "spree&tote" {
		t.Errorf("Candidate should be spree&tote, but was %s", candidate)
	}
}

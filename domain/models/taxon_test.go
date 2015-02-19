package models

import "testing"

func TestTaxon_PrettyName(t *testing.T) {
	taxon := Taxon{
		Id:   1,
		Name: "Brand",
	}

	prettyName := taxon.PrettyName()

	if prettyName != "Brand" {
		t.Error("Wrong default pretty name", prettyName)
	}
}

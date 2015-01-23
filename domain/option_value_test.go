package domain

import "testing"

func TestOptionValueStructure(t *testing.T) {
	expected := `{"id":1,"name":"option val1","presentation":` +
		`"presentation","option_type_name":"option1",` +
		`"option_type_id":1,"option_type_presentation":"ot presentation"}`

	optionValue := OptionValue{
		ID:                     1,
		Name:                   "option val1",
		Presentation:           "presentation",
		OptionTypeName:         "option1",
		OptionTypeID:           1,
		OptionTypePresentation: "ot presentation",
	}

	AssertEqualJson(t, optionValue, expected)

}

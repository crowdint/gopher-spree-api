package json

import "testing"

func TestOptionTypeStructure(t *testing.T) {
	expected := `{"id":1,"name":"option1","presentation":"presentation","position":20}`

	optionType := OptionType{
		Id:           1,
		Name:         "option1",
		Presentation: "presentation",
		Position:     20,
	}
	AssertEqualJson(t, optionType, expected)
}

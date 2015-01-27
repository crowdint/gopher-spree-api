package json

import "testing"

func TestProductPropertyStructure(t *testing.T) {
	expected := `{"id":1,"product_id":1,"property_id":1,"value":"some value","property_name":"some name"}`

	productProperty := ProductProperty{
		ID:           1,
		ProductID:    1,
		PropertyID:   1,
		Value:        "some value",
		PropertyName: "some name",
	}
	AssertEqualJson(t, productProperty, expected)
}

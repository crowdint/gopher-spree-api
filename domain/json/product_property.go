package json

type ProductProperty struct {
	Id           int64  `json:"id"`
	ProductId    int64  `json:"product_id"`
	PropertyId   int64  `json:"property_id"`
	Value        string `json:"value"`
	PropertyName string `json:"property_name"`
}

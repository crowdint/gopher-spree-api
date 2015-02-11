package json

type ProductProperty struct {
	Id           int64  `json:"id"`
	ProductID    int64  `json:"product_id"`
	PropertyID   int64  `json:"property_id"`
	Value        string `json:"value"`
	PropertyName string `json:"property_name"`
}

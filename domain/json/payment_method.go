package json

type PaymentMethod struct {
	Id          int64  `json:"id"`
	Environment string `json:"environment"`
	Name        string `json:"name"`
}

package json

type ProductResponse struct {
	Products    []*Product `json:"products"`
	Count       int        `json:"count"`
	Pages       int        `json:"pages"`
	CurrentPage int        `json:"current_page"`
}

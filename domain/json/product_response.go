package json

type ProductResponse struct {
	Count       int        `json:"count"`
	TotalCount  int        `json:"total_count"`
	CurrentPage int        `json:"current_page"`
	PerPage     int        `json:"per_page"`
	Pages       int        `json:"pages"`
	Products    []*Product `json:"products"`
}

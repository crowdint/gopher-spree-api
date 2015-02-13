package json

type Taxonomy struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
	Root Taxon  `json:"root"`
}

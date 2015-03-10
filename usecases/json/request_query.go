package json

const (
	GRANSAK = iota
	ELASTIC_SEARCH
)

type RequestQuery struct {
	Type   int
	Params []interface{}
	Query  string
}

package domain

type Adjustable interface {
	AdjustableId() int64
	AdjustableCurrency() string
	SpreeClass() string
}

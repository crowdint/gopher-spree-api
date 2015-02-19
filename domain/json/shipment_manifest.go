package json

type ShipmentManifest struct {
	States    map[string]int64 `json:"states"`
	Quantity  int64            `json:"quantity"`
	VariantId int64            `json:"variant_id"`
}

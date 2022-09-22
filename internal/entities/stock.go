package entities

const (
	IN  string = "in"
	OUT string = "out"
)

type Stock struct {
	TypeMovement string `json:"type_movement"`
	Quantity     int    `json:"quantity"`
	Invoice      int    `json:"invoice,omitempty"`
	Date         string `json:"date"`
	SupplierId   int    `json:"supplier_id"`
	ProductId    int    `json:"product_id"`
}

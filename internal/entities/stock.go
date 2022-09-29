package entities

const (
	IN  string = "in"
	OUT string = "out"
)

type Stock struct {
	TypeMovement string `json:"type_movement" db:"type_movement"`
	Quantity     int    `json:"quantity" db:"quantity"`
	Invoice      int    `json:"invoice,omitempty" db:"invoice"`
	Date         string `json:"date" db:"date"`
	SupplierId   int    `json:"supplier_id" db:"supplier_id"`
	ProductId    int    `json:"product_id" db:"product_id"`
	UserId       int    `json:"user_id" db:"user_id"`
}

type StockResult struct {
	Movements       []Stock `json:"movements"`
	CurrentQuantity int     `json:"current_quantity"`
}

type StockRepository interface {
	Add(stock Stock) error
	Remove(movId int) error
	GetStock(productId int) ([]Stock, error)
}

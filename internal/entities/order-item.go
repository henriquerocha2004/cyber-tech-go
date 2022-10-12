package entities

import "errors"

const (
	TypeProduct string = "product"
	TypeService string = "service"
)

type OrderItem struct {
	Id        int      `json:"id,omitempty" db:"id,omitempty"`
	ProductId int      `json:"product_id" db:"product_id"`
	OrderId   int      `json:"order_id" db:"order_id"`
	Type      string   `json:"type" db:"type"`
	Quantity  int      `json:"quantity" db:"quantity"`
	Value     float64  `json:"value" db:"value"`
	Product   *Product `json:"product,omitempty"`
}

func (o *OrderItem) SetTypeItem(itemType string) error {
	switch itemType {
	case TypeProduct, TypeService:
		o.Type = itemType
		return nil
	default:
		return errors.New("invalid order type")
	}
}

func (o *OrderItem) Subtotal() float64 {
	return o.Value * float64(o.Quantity)
}

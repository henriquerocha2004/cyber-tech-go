package entities

type OrderPayment struct {
	Id               int     `json:"id,omitempty" db:"id,omitempty"`
	OrderId          int     `json:"order_id" db:"order_id"`
	Description      string  `json:"description" db:"description"`
	TotalValue       float64 `json:"total_value" db:"total_value"`
	Installments     int     `json:"installments" db:"installments"`
	InstallmentValue float64 `json:"installment_value" db:"installment_value"`
}

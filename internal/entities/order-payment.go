package entities

type Payment struct {
	Description      string  `json:"description"`
	TotalValue       float64 `json:"total_value"`
	Installments     int     `json:"installments"`
	InstallmentValue float64 `json:"installment_value"`
}

package entities

type Equipment struct {
	Id           int    `json:"id,omitempty" db:"id,omitempty"`
	Description  string `json:"description" db:"description"`
	Defect       string `json:"defect" db:"defect"`
	Observations string `json:"observations" db:"observations"`
	OrderId      int    `json:"order_id" db:"order_id"`
}

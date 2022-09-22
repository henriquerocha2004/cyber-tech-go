package entities

type Service struct {
	Id          int     `json:"id,omitempty"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}

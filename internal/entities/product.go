package entities

type Product struct {
	Id              int     `json:"id,omitempty"`
	Name            string  `json:"name"`
	ProductGroup    int     `json:"product_group"`
	MinQuantity     int     `json:"min_quantity"`
	MaxQuantity     int     `json:"max_quantity"`
	CostValue       float64 `json:"cost_value"`
	Value           float64 `json:"value"`
	CurrentQuantity int     `json:"current_quantity"`
}

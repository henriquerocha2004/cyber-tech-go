package entities

type Product struct {
	Id              int     `json:"id,omitempty" db:"id,omitempty"`
	Name            string  `json:"name" db:"name" validate:"required"`
	ProductGroup    int     `json:"product_group" db:"product_group" validate:"required"`
	MinQuantity     int     `json:"min_quantity" db:"min_quantity" validate:"required"`
	MaxQuantity     int     `json:"max_quantity" db:"max_quantity" validate:"required"`
	CostValue       float64 `json:"cost_value" db:"cost_value"`
	Value           float64 `json:"value" db:"value" validate:"required"`
	CurrentQuantity int     `json:"current_quantity"`
}

type ProductRepository interface {
	Create(product Product) error
	Update(product Product) error
	Delete(id int) error
	FindOne(id int) (Product, error)
	FindAll() ([]Product, error)
}

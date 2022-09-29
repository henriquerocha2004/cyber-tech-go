package entities

type Supplier struct {
	Id       int    `json:"id,omitempty" db:"id,omitempty"`
	Name     string `json:"name" db:"name" validate:"required"`
	Document string `json:"document" db:"document"`
	Address  string `json:"address" db:"address"`
	District string `json:"district" db:"district"`
	City     string `json:"city" db:"city"`
	State    string `json:"state" db:"state"`
}

type SupplierRepository interface {
	Create(supplier Supplier) error
	Update(supplier Supplier) error
	Delete(id int) error
	FindOne(id int) (Supplier, error)
	FindAll() ([]Supplier, error)
}

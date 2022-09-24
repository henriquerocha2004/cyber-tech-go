package entities

type ProductCategory struct {
	Description string `json:"description" validate:"required"`
	Id          int    `json:"id,omitempty"`
}

type ProductCategoryRepository interface {
	Create(productCategory ProductCategory) error
	Update(productCategory ProductCategory) error
	Delete(id int) error
	FindOne(id int) (ProductCategory, error)
	FindAll() ([]ProductCategory, error)
}

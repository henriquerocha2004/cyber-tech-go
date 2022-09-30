package entities

type OrderServiceStatus struct {
	Id              int    `json:"id,omitempty" db:"id,omitempty"`
	Description     string `json:"description" db:"description" validate:"required"`
	LaunchFinancial bool   `json:"launch_financial" db:"launch_financial" validate:"required"`
	Color           string `json:"color" db:"color" validate:"required"`
}

type OrderServiceStatusRepository interface {
	Create(status OrderServiceStatus) error
	Update(status OrderServiceStatus) error
	Delete(id int) error
	FindOne(id int) (OrderServiceStatus, error)
	FindAll() ([]OrderServiceStatus, error)
}

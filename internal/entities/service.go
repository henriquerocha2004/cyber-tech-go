package entities

type Service struct {
	Id          int     `json:"id,omitempty"`
	Description string  `json:"description" validate:"required"`
	Price       float64 `json:"price" validate:"required"`
}

type ServiceRepository interface {
	Create(service Service) error
	Update(service Service) error
	Delete(serviceId int) error
	FindOne(serviceId int) (Service, error)
	FindAll() ([]Service, error)
}

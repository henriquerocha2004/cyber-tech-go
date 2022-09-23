package mysql

import (
	"github.com/henriquerocha2004/cyber-tech-go/internal/entities"
	"github.com/jmoiron/sqlx"
)

type ServiceRepository struct {
	mysqlConnection *sqlx.DB
}

func NewServiceRepository(mysqlConn *sqlx.DB) *ServiceRepository {
	return &ServiceRepository{
		mysqlConnection: mysqlConn,
	}
}

func (s *ServiceRepository) Create(service entities.Service) error {
	query := `INSERT INTO services (description, price) VALUES (:description, :price)`
	_, err := s.mysqlConnection.NamedExec(query, service)
	if err != nil {
		return err
	}
	return nil
}

func (s *ServiceRepository) Update(service entities.Service) error {
	query := `UPDATE services SET description = :description, price = :price WHERE id = :id`
	_, err := s.mysqlConnection.NamedExec(query, service)
	if err != nil {
		return err
	}
	return nil
}

func (s *ServiceRepository) Delete(serviceId int) error {
	query := `DELETE FROM services WHERE id = ?`
	s.mysqlConnection.MustExec(query, serviceId)
	return nil
}

func (s *ServiceRepository) FindOne(serviceId int) (entities.Service, error) {
	service := entities.Service{}
	query := `SELECT id, description, price FROM services WHERE id = ?`
	err := s.mysqlConnection.Get(&service, query, serviceId)
	return service, err
}

func (s *ServiceRepository) FindAll() ([]entities.Service, error) {
	var services []entities.Service
	query := `SELECT id, description, price FROM services`
	err := s.mysqlConnection.Select(&services, query)
	return services, err
}

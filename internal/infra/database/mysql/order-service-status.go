package mysql

import (
	"github.com/henriquerocha2004/cyber-tech-go/internal/entities"
	"github.com/jmoiron/sqlx"
)

type OrderServiceStatusRepository struct {
	connection *sqlx.DB
}

func NewOrderServiceStatusRepository(conn *sqlx.DB) *OrderServiceStatusRepository {
	return &OrderServiceStatusRepository{
		connection: conn,
	}
}

func (o *OrderServiceStatusRepository) Create(status entities.OrderServiceStatus) error {
	query := `
		INSERT INTO order_service_status 
			(description, launch_financial, color)
		VALUES
			(:description,:launch_financial,:color)	
	`
	_, err := o.connection.NamedExec(query, status)
	return err
}

func (o *OrderServiceStatusRepository) Update(status entities.OrderServiceStatus) error {
	query := `
		UPDATE order_service_status 
			SET description = :description, launch_financial = :launch_financial, color = :color
		WHERE id = :id
	`
	_, err := o.connection.NamedExec(query, status)
	return err
}

func (o *OrderServiceStatusRepository) Delete(id int) error {
	query := `DELETE FROM order_service_status WHERE id = ?`
	o.connection.MustExec(query, id)
	return nil
}

func (o *OrderServiceStatusRepository) FindOne(id int) (entities.OrderServiceStatus, error) {
	var status entities.OrderServiceStatus
	query := `SELECT id, description, launch_financial, color FROM order_service_status WHERE id = ?`
	err := o.connection.Get(&status, query, id)
	return status, err
}

func (o *OrderServiceStatusRepository) FindAll() ([]entities.OrderServiceStatus, error) {
	var statuses []entities.OrderServiceStatus
	query := `SELECT id, description, launch_financial, color FROM order_service_status`
	err := o.connection.Select(&statuses, query)
	return statuses, err
}

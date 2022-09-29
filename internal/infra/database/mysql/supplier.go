package mysql

import (
	"github.com/henriquerocha2004/cyber-tech-go/internal/entities"
	"github.com/jmoiron/sqlx"
)

type SupplierRepository struct {
	mysqlConnection *sqlx.DB
}

func NewSupplierRepository(conn *sqlx.DB) *SupplierRepository {
	return &SupplierRepository{
		mysqlConnection: conn,
	}
}
func (s *SupplierRepository) Create(supplier entities.Supplier) error {
	query := `
		INSERT INTO suppliers 
			(name, document, address, district, city, state) 
		VALUES
		 	(:name, :document, :address, :district, :city, :state)	
	`
	_, err := s.mysqlConnection.NamedExec(query, supplier)
	return err
}

func (s *SupplierRepository) Update(supplier entities.Supplier) error {
	query := `
		UPDATE suppliers SET name = :name, document = :document, address = :address,
			district = :district, city = :city, state = :state WHERE id = :id
	`
	_, err := s.mysqlConnection.NamedExec(query, supplier)
	return err
}

func (s *SupplierRepository) Delete(id int) error {
	query := `DELETE FROM suppliers WHERE id = ?`
	s.mysqlConnection.MustExec(query, id)
	return nil
}

func (s *SupplierRepository) FindOne(id int) (entities.Supplier, error) {
	var supplier entities.Supplier
	query := `
		SELECT id,name, document, address, district, city, state 
			FROM suppliers WHERE id = ?
	`
	err := s.mysqlConnection.Get(&supplier, query, id)
	return supplier, err
}

func (s *SupplierRepository) FindAll() ([]entities.Supplier, error) {
	var suppliers []entities.Supplier
	query := `SELECT id,name, document, address, district, city, state FROM suppliers`
	err := s.mysqlConnection.Select(&suppliers, query)
	return suppliers, err
}

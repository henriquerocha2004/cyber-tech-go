package mysql

import (
	"github.com/henriquerocha2004/cyber-tech-go/internal/entities"
	"github.com/jmoiron/sqlx"
)

type ProductRepository struct {
	mysqlConnection *sqlx.DB
}

func NewProductRepository(conn *sqlx.DB) *ProductRepository {
	return &ProductRepository{
		mysqlConnection: conn,
	}
}

func (p *ProductRepository) Create(product entities.Product) error {
	query := `
		INSERT INTO products 
			(name, product_group, min_quantity, max_quantity, cost_value, value)
		VALUES
			(:name,:product_group,:min_quantity,:max_quantity,:cost_value,:value)	
	`
	_, err := p.mysqlConnection.NamedExec(query, product)
	return err
}

func (p *ProductRepository) Update(product entities.Product) error {
	query := `
		UPDATE products SET name = :name, product_group = :product_group, min_quantity = :min_quantity, 
			max_quantity = :max_quantity, cost_value = :cost_value, value = :value WHERE id = :id
	`
	_, err := p.mysqlConnection.NamedExec(query, product)
	return err
}

func (p *ProductRepository) Delete(id int) error {
	query := `DELETE FROM products WHERE id = ?`
	p.mysqlConnection.MustExec(query, id)
	return nil
}

func (p *ProductRepository) FindOne(id int) (entities.Product, error) {
	product := entities.Product{}
	query := `SELECT name, product_group, min_quantity, max_quantity, cost_value, value FROM products WHERE id = ?`
	err := p.mysqlConnection.Get(&product, query, id)
	return product, err
}

func (p *ProductRepository) FindAll() ([]entities.Product, error) {
	products := []entities.Product{}
	query := `SELECT name, product_group, min_quantity, max_quantity, cost_value, value FROM products`
	err := p.mysqlConnection.Select(&products, query)
	return products, err
}

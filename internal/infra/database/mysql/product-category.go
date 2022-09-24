package mysql

import (
	"github.com/henriquerocha2004/cyber-tech-go/internal/entities"
	"github.com/jmoiron/sqlx"
)

type ProductCategoryRepository struct {
	mysqlConnection *sqlx.DB
}

func NewProductCategoryRepository(mysqlConn *sqlx.DB) *ProductCategoryRepository {
	return &ProductCategoryRepository{
		mysqlConnection: mysqlConn,
	}
}

func (p *ProductCategoryRepository) Create(productCategory entities.ProductCategory) error {
	query := `INSERT INTO product_categories (description) VALUES (:description)`
	_, err := p.mysqlConnection.NamedExec(query, productCategory)
	return err
}

func (p *ProductCategoryRepository) Update(productCategory entities.ProductCategory) error {
	query := `UPDATE product_categories SET description = :description WHERE id = :id`
	_, err := p.mysqlConnection.NamedExec(query, productCategory)
	return err
}

func (p *ProductCategoryRepository) Delete(id int) error {
	query := `DELETE FROM product_categories WHERE id = ?`
	p.mysqlConnection.MustExec(query, id)
	return nil
}

func (p *ProductCategoryRepository) FindOne(id int) (entities.ProductCategory, error) {
	var category entities.ProductCategory
	query := `SELECT id, description FROM product_categories WHERE id = ?`
	err := p.mysqlConnection.Get(&category, query, id)
	return category, err
}

func (p *ProductCategoryRepository) FindAll() ([]entities.ProductCategory, error) {
	var categories []entities.ProductCategory
	query := `SELECT id, description FROM product_categories`
	err := p.mysqlConnection.Select(&categories, query)
	return categories, err
}

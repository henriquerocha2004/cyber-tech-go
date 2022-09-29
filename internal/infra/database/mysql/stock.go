package mysql

import (
	"github.com/henriquerocha2004/cyber-tech-go/internal/entities"
	"github.com/jmoiron/sqlx"
)

type StockRepository struct {
	mysqlConnection *sqlx.DB
}

func NewStockRepository(connection *sqlx.DB) *StockRepository {
	return &StockRepository{
		mysqlConnection: connection,
	}
}

func (s *StockRepository) Add(stock entities.Stock) error {
	query := `
		INSERT INTO stock 
			(type_movement, quantity, invoice, date, supplier_id, product_id, user_id)
		VALUES
			(:type_movement,:quantity,:invoice,:date,:supplier_id,:product_id,:user_id)
	`
	_, err := s.mysqlConnection.NamedExec(query, stock)
	return err
}

func (s *StockRepository) Remove(movId int) error {
	query := `DELETE FROM stock WHERE id = ?`
	s.mysqlConnection.MustExec(query, movId)
	return nil
}

func (s *StockRepository) GetStock(productId int) ([]entities.Stock, error) {
	var stock []entities.Stock
	query := `SELECT type_movement, quantity, invoice, date, supplier_id, product_id, user_id FROM stock WHERE product_id = ?`
	err := s.mysqlConnection.Select(&stock, query, productId)
	return stock, err
}

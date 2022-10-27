package actions

import (
	"errors"
	"log"
	"time"

	"github.com/henriquerocha2004/cyber-tech-go/internal/entities"
)

type StockActions struct {
	stockRepository entities.StockRepository
}

type StockInput struct {
	TypeMovement string `json:"type_movement" validate:"required"`
	Quantity     int    `json:"quantity" validate:"required"`
	Invoice      string `json:"invoice,omitempty"`
	Date         string `json:"date"`
	SupplierId   int    `json:"supplier_id"`
	ProductId    int    `json:"product_id" validate:"required"`
	UserId       int    `json:"user_id" validate:"required"`
}

type StockOutput struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

func NewStockActions(stockRepo entities.StockRepository) *StockActions {
	return &StockActions{
		stockRepository: stockRepo,
	}
}

func (s *StockActions) Add(stockInput StockInput) StockOutput {
	output := StockOutput{}
	stock := entities.Stock{
		TypeMovement: stockInput.TypeMovement,
		Quantity:     stockInput.Quantity,
		Invoice:      stockInput.Invoice,
		Date:         time.Now().Format("2006-01-02 15:04"),
		SupplierId:   stockInput.SupplierId,
		ProductId:    stockInput.ProductId,
		UserId:       stockInput.UserId,
	}
	log.Println(s.stockRepository)
	err := s.stockRepository.Add(stock)
	if err != nil {
		log.Println(err)
		output.Error = true
		output.Message = `Error in add product in stock`
		return output
	}

	output.Error = true
	output.Message = `Product added in stock successfully`
	return output
}

func (s *StockActions) Delete(stockId int) StockOutput {
	output := StockOutput{}
	err := s.stockRepository.Remove(stockId)
	if err != nil {
		log.Println(err)
		output.Error = true
		output.Message = `Error in delete product in stock`
		return output
	}

	output.Error = false
	output.Message = `Product added in stock successfully`
	return output
}

func (s *StockActions) FindStock(productId int) StockOutput {
	output := StockOutput{}
	stock, err := s.stockRepository.GetStock(productId)
	if err != nil {
		log.Println(err)
		output.Error = true
		output.Message = `Error in delete product in stock`
		return output
	}
	stockResult := s.calcQuantityProductInStock(stock)
	output.Error = false
	output.Data = stockResult
	return output
}

func (s *StockActions) CreateStockByOrderProcess(orderInput ServiceOrderInput) error {
	if !s.validateOrder(&orderInput) {
		return nil
	}

	for _, item := range orderInput.Items {

		if item.Type == entities.TypeService {
			continue
		}

		var stockInput StockInput
		stockInput.TypeMovement = entities.OUT
		stockInput.Quantity = item.Quantity
		stockInput.ProductId = item.ProductId
		stockInput.UserId = 1

		output := s.Add(stockInput)

		if output.Error {
			return errors.New(output.Message)
		}
	}

	return nil
}

func (s *StockActions) validateOrder(orderInput *ServiceOrderInput) bool {
	if !orderInput.Paid {
		return false
	}

	if len(orderInput.Items) < 1 {
		return false
	}

	var productsId []int

	for _, item := range orderInput.Items {
		if item.Type == entities.TypeProduct {
			productsId = append(productsId, item.ProductId)
		}
	}

	return len(productsId) >= 1
}

func (s *StockActions) calcQuantityProductInStock(stockMovements []entities.Stock) entities.StockResult {
	var result entities.StockResult

	for _, movement := range stockMovements {
		if movement.TypeMovement == entities.IN {
			result.CurrentQuantity += movement.Quantity
		}

		if movement.TypeMovement == entities.OUT {
			result.CurrentQuantity -= movement.Quantity
		}
	}

	result.Movements = stockMovements
	return result
}

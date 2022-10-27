package mysql

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/henriquerocha2004/cyber-tech-go/internal/entities"
	"github.com/jmoiron/sqlx"
)

type OrderServiceCommandRepository struct {
	connection *sqlx.DB
}

func NewOrderServiceCommandRepository(conn *sqlx.DB) *OrderServiceCommandRepository {
	return &OrderServiceCommandRepository{
		connection: conn,
	}
}

func (o *OrderServiceCommandRepository) Create(order entities.OrderService) (int, error) {
	order.CreatedAt = time.Now().Format("2006-01-02 15:04")
	order.UpdatedAt = time.Now().Format("2006-01-02 15:04")
	tx := o.connection.MustBegin()

	query := `
		INSERT INTO order_service
		    (number, description, user_id, status_id, created_at, updated_at, paid)
		VALUES 
		    (:number,:description,:user_id,:status_id,:created_at,:updated_at,:paid)
   `
	result, err := tx.NamedExec(query, order)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	idInserted, err := result.LastInsertId()
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	err = o.processItems(order.Items, int(idInserted), tx)
	if err != nil {
		return 0, err
	}

	err = o.processEquipments(order.Equipments, int(idInserted), tx)
	if err != nil {
		return 0, err
	}

	err = o.processOrderPayments(order.Payments, int(idInserted), tx)
	if err != nil {
		return 0, err
	}

	err = tx.Commit()
	if err != nil {
		return 0, err
	}

	return int(idInserted), nil
}

func (o *OrderServiceCommandRepository) Update(order entities.OrderService) error {
	tx := o.connection.MustBegin()
	order.UpdatedAt = time.Now().Format("2006-01-02 15:04")
	query := `
		UPDATE order_service SET 
		    number = :number, description = :description, user_id = :user_id, status_id = :status_id, updated_at = :updated_at, paid = :paid,
			close_date = :close_date 
		WHERE id = :id
	`
	_, err := tx.NamedExec(query, order)
	if err != nil {
		tx.Rollback()
		return err
	}

	err = o.processItems(order.Items, order.Id, tx)
	if err != nil {
		return err
	}

	err = o.processEquipments(order.Equipments, order.Id, tx)
	if err != nil {
		return err
	}

	err = o.processOrderPayments(order.Payments, order.Id, tx)
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return err
}

func (o *OrderServiceCommandRepository) processItems(items []entities.OrderItem, orderId int, tx *sqlx.Tx) error {

	if len(items) < 1 {
		return nil
	}

	for index, item := range items {
		item.OrderId = int(orderId)
		items[index] = item
	}
	err := o.syncItems(items, tx)
	if err != nil {
		tx.Rollback()
		return err
	}
	return nil
}

func (o *OrderServiceCommandRepository) syncItems(items []entities.OrderItem, tx *sqlx.Tx) error {
	queryDelete := `DELETE FROM order_items WHERE order_id = ?`
	tx.MustExec(queryDelete, items[0].OrderId)
	queryInsert := `INSERT INTO order_items (product_id, order_id, type, quantity, value) VALUES (:product_id, :order_id, :type, :quantity, :value)`
	_, err := tx.NamedExec(queryInsert, items)
	return err
}

func (o *OrderServiceCommandRepository) processEquipments(equipments []entities.Equipment, orderId int, tx *sqlx.Tx) error {
	if len(equipments) < 1 {
		return nil
	}

	for index, equipment := range equipments {
		equipment.OrderId = orderId
		equipments[index] = equipment
	}

	err := o.syncEquipments(equipments, tx)
	if err != nil {
		tx.Rollback()
		return err
	}

	return nil
}

func (o *OrderServiceCommandRepository) syncEquipments(equipments []entities.Equipment, tx *sqlx.Tx) error {
	queryDelete := `DELETE FROM equipments WHERE order_id = ?`
	tx.MustExec(queryDelete, equipments[0].OrderId)
	queryInsert := `INSERT INTO equipments (description, defect, observations, order_id) VALUES (:description, :defect, :observations, :order_id)`
	_, err := tx.NamedExec(queryInsert, equipments)
	return err
}

func (o *OrderServiceCommandRepository) processOrderPayments(orderPayments []entities.OrderPayment, orderId int, tx *sqlx.Tx) error {
	if len(orderPayments) < 1 {
		return nil
	}

	for index, orderPayment := range orderPayments {
		orderPayment.OrderId = orderId
		orderPayments[index] = orderPayment
	}

	err := o.addOrderPayments(orderPayments, tx)
	if err != nil {
		tx.Rollback()
		return err
	}

	return nil
}

func (o *OrderServiceCommandRepository) addOrderPayments(orderPayments []entities.OrderPayment, tx *sqlx.Tx) error {
	queryDelete := `DELETE FROM order_payments WHERE order_id = ?`
	tx.MustExec(queryDelete, orderPayments[0].OrderId)
	queryInsert := `
			INSERT INTO order_payments 
			    (order_id, description, total_value, installments, installment_value) 
			VALUES 
			    (:order_id, :description, :total_value, :installments, :installment_value)`
	_, err := tx.NamedExec(queryInsert, orderPayments)
	return err
}

type OrderServiceQueryRepository struct {
	connection *sqlx.DB
}

func NewOrderServiceQueryRepository(conn *sqlx.DB) *OrderServiceQueryRepository {
	return &OrderServiceQueryRepository{
		connection: conn,
	}
}

func (o *OrderServiceQueryRepository) FindOne(id int) (entities.OrderService, error) {
	var serviceOrder entities.OrderService
	query := `SELECT id, number, description, user_id, close_date, status_id, created_at, paid 
		FROM order_service	WHERE id = ?`
	err := o.connection.Get(&serviceOrder, query, id)
	if err != nil {
		return serviceOrder, err
	}
	items, err := o.GetOrderItems(serviceOrder.Id)
	if err != nil {
		return serviceOrder, err
	}
	serviceOrder.Items = items

	equipments, err := o.GetEquipments(serviceOrder.Id)
	if err != nil {
		return serviceOrder, err
	}
	serviceOrder.Equipments = equipments

	payments, err := o.GetPayments(serviceOrder.Id)
	if err != nil {
		return serviceOrder, err
	}
	serviceOrder.Payments = payments

	return serviceOrder, err
}

func (o *OrderServiceQueryRepository) FindAll() ([]entities.OrderService, error) {
	var ordersService []entities.OrderService
	query := `SELECT id, number, description, user_id, close_date, status_id, paid, created_at FROM order_service`
	err := o.connection.Select(&ordersService, query)
	return ordersService, err
}

func (o *OrderServiceQueryRepository) GetOrderItems(orderId int) ([]entities.OrderItem, error) {
	var orderItems []entities.OrderItem
	query := `SELECT id, product_id, order_id, type, quantity, value FROM order_items WHERE order_id = ?`
	err := o.connection.Select(&orderItems, query, orderId)
	if err != nil {
		return nil, err
	}

	if len(orderItems) < 1 {
		return orderItems, nil
	}

	var productsIds []int
	for _, orderItem := range orderItems {
		productsIds = append(productsIds, orderItem.ProductId)
	}

	products, err := o.getProductsItem(productsIds)
	if err != nil {
		return nil, err
	}

	for index, orderItem := range orderItems {
		for _, product := range products {
			if orderItem.ProductId == product.Id {
				orderItem.Product = &product
			}
		}
		orderItems[index] = orderItem
	}
	return orderItems, err
}

func (o *OrderServiceQueryRepository) GetEquipments(orderId int) ([]entities.Equipment, error) {
	var equipments []entities.Equipment
	query := `SELECT id, description, defect, observations, order_id FROM equipments WHERE order_id = ?`
	err := o.connection.Select(&equipments, query, orderId)
	return equipments, err
}

func (o *OrderServiceQueryRepository) GetPayments(orderId int) ([]entities.OrderPayment, error) {
	var payments []entities.OrderPayment
	query := `SELECT id, order_id, description, total_value,installments,installment_value FROM order_payments WHERE order_id = ?`
	err := o.connection.Select(&payments, query, orderId)
	return payments, err
}

func (o *OrderServiceQueryRepository) getProductsItem(productIds []int) ([]entities.Product, error) {
	s, _ := json.Marshal(productIds)
	sIds := strings.Trim(string(s), "[]")
	products := []entities.Product{}
	query := fmt.Sprintf("SELECT * FROM products WHERE id IN (%s)", sIds)
	err := o.connection.Select(&products, query)
	return products, err
}

package actions

import (
	"log"
	"math/rand"
	"strconv"
	"time"

	"github.com/henriquerocha2004/cyber-tech-go/internal/entities"
)

type ServiceOrderActions struct {
	orderServiceCommandRepository entities.OrderServiceCommandRepository
	orderServiceQueryRepository   entities.OrderServiceQueryRepository
	orderServiceStatusRepository  entities.OrderServiceStatusRepository
	productRepository             entities.ProductRepository
}

type ServiceOrderInput struct {
	Id          int                 `json:"id,omitempty"`
	Number      string              `json:"number"`
	Description string              `json:"description" validate:"required"`
	UserId      int                 `json:"user_id" validate:"required,numeric,gt=0"`
	CloseDate   string              `json:"close_date"`
	StatusId    int                 `json:"status_id" validate:"required"`
	Equipments  []EquipmentInput    `json:"equipments"`
	Items       []ItemsInput        `json:"items"`
	Paid        bool                `json:"paid"`
	Payments    []OrderPaymentInput `json:"payments"`
	Total       float64             `json:"total"`
}

type EquipmentInput struct {
	Id           int    `json:"id,omitempty"`
	Description  string `json:"description" validate:"required"`
	Defect       string `json:"defect" validate:"required"`
	Observations string `json:"observations"`
	OrderId      int    `json:"order_id" validate:"required"`
}

type ItemsInput struct {
	Id        int     `json:"id,omitempty"`
	ProductId int     `json:"product_id" validate:"required"`
	OrderId   int     `json:"order_id" validate:"required"`
	Type      string  `json:"type" validate:"required"`
	Quantity  int     `json:"quantity" validate:"required"`
	Value     float64 `json:"value" validate:"required"`
}

type OrderPaymentInput struct {
	Id               int     `json:"id,omitempty"`
	OrderId          int     `json:"order_id" validate:"required"`
	Description      string  `json:"description" validate:"required"`
	TotalValue       float64 `json:"total_value" validate:"required"`
	Installments     int     `json:"installments" validate:"required"`
	InstallmentValue float64 `json:"installment_value" validate:"required"`
}

type OrderOutput struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

func NewServiceOrderActions(
	orderServiceComRepo entities.OrderServiceCommandRepository,
	orderServiceQueryRepo entities.OrderServiceQueryRepository,
	orderServiceStatusRepo entities.OrderServiceStatusRepository,
	productRepository entities.ProductRepository) *ServiceOrderActions {
	return &ServiceOrderActions{
		orderServiceCommandRepository: orderServiceComRepo,
		orderServiceQueryRepository:   orderServiceQueryRepo,
		orderServiceStatusRepository:  orderServiceStatusRepo,
		productRepository:             productRepository,
	}
}

func (s *ServiceOrderActions) Create(input ServiceOrderInput) OrderOutput {
	orderStatus, err := s.getOrderStatus(input.StatusId)
	if err != nil || orderStatus.Id == 0 {
		return OrderOutput{
			Error:   true,
			Message: "Failed to create order: Failed to get status order",
		}
	}

	order := &entities.OrderService{
		Number:      s.generateServiceOrderNumber(),
		Description: input.Description,
		StatusId:    input.StatusId,
		Status:      orderStatus,
		Paid:        input.Paid,
	}

	order.Items, err = s.generateItems(input.Items)
	if err != nil {
		return OrderOutput{
			Error:   true,
			Message: "Failed to create order: Failed to generate items",
		}
	}

	order.Total, err = order.GetTotal()
	if err != nil {
		return OrderOutput{
			Error:   true,
			Message: "Failed to create order: Failed to calculate total of items",
		}
	}

	order.Equipments = s.generateEquipments(input.Equipments)
	order.UserId = input.UserId
	order.CloseDate = input.CloseDate
	order.Payments = s.generateOrderPayments(input.Payments)
	_, err = s.orderServiceCommandRepository.Create(*order)
	if err != nil {
		log.Println(err)
		return OrderOutput{
			Error:   true,
			Message: "Failed to create order!",
		}
	}
	return OrderOutput{
		Error:   false,
		Message: "Order created successfully",
		Data:    map[string]string{"order_number": order.Number},
	}
}

func (s *ServiceOrderActions) Update(input ServiceOrderInput) OrderOutput {
	if input.Number == "" {
		return OrderOutput{
			Error:   true,
			Message: "Failed to create order: Order Number not informed",
		}
	}

	orderStatus, err := s.getOrderStatus(input.StatusId)
	if err != nil {
		return OrderOutput{
			Error:   true,
			Message: "Failed to create order: Failed to get status order",
		}
	}

	order := &entities.OrderService{
		Id:          input.Id,
		Number:      input.Number,
		Description: input.Description,
		StatusId:    input.StatusId,
		Status:      orderStatus,
		Paid:        input.Paid,
	}

	order.Items, err = s.generateItems(input.Items)
	if err != nil {
		return OrderOutput{
			Error:   true,
			Message: "Failed to create order: Failed to generate items",
		}
	}

	order.Total, err = order.GetTotal()
	if err != nil {
		return OrderOutput{
			Error:   true,
			Message: "Failed to create order: Failed to calculate total of items",
		}
	}

	order.Equipments = s.generateEquipments(input.Equipments)
	order.UserId = input.UserId
	order.CloseDate = input.CloseDate
	order.Payments = s.generateOrderPayments(input.Payments)

	if order.Paid {
		order.CloseDate = time.Now().Format("2006-01-02 15:04")
	}

	if order.Paid && len(order.Payments) < 1 {
		return OrderOutput{
			Error:   true,
			Message: "Failed to create order: Payments not informed",
		}
	}

	err = s.orderServiceCommandRepository.Update(*order)
	if err != nil {
		log.Println(err)
		return OrderOutput{
			Error:   true,
			Message: "Failed to update order!",
		}
	}

	return OrderOutput{
		Error:   false,
		Message: "Order updated successfully",
	}
}

func (s *ServiceOrderActions) GetOne(id int) OrderOutput {
	order, err := s.orderServiceQueryRepository.FindOne(id)
	if err != nil {
		log.Println(err)
		return OrderOutput{
			Error:   true,
			Message: "Failed to get order!",
		}
	}

	return OrderOutput{
		Error: false,
		Data:  order,
	}
}

func (s *ServiceOrderActions) GetAll() OrderOutput {
	orders, err := s.orderServiceQueryRepository.FindAll()
	if err != nil {
		log.Println(err)
		return OrderOutput{
			Error:   true,
			Message: "Failed to get orders!",
		}
	}

	return OrderOutput{
		Error: false,
		Data:  orders,
	}
}

func (s *ServiceOrderActions) getOrderStatus(id int) (*entities.OrderServiceStatus, error) {
	orderStatus, err := s.orderServiceStatusRepository.FindOne(id)
	if err != nil || orderStatus.Id == 0 {
		return nil, err
	}
	return &orderStatus, nil
}

func (s *ServiceOrderActions) generateServiceOrderNumber() string {
	currentTime := time.Now().Format("200601021504")
	randomInt := rand.Intn(1000)
	return currentTime + strconv.Itoa(randomInt)
}

func (s *ServiceOrderActions) generateItems(inputOrderItems []ItemsInput) ([]entities.OrderItem, error) {
	log.Println(inputOrderItems)

	if len(inputOrderItems) < 1 {
		return []entities.OrderItem{}, nil
	}

	var orderItems []entities.OrderItem

	for _, item := range inputOrderItems {
		var orderItem entities.OrderItem
		orderItem.OrderId = item.OrderId
		err := orderItem.SetTypeItem(item.Type)
		if err != nil {
			return []entities.OrderItem{}, err
		}
		orderItem.Quantity = item.Quantity
		orderItem.ProductId = item.ProductId
		orderItem.Value = item.Value
		orderItems = append(orderItems, orderItem)
	}
	return orderItems, nil
}

func (s *ServiceOrderActions) generateEquipments(inputEquipments []EquipmentInput) []entities.Equipment {
	if len(inputEquipments) < 1 {
		return []entities.Equipment{}
	}

	var equipments []entities.Equipment
	for _, inputEquipment := range inputEquipments {
		equipment := &entities.Equipment{
			Description:  inputEquipment.Description,
			Defect:       inputEquipment.Defect,
			Observations: inputEquipment.Observations,
		}
		equipments = append(equipments, *equipment)
	}

	return equipments
}

func (s *ServiceOrderActions) generateOrderPayments(inputOrderPayments []OrderPaymentInput) []entities.OrderPayment {
	if len(inputOrderPayments) < 1 {
		return []entities.OrderPayment{}
	}

	var orderPayments []entities.OrderPayment
	for _, inputOrderPayment := range inputOrderPayments {
		orderPayment := &entities.OrderPayment{
			Description:      inputOrderPayment.Description,
			TotalValue:       inputOrderPayment.TotalValue,
			Installments:     inputOrderPayment.Installments,
			InstallmentValue: inputOrderPayment.InstallmentValue,
		}

		orderPayments = append(orderPayments, *orderPayment)
	}
	return orderPayments
}

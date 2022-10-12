package entities

type OrderService struct {
	Id          int                 `json:"id,omitempty" db:"id,omitempty"`
	Number      string              `json:"number" db:"number"`
	Description string              `json:"description" db:"description"`
	UserId      int                 `json:"user_id" db:"user_id"`
	CloseDate   string              `json:"close_date" db:"close_date,omitempty"`
	Status      *OrderServiceStatus `json:"status,omitempty"`
	StatusId    int                 `json:"status_id" db:"status_id"`
	Equipments  []Equipment         `json:"equipments,omitempty"`
	Items       []OrderItem         `json:"items,omitempty"`
	CreatedAt   string              `json:"created_at" db:"created_at"`
	UpdatedAt   string              `json:"updated_at,omitempty" db:"updated_at"`
	CreatedBy   *User               `json:"created_by,omitempty" db:"created_by"`
	Paid        bool                `json:"paid" db:"paid"`
	Payments    []OrderPayment      `json:"payments,omitempty"`
	Total       float64             `json:"total" db:"total"`
}

func (o *OrderService) AddEquipment(equipment Equipment) {
	o.Equipments = append(o.Equipments, equipment)
}

func (o *OrderService) AddItem(item OrderItem) {
	o.Items = append(o.Items, item)
}

func (o *OrderService) AddPaymentForm(paymentForm OrderPayment) {
	o.Payments = append(o.Payments, paymentForm)
}

func (o *OrderService) GetTotal() (float64, error) {
	if len(o.Items) < 1 {
		return 0.00, nil
	}

	var total float64
	for _, item := range o.Items {
		total += item.Subtotal()
	}
	return total, nil
}

type OrderServiceCommandRepository interface {
	Create(orderService OrderService) (int, error)
	Update(orderService OrderService) error
}

type OrderServiceQueryRepository interface {
	FindOne(orderId int) (OrderService, error)
	FindAll() ([]OrderService, error)
	GetOrderItems(orderId int) ([]OrderItem, error)
	GetEquipments(orderId int) ([]Equipment, error)
}

type SendOrderServiceEvent interface {
	Send(order OrderService) error
}

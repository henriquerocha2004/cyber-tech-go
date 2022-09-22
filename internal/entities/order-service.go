package entities

type OrderService struct {
	Id           int                `json:"id,omitempty"`
	Number       string             `json:"number"`
	UserId       int                `json:"user_id"`
	Quantity     int                `json:"quantity"`
	CloseDate    string             `json:"close_date"`
	Status       OrderServiceStatus `json:"status"`
	Equipments   []Equipment        `json:"equipments"`
	Items        []OrderItem        `json:"items"`
	CreatedAt    string             `json:"created_at"`
	CreatedBy    User               `json:"created_by,omitempty"`
	Paid         bool               `json:"paid"`
	PaymentForms []PaymentForm      `json:"payment_forms"`
}

func (o *OrderService) AddEquipment(equipment Equipment) {
	o.Equipments = append(o.Equipments, equipment)
}

func (o *OrderService) AddItem(item OrderItem) {
	o.Items = append(o.Items, item)
}

func (o *OrderService) AddPaymentForm(paymentForm PaymentForm) {
	o.PaymentForms = append(o.PaymentForms, paymentForm)
}

func (o *OrderService) GetTotal() (float64, error) {
	var total float64
	for _, item := range o.Items {
		total += item.Subtotal()
	}
	return total, nil
}

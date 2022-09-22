package entities

const (
	STATUS_ACTIVE       string = "active"
	STATUS_INACTIVE     string = "inactive"
	SITUATION_PENDING   string = "pending"
	SITUATION_PAID      string = "paid"
	SITUATION_CANCELLED string = "cancelled"
	ENTRY_MOVEMENT      string = "entry"
	OUT_MOVEMENT        string = "out"
)

type FinancialMovement struct {
	Id               int     `json:"id,omitempty"`
	Status           string  `json:"status"`
	Situation        string  `json:"situation"`
	Type             string  `json:"type"`
	BudgetAccountId  int     `json:"budget_account_id"`
	Value            float64 `json:"value"`
	PaidValue        float64 `json:"paid_value"`
	IssueDate        string  `json:"issue_date"`
	DueDate          string  `json:"due_date"`
	PaymentForecast  string  `json:"payment_forecast"`
	CompensationDate string  `json:"compensation_date"`
	PayDay           string  `json:"payday"`
	Observation      string  `json:"observation"`
	OrderId          string  `json:"order_id,omitempty"`
}

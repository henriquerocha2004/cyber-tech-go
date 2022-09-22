package entities

const (
	BUDGET_ACCOUNT_TYPE_REVENUE = "revenue"
	BUDGET_ACCOUNT_TYPE_EXPENSE = "expense"
)

type BudgetAccount struct {
	Id          int    `json:"id,omitempty"`
	Description string `json:"description"`
	Type        string `json:"type"`
}

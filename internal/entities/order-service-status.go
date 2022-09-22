package entities

type OrderServiceStatus struct {
	Id              int    `json:"id,omitempty"`
	Description     string `json:"description"`
	LaunchFinancial bool   `json:"launch_financial"`
	Color           string `json:"color"`
}

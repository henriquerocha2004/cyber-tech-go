package entities

type Equipment struct {
	Id           int    `json:"id,omitempty"`
	Description  string `json:"description"`
	Defect       string `json:"defect"`
	Observations string `json:"observations"`
}

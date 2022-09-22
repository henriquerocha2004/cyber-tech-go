package entities

type Supplier struct {
	Id       int    `json:"id,omitempty"`
	Name     string `json:"name"`
	Document string `json:"document"`
	Address  string `json:"address"`
	District string `json:"district"`
	City     string `json:"city"`
	State    string `json:"state"`
}

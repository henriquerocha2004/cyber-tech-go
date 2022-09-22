package entities

type Address struct {
	Id       int    `json:"id,omitempty" db:"id"`
	Street   string `json:"street" db:"street"`
	City     string `json:"city" db:"city"`
	District string `json:"district" db:"district"`
	State    string `json:"state" db:"state"`
	Country  string `json:"country" db:"country"`
	ZipCode  string `json:"zip_code" db:"zip_code"`
	Type     string `json:"type" db:"type"`
	UserId   int    `json:"user_id" db:"user_id"`
}

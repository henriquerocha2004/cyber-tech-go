package entities

type Contact struct {
	Id         int    `json:"id,omitempty" db:"id"`
	Type       string `json:"type" db:"type"`
	Phone      string `json:"phone" db:"phone"`
	IsWhatsApp bool   `json:"is_whatsapp" db:"is_whatsapp"`
	UserId     int    `json:"user_id" db:"user_id"`
}

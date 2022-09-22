package entities

type Business struct {
	Id            int    `json:"id,omitempty"`
	FantasyName   string `json:"fantasy_name"`
	SocialReason  string `json:"social_reason"`
	Address       string `json:"address"`
	City          string `json:"city"`
	District      string `json:"district"`
	State         string `json:"state"`
	Country       string `json:"country"`
	ZipCode       string `json:"zip_code"`
	Status        string `json:"status"`
	BlockedReason string `json:"blocked_reason"`
}

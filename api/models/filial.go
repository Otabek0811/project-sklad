package models

type FilialPrimaryKey struct {
	Id string `json:"id"`
}

type CreateFilial struct {
	Name        string   `json:"name"`
	Address     string   `json:"address"`
	PhoneNumber string   `json:"phone_number"`
}

type Filial struct {
	Id          string     `json:"id"`
	Name        string     `json:"name"`
	Address     string     `json:"address"`
	PhoneNumber string     `json:"phone_number"`
	CreatedAt   string     `json:"created_at"`
	UpdatedAt   string     `json:"updated_at"`
}

type UpdateFilial struct {
	Id          string   `json:"id"`
	Name        string   `json:"name"`
	Address     string   `json:"address"`
	PhoneNumber string   `json:"phone_number"`	
}

type FilialGetListRequest struct {
	Offset int    `json:"offset"`
	Limit  int    `json:"limit"`
	Search string `json:"search"`
}

type FilialGetListResponse struct {
	Count   int       `json:"count"`
	Filials []*Filial `json:"filials"`
}

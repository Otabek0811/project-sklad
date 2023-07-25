package models

type RemainderPrimaryKey struct {
	Id string `json:"id"`
}

type CreateRemainder struct {
	FilialID   string  `json:"filial_id"`
	CategoryID string  `json:"categroy_id"`
	Name       string  `json:"name"`
	Barcode    string  `json:"barcode"`
	Amount     int     `json:"amount"`
	Price      float64 `json:"price"`
}

type Remainder struct {
	Id         string  `json:"id"`
	FilialID   string  `json:"filial_id"`
	CategoryID string  `json:"categroy_id"`
	Name       string  `json:"name"`
	Barcode    string  `json:"barcode"`
	Amount     int     `json:"amount"`
	Price      float64 `json:"price"`
	TotalPrice float64 `json:"total_price"`
	CreatedAt  string  `json:"created_at"`
	UpdatedAt  string  `json:"updated_at"`
}

type UpdateRemainder struct {
	Id         string  `json:"id"`
	FilialID   string  `json:"filial_id"`
	CategoryID string  `json:"categroy_id"`
	Name       string  `json:"name"`
	Barcode    string  `json:"barcode"`
	Amount     int     `json:"amount"`
	Price      float64 `json:"price"`
}

type RemainderGetListRequest struct {
	Offset int    `json:"offset"`
	Limit  int    `json:"limit"`
	SearchByFilial string `json:"search_by_filial"`
	SearchByCategory string `json:"search_by_category"`
	SearchByBarcode string `json:"search_by_barcode"`
}

type RemainderGetListResponse struct {
	Count      int          `json:"count"`
	Remainders []*Remainder `json:"Remainders"`
}

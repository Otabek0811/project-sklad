package models

type ComingProductsPrimaryKey struct {
	Id string `json:"id"`
}

type CreateComingProducts struct {
	ComingID   string  `json:"coming_id"`
	CategoryID string  `json:"categroy_id"`
	Name       string  `json:"name"`
	Barcode    string  `json:"barcode"`
	Amount     int     `json:"amount"`
	Price      float64 `json:"price"`
}

type ComingProducts struct {
	Id         string  `json:"id"`
	ComingID   string  `json:"coming_id"`
	CategoryID string  `json:"categroy_id"`
	Name       string  `json:"name"`
	Barcode    string  `json:"barcode"`
	Amount     int     `json:"amount"`
	Price      float64 `json:"price"`
	TotalPrice float64 `json:"total_price"`
	CreatedAt  string  `json:"created_at"`
	UpdatedAt  string  `json:"updated_at"`
}

type UpdateComingProducts struct {
	Id         string  `json:"id"`
	ComingID   string  `json:"coming_id"`
	CategoryID string  `json:"categroy_id"`
	Name       string  `json:"name"`
	Barcode    string  `json:"barcode"`
	Amount     int     `json:"amount"`
	Price      float64 `json:"price"`
}

type ComingProductsGetListRequest struct {
	Offset             int    `json:"offset"`
	Limit              int    `json:"limit"`
	SearchBYComingID  string   `json:"search_by_coming_id"`
	SearchByBarcode    string `json:"search_by_barcode"`
	SearchByCategoryID string `json:"search_by_category"`
}

type ComingProductsGetListResponse struct {
	Count          int               `json:"count"`
	ComingProducts []*ComingProducts `json:"ComingProducts"`
}

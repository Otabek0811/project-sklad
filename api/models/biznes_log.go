package models


type Scan_Barcode struct{
	ComingID   string  `json:"coming_id"`
	Barcode    string   `json:"barcode"`
}

type ResponseBarcode struct{
	ComingID   string  `json:"coming_id"`
	Name       string  `json:"name"`
	Barcode    string   `json:"barcode"`
	Price      float64 `json:"price"`
	CategoryId string  `json:"category_id"`
}

type RespProduct struct{ 
	FilialID string `json:"filial_id"`
	ComingID   string  `json:"coming_id"`
	CategoryID string  `json:"categroy_id"`
	Name       string  `json:"name"`
	Barcode    string  `json:"barcode"`
	Amount     int     `json:"amount"`
	Price      float64 `json:"price"`
	TotalPrice float64 `json:"total_price"`
}
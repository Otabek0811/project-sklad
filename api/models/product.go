package models

type ProductPrimaryKey struct {
	Id string `json:"id"`
}
type ProductBarcodeKey struct {
	Barcode string `json:"barcode"`
}

type CreateProduct struct {
	Name       string  `json:"name"`
	Price      float64 `json:"price"`
	Barcode    string   `json:"barcode"`
	CategoryId string  `json:"category_id"`
}

type Product struct {
	Id         string  `json:"id"`
	Name       string  `json:"name"`
	Price      float64 `json:"price"`
	Barcode    string   `json:"barcode"`
	CategoryId string  `json:"category_id"`
	CreatedAt  string  `json:"created_at"`
	UpdatedAt  string  `json:"updated_at"`
}

type UpdateProduct struct {
	Id         string  `json:"id"`
	Name       string  `json:"name"`
	Price      float64 `json:"price"`
	Barcode    string   `json:"barcode"`
	CategoryId string  `json:"category_id"`
}

type ProductGetListRequest struct {
	Offset int    `json:"offset"`
	Limit  int    `json:"limit"`
	SearchByBarcode string `json:"search_by_barcode"`
	SearchByName string `json:"search_by_name"`

}

type ProductGetListResponse struct {
	Count    int        `json:"count"`
	Products []*Product `json:"products"`
}

package entity

import "time"

type Mexc struct {
	Name string `json:"hello"`
}

type ProductInformation struct {
	ProductReference  string     `json:"product_reference"`
	ProductName       string     `json:"product_name"`
	AddedDateUtc      *time.Time `json:"added_date_utc"`
	AddedDate         string     `json:"added_date"`
	Status            string     `json:"status"`
	ProductCategory   string     `json:"product_category"`
	Price             float64    `json:"price"`
	StockLocation     string     `json:"stock_location"`
	Supplier          string     `json:"supplier"`
	AvailableQuantity int64      `json:"available_quantity"`
}

type Product struct {
	Id           string
	Reference    string
	Name         string
	AddedDate    *time.Time
	Status       string
	CategoryId   string
	Price        float64
	StockCity    string
	SupplierCity string
	SupplierId   string
	Quantity     string
}

type GetListProductInformationResponse struct {
	ListProductInformation []*ProductInformation
}

type GetListProductInformationRequest struct {
	Limit  int `json:"limit" query:"limit"`
	Offset int `json:"offset" query:"offset"`
}

type ProductStatisticPerCategory struct {
	CategoryId   string  `json:"category_id"`
	CategoryName string  `json:"category_name"`
	TotalProduct int64   `json:"total_product"`
	Percentage   float64 `json:"percentage"`
}

type GetListProductStatisticPerCategoryResponse struct {
	ListProductStatistic []*ProductStatisticPerCategory
}

type ProductStatisticPerSupplier struct {
	SupplierId   string  `json:"supplier_id"`
	SupplierName string  `json:"supplier_name"`
	TotalProduct int64   `json:"total_product"`
	Percentage   float64 `json:"percentage"`
}

type GetListProductStatisticPerSupplierResponse struct {
	ListProductStatistic []*ProductStatisticPerSupplier
}

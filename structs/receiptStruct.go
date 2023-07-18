package structs

type Item struct {
	ShortDescription string `json:"shortDescription"  validate:"required"`
	Price            string `json:"price"  validate:"required,money"`
}

type Receipt struct {
	Retailer     string `json:"retailer" validate:"required"`
	PurchaseDate string `json:"purchaseDate"  validate:"required,date"`
	PurchaseTime string `json:"purchaseTime" validate:"required,time"`
	Items        []Item `json:"items"  validate:"required,dive"`
	Total        string `json:"total" validate:"required,money"`
}

package http_dto

//
// ProductDTO
// @Description: 产品DTO
//
type ProductDTO struct {
	ShopId      uint    `json:"shop_id"`
	ProductName string  `json:"product_name"`
	OriPrice    float64 `json:"ori_price"`
	Price       float64 `json:"price"`
	Brief       string  `json:"brief"`
	Content     string  `json:"content"`
	Pic         string  `json:"pic"`
	Status      int     `json:"status"`
	CategoryId  int     `json:"category_id"`
	SoldNum     int     `json:"sold_num"`
	TotalStocks int     `json:"total_stocks"`
	PutOnTime   int64   `json:"put_on_time"`
}

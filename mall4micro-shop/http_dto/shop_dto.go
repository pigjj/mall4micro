package http_dto

type ShopDTO struct {
	ID         uint   `json:"id"`
	ShopName   string `json:"shop_name"`
	ShopDesc   string `json:"shop_desc"`
	ShopPic    string `json:"shop_pic"`
	ShopStatus int    `json:"shop_status"`
}

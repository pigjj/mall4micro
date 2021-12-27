package http_dto

//
// CategoryDTO
// @Description: 产品种类DTO
//
type CategoryDTO struct {
	ID           uint   `json:"id"`
	ShopId       uint   `json:"shop_id"`
	ParentId     uint   `json:"parent_id"`
	CategoryName string `json:"category_name"`
	Icon         string `json:"icon"`
	Pic          string `json:"pic"`
	Status       int    `json:"status"`
}

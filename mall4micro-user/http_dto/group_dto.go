package http_dto

//
// GroupDTO
// @Description: 权限组数据传输DTO
//
type GroupDTO struct {
	ID             uint            `json:"id"`
	GroupName      string          `json:"group_name"`
	GroupDesc      string          `json:"group_desc"`
	ShopId         uint            `json:"shop_id"`
	PermissionList []PermissionDTO `json:"permission_list"`
}

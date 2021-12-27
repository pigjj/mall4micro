package http_dto

//
// PermissionDTO
// @Description: 权限数据传输DTO
//
type PermissionDTO struct {
	ID             uint   `json:"id"`
	PermissionName string `json:"permission_name"`
	PermissionDesc string `json:"permission_desc"`
}

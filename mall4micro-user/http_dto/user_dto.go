package http_dto

//
// UserDTO
// @Description: 用户数据传输DTO
//
type UserDTO struct {
	ID             uint            `json:"id"`
	Username       string          `json:"username"`
	Email          string          `json:"email"`
	Mobile         string          `json:"mobile"`
	Status         int             `json:"status"`
	GroupList      []GroupDTO      `json:"group_list"`
	PermissionList []PermissionDTO `json:"permission_list"`
}

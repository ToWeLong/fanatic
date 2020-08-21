package form

type PermissionForm struct {
	UserId uint `json:"user_id" binding:"required,gte=1"`
	PermissionId uint `json:"permission_id" binding:"required,gte=1"`
}

type PermissionEditForm struct {
	UserId uint `json:"user_id" binding:"required,gte=1"`
	NewPermissionId uint `json:"new_permission_id" binding:"required,gte=1"`
	OldPermissionId uint `json:"old_permission_id" binding:"required,gte=1"`
}

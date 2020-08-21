package service

import (
	"errors"
	"fanatic/lib/redis"
	"fanatic/model"
	"fanatic/vo"
	"fmt"
	"log"
)

type PermissionService interface {
	RegisterPermission(permission, module string, mount int) bool
	FindPermission(permission string) bool
	FindPermissionById(permissionId uint) model.Permission
	FindAllPermission() map[string][]vo.Permission
	FindPermissionByIds(ids []uint) []model.Permission
	FindUserPermission(uid uint) vo.UserInfo
	// user_permission相关接口
	EditUserPermission(uid, oldPermissionId, newPermissionId uint) error
	RegisterUserPermission(uid, permissionId uint) error
	DeleteUserPermission(uid, permissionId uint) error
}
type permissionService struct {
}

func NewPermissionService() PermissionService {
	return &permissionService{}
}

func (p permissionService) RegisterPermission(permission, module string, mount int) bool {
	permiss := model.Permission{
		Name:   permission,
		Module: module,
		Mount:  mount,
	}
	pSerice := model.DB.Create(&permiss)
	if pSerice.Error != nil {
		return false
	}
	return true
}

// 为空是true
// 不为空是false
func (p *permissionService) FindPermission(permission string) bool {
	var permiss model.Permission
	model.DB.Where("name = ?", permission).First(&permiss)
	return permiss.Name == ""
}

func (p *permissionService) FindPermissionById(permissionId uint) model.Permission {
	var permiss model.Permission
	model.DB.Where("id = ?", permissionId).First(&permiss)
	return permiss
}

func (p *permissionService) FindPermissionByIds(ids []uint) []model.Permission {
	var permiss []model.Permission
	model.DB.Where("id IN (?)", ids).Find(&permiss)
	return permiss
}

func (p *permissionService) FindAllPermission() map[string][]vo.Permission {
	var permiss []model.Permission
	var permissVo []vo.Permission

	model.DB.Select("module").Group("module").Find(&permiss)
	m := make(map[string][]vo.Permission, len(permiss))
	for i := 0; i < len(permiss); i++ {
		module := permiss[i].Module
		model.DB.Where("module = ?", module).Find(&permissVo)
		m[module] = permissVo
	}
	return m
}

func (p *permissionService) FindUserPermission(uid uint) vo.UserInfo {
	var (
		userInfo      vo.UserInfo
		permissList   []vo.Permission
		permiss       []model.Permission
		userPermiss   []model.UserPermission
		permissionIds []uint
	)

	// 从userPermiss表中拿到 uid 和 permission_id
	model.DB.Where("user_id = ?", uid).Find(&userPermiss)

	for i := 0; i < len(userPermiss); i++ {
		// 一组 permission_id
		id := userPermiss[i].PermissionId
		permissionIds = append(permissionIds, id)
	}
	if len(permissionIds) > 0 {
		var uService = NewUserService()
		user := uService.FindOneById(int32(uid))
		userInfo.ID = user.ID
		userInfo.Account = user.Account
		permiss = p.FindPermissionByIds(permissionIds)
		permissList = make([]vo.Permission, len(permissionIds))
		for i := 0; i < len(permissionIds); i++ {
			permissList[i].ID = permiss[i].ID
			permissList[i].Name = permiss[i].Name
		}
		userInfo.Permission = permissList
	}
	return userInfo
}

// user_permission相关接口
var redisService = redis.NewRedis()

func (p permissionService) EditUserPermission(uid, oldPermissionId, newPermissionId uint) error {
	var userPermission model.UserPermission
	if p.FindPermissionById(newPermissionId).Name == "" {
		return errors.New("该权限不存在")
	}
	model.DB.Where("user_id = ? AND permission_id = ?", uid, newPermissionId).First(&userPermission)
	// 判断user_permission表中是否存在该记录
	if userPermission.PermissionId > 0 {
		return errors.New("该用户权限已存在")
	}
	update := model.DB.Model(&userPermission).
		Where("user_id = ? AND permission_id = ?", uid, oldPermissionId).
		Update("permission_id", newPermissionId)
	key := fmt.Sprintf("user:%d:permission",uid)
	err := redisService.Delete(key)
	if err != nil {
		log.Println("redis err",err)
	}
	return update.Error
}

func (p permissionService) RegisterUserPermission(uid, permissionId uint) error {
	var userPermission1 model.UserPermission

	// 判断permission表中是否存在该权限
	if p.FindPermissionById(permissionId).Name == "" {
		return errors.New("该权限不存在")
	}
	model.DB.Where("user_id = ? AND permission_id = ?", uid, permissionId).First(&userPermission1)
	// 判断user_permission表中是否存在该记录
	if userPermission1.PermissionId > 0 {
		return errors.New("该用户权限已存在")
	}
	userPermission := model.UserPermission{
		UserId:       uid,
		PermissionId: permissionId,
	}
	key := fmt.Sprintf("user:%d:permission",uid)
	err := redisService.Delete(key)
	if err != nil {
		log.Println("redis err",err)
	}
	return model.DB.Create(&userPermission).Error
}
func (p *permissionService) DeleteUserPermission(uid, permissionId uint) error {
	var userPermission model.UserPermission
	model.DB.Where("user_id = ? AND permission_id = ?", uid, permissionId).First(&userPermission)
	if userPermission.PermissionId == 0 {
		return errors.New("该用户权限不存在")
	}
	db := model.DB.Where("user_id = ? AND permission_id = ?", uid, permissionId).Delete(&userPermission)
	key := fmt.Sprintf("user:%d:permission",uid)
	err := redisService.Delete(key)
	if err != nil {
		log.Println("redis err",err)
	}
	return db.Error
}

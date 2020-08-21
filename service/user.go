package service

import (
	"fanatic/lib/crypt"
	"fanatic/lib/token"
	"fanatic/model"
)

type UserService interface {
	UserLogin(account, password string) (bool, string, string)
	RegisterUser(account, password string) bool
	FindOneByAccount(account string) model.User
	FindOneById(id int32) model.User
}

type userService struct {
}

func NewUserService() UserService {
	return &userService{}
}

// 登录
// return false "" "" 表示登录失败
// return true access_token,refresh_token 表示登录成功
func (u *userService) UserLogin(account, password string) (bool, string, string) {
	user := u.FindOneByAccount(account)
	if user.Account != "" {
		access_token := token.CreateToken(token.ACCESS, int32(user.ID))
		refresh_token := token.CreateToken(token.REFRESH, int32(user.ID))
		return crypt.VerfifyPsw([]byte(password), []byte(user.Password)), access_token, refresh_token
	}
	return false, "", ""
}

// 注册
// return true: 表示含有这条记录
// return false: 表示无这条记录
func (u *userService) RegisterUser(account, password string) bool {
	password = string(crypt.EncodePassword(password))
	user := model.User{
		Account:  account,
		Password: password,
	}
	createUser := model.DB.Create(&user)
	// 判断数据库是否存在该字段
	if createUser.Error != nil {
		return true
	}
	return false
}

func (u *userService) FindOneByAccount(account string) model.User {
	var user model.User
	model.DB.Where("account = ?", account).First(&user)
	return user
}

func (u *userService) FindOneById(id int32) model.User {
	var user model.User
	model.DB.Where("id = ?", id).First(&user)
	return user
}


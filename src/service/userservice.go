package service

import (
	"chatroom/src/dataorm"
	"errors"
)

type userService interface {
	CreateUser(UserName, FirstName, LastName, Phone, Email, Password string) error
	UserInfo(username string) (dataorm.User, error)
	Login(username, password string) error
}

type Userservice struct {
}

//@function: 新建用户
//@param: 输入参数
//@return: 如果失败，返回错误

func (r Userservice) CreateUser(UserName, FirstName, LastName, Phone, Email, Password string) error {
	user := dataorm.User{
		Name:      UserName,
		Firstname: FirstName,
		Lastname:  LastName,
		Phone:     Phone,
		Email:     Email,
		Password:  Password,
		Status:    "",
	}
	if err := dataorm.Insert(user); err != nil {
		return err
	}
	return nil
}

//@function: 返回用户信息
//@param: username：用户名
//@return: 如果失败，返回错误

func (r Userservice) UserInfo(username string) (dataorm.User, error) {
	users, err := dataorm.Query("User", nil, []string{"name"}, []string{username})
	if err != nil {
		return dataorm.User{}, err
	}

	uservalue, _ := users.([]dataorm.User)
	return uservalue[0], nil
}

//@function: 用户登陆
//@return：如果登陆失败，返回{}与error，成功返回user信息
func (r Userservice) Login(username string, password string) (dataorm.User, error) {
	users, err := dataorm.Query("User", nil, []string{"name", "password"}, []string{username, password})
	if err != nil {
		return dataorm.User{}, err
	}
	uservalue, _ := users.([]dataorm.User)
	if len(uservalue) != 1 {
		return dataorm.User{}, errors.New("wrong password")
	}
	return uservalue[0], nil
}

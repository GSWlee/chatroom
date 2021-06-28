package service

import "chatroom/src/dataorm"

type userService interface {
	CreateUser(FirstName, LastName, Phone, Email, Password string) error
	UserInfo(username string) (dataorm.User, error)
	//todo
	Login()
}

type Userservice struct {
}

//@function: 新建用户
//@param: 输入参数
//@return: 如果失败，返回错误

func (r Userservice) CreateUser(FirstName, LastName, Phone, Email, Password string) error {
	user := dataorm.User{
		Name:      FirstName + " " + LastName,
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

func (r Userservice) Userinfo(username string) (dataorm.User, error) {
	users, err := dataorm.Query("User", nil, []string{"name"}, []string{username})
	if err != nil {
		return dataorm.User{}, err
	}

	uservalue, _ := users.([]dataorm.User)
	return uservalue[0], nil
}

//todo
func (r Userservice) Login() {

}

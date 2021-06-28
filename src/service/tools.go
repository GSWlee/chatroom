package service

import (
	"chatroom/src/dataorm"
	"errors"
	"strconv"
)

func getRoomID(RoomName string) (int, error) {
	rooms, err := dataorm.Query("Room", nil, []string{"name"}, []string{RoomName})
	if err != nil {
		return 0, err
	}

	roomvalue, _ := rooms.([]dataorm.Room)
	room := roomvalue[0]
	return room.ID, nil
}

func getUserID(UserName string) (int, error) {
	users, err := dataorm.Query("User", nil, []string{"name"}, []string{UserName})
	if err != nil {
		return 0, err
	}

	uservalue, _ := users.([]dataorm.User)
	user := uservalue[0]
	return user.ID, nil
}

func getUserName(UserId int) string {
	users, err := dataorm.Query("User", nil, []string{"i_d"}, []string{strconv.Itoa(UserId)})
	if err != nil {
		return ""
	}

	uservalue, _ := users.([]dataorm.User)
	user := uservalue[0]
	return user.Name
}

//@function: 根据用户名返回id与所在房间id
//@param: username:用户名
//@return: 如果失败，返回错误信息

func GetTwoID(UserName string) (int, int, error) {
	users, err := dataorm.Query("User", nil, []string{"name"}, []string{UserName})
	if err != nil {
		return 0, 0, err
	}

	uservalue, _ := users.([]dataorm.User)
	user := uservalue[0]
	if user.Status == "" {
		return 0, 0, errors.New("User not in anyroom!!")
	}
	roomid, err := getRoomID(user.Status)
	if err != nil {
		return 0, 0, err
	}
	return user.ID, roomid, nil
}

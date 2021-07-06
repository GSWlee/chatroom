package service

import (
	"chatroom/src/dataorm"
	"errors"
	"strconv"
)

//房间服务的接口

type roomService interface {
	CreateRoom(roomname string, username string) (int, error)
	EnterRoom(roomid int, username string) error
	LeaveRoom(username string) error
	GetAllUser(roomname string) ([]string, error)
	GetAllRoom() ([]Room, error)
}

//房间服务提供商

type RoomService struct {
}

type Room struct {
	Name string `json:"name"`
	Id   string `json:"id"`
}

//@function: 创建一个新房间
//@param1: roomname:房间名
//@param2: username:创建用户名
//@return: 如果创建失败，返回错误信息

func (r RoomService) CreateRoom(roomname string, username string) (int, error) {
	room := dataorm.Room{
		Name:    roomname,
		Creator: username,
		Data:    "",
	}
	if err := dataorm.Insert(room); err != nil {
		return 0, err
	}
	ID, err := getRoomID(roomname)
	if err != nil {
		return 0, err
	}
	return ID, nil
}

//@function: 进入房间
//@param1: roomname:房间名
//@param2: username:用户名
//@return: 如果进入失败，返回错误信息

//func (r RoomService) EnterRoom(roomname string, username string) error {
//	users, err := dataorm.Query("User", nil, []string{"name"}, []string{username})
//	if err != nil {
//		return err
//	}
//
//	uservalue, _ := users.([]dataorm.User)
//	user := uservalue[0]
//
//	roomid, err := getRoomID(roomname)
//	if err != nil {
//		return err
//	}
//
//	if user.Status == "" {
//		user.Status = roomname
//		if err = dataorm.Update(user); err != nil {
//			return err
//		}
//		tmp := dataorm.Userinroom{
//			Userid: user.ID,
//			Roomid: roomid,
//		}
//		if err = dataorm.Insert(tmp); err != nil {
//			return err
//		}
//	} else {
//		user.Status = roomname
//		if err = dataorm.Update(user); err != nil {
//			return err
//		}
//		userinrooms, err := dataorm.Query("Userinroom", nil, []string{"Userid"}, []string{strconv.Itoa(user.ID)})
//		if err != nil {
//			return err
//		}
//
//		userinroomvalue, _ := userinrooms.([]dataorm.Userinroom)
//		uir := userinroomvalue[0]
//		uir.Roomid = roomid
//		if err = dataorm.Insert(uir); err != nil {
//			return err
//		}
//	}
//	return nil
//}

func (r RoomService) EnterRoom(roomid int, username string) error {
	users, err := dataorm.Query("User", nil, []string{"name"}, []string{username})
	if err != nil {
		return err
	}

	uservalue, _ := users.([]dataorm.User)
	if uservalue == nil {
		return errors.New("User not find")
	}
	user := uservalue[0]

	roomname := getRoomName(roomid)
	if roomname == "" {
		return errors.New("Wrong room")
	}

	if user.Status == "" {
		user.Status = roomname
		if err = dataorm.Update(user); err != nil {
			return err
		}
		tmp := dataorm.Userinroom{
			Userid: user.ID,
			Roomid: roomid,
		}
		if err = dataorm.Insert(tmp); err != nil {
			return err
		}
	} else {
		user.Status = roomname
		if err = dataorm.Update(user); err != nil {
			return err
		}
		userinrooms, err := dataorm.Query("Userinroom", nil, []string{"Userid"}, []string{strconv.Itoa(user.ID)})
		if err != nil {
			return err
		}

		userinroomvalue, _ := userinrooms.([]dataorm.Userinroom)
		uir := userinroomvalue[0]
		uir.Roomid = roomid
		if err = dataorm.Insert(uir); err != nil {
			return err
		}
	}
	return nil
}

//@function: 离开房间
//@param: username:用户名
//@return: 如果离开失败，返回错误信息

func (r RoomService) LeaveRoom(username string) error {
	users, err := dataorm.Query("User", nil, []string{"name"}, []string{username})
	if err != nil {
		return err
	}

	uservalue, _ := users.([]dataorm.User)
	if uservalue == nil {
		return errors.New("User not find")
	}
	user := uservalue[0]

	user.Status = ""
	if err = dataorm.Update(user); err != nil {
		return err
	}
	userinrooms, err := dataorm.Query("Userinroom", nil, []string{"Userid"}, []string{strconv.Itoa(user.ID)})
	if err != nil {
		return err
	}

	userinroomvalue, _ := userinrooms.([]dataorm.Userinroom)
	if userinroomvalue == nil {
		return errors.New("Userinroom not find")
	}
	uir := userinroomvalue[0]
	if err = dataorm.Delete("Userinroom", uir.ID); err != nil {
		return err
	}
	return nil
}

//@function: 查询房间内所有用户
//@param: roomname:房间名
//@return: 如果失败，返回错误信息，如果成功，返回用户列表[]string

func (r RoomService) GetAllUser(roomname string) ([]string, error) {

	roomsID, err := getRoomID(roomname)
	if err != nil {
		return nil, err
	}

	userinrooms, err := dataorm.Query("Userinroom", nil, []string{"roomid"}, []string{strconv.Itoa(roomsID)})
	if err != nil {
		return nil, err
	}
	users := []string{}
	value, _ := userinrooms.([]dataorm.Userinroom)
	for _, v := range value {
		name := getUserName(v.Userid)
		if name != "" {
			users = append(users, name)
		}
	}
	return users, nil
}

//@function: 返回所有房间列表
//@return: 如果失败，返回错误信息，如果成功，返回房间列表[]string
func (s RoomService) GetAllRoom() ([]Room, error) {
	values, err := dataorm.Query("Room", nil, nil, nil)
	if err != nil {
		return nil, err
	}
	rooms, _ := values.([]dataorm.Room)
	roomnames := []Room{}
	for _, v := range rooms {
		roomnames = append(roomnames, Room{
			Name: v.Name,
			Id:   strconv.Itoa(v.ID),
		})
	}
	return roomnames, nil
}

//@function: 返回房间名
func (s RoomService) GetRoomName(ID int) string {
	name := getRoomName(ID)
	return name
}

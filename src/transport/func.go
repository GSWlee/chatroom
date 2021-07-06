package transport

import (
	"chatroom/src/service"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

type Roominfo struct {
	Name string `form:"name" json:"name" binding:"required"`
}

var roomsvc = service.RoomService{}

//todo
//增加获取建房者信息的步骤
func Room(c *gin.Context) {
	room := Roominfo{}

	err := c.ShouldBindJSON(&room)
	if err != nil {
		c.String(http.StatusBadRequest, "Invalid input")
	}
	user := "11"
	id, err := roomsvc.CreateRoom(room.Name, user)
	if err != nil {
		c.String(http.StatusBadRequest, "Invalid input")
	}
	c.String(200, strconv.Itoa(id))
}

//todo
//增加获取进入房间者用户的步骤
func RoomEnter(c *gin.Context) {
	id := c.Param("roomid")
	user := "yg s"
	Id, _ := strconv.Atoi(id)
	err := roomsvc.EnterRoom(Id, user)
	if err != nil {
		c.String(400, "Invalid Room ID")
	}
	c.String(200, "Enter the Room")
}

//todo
//增加获取离开房间者用户的步骤
func RoomLeave(c *gin.Context) {
	user := "yg s"
	err := roomsvc.LeaveRoom(user)
	if err != nil {
		log.Println(err)
		c.String(400, "Error")
	}
	c.String(200, "Left the room")
}

//@function: 获取房间名字
func RoomName(c *gin.Context) {
	id := c.Param("roomid")
	Id, _ := strconv.Atoi(id)
	name := roomsvc.GetRoomName(Id)
	if name == "" {
		c.String(400, "Invalid Room ID")
	}
	c.String(200, name)
}

//@function获取房间用户
func RoomUsers(c *gin.Context) {
	id := c.Param("roomid")
	Id, _ := strconv.Atoi(id)
	name := roomsvc.GetRoomName(Id)
	if name == "" {
		c.String(400, "Invalid Room ID")
	}
	userList, err := roomsvc.GetAllUser(name)
	if err != nil {
		c.String(400, "Invalid Room ID")
	}
	type user struct {
		Username string `json:"username"`
	}
	data := []user{}
	for _, v := range userList {
		data = append(data, user{Username: v})
	}
	c.JSON(200, data)
	//data:=[]byte{}
	//for _,v:=range userList{
	//	data=append(data,[]byte(v)...)
	//}
	//c.Data(200,"array",data)
}

//@function 返回房间用户
func RoomList(c *gin.Context) {
	roomlist, err := roomsvc.GetAllRoom()
	if err != nil {
		c.String(400, "Error")
	}
	c.JSON(200, roomlist)
}

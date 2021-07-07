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

type Message struct {
	Id   string `json:"id"`
	Text string `json:"text"`
}

type RoomControlData struct {
	PageIndex int `json:"pageIndex"`
	PageSize  int `json:"pageSize"`
}

type MessageControlData struct {
	PageIndex int `json:"pageIndex"`
	PageSize  int `json:"pageSize"`
}

type User struct {
	Username  string `json:"username"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Phone     string `json:"phone"`
}

func getNewRoomControlData() RoomControlData {
	return RoomControlData{
		PageIndex: 0,
		PageSize:  100,
	}
}

func getNewMessageControlData() MessageControlData {
	return MessageControlData{
		PageIndex: -1,
		PageSize:  100,
	}
}

var roomsvc = service.RoomService{}
var messagesvc = service.MessageService{}
var usersvc = service.Userservice{}

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

//@function 返回房间列表
func RoomList(c *gin.Context) {
	roomcontral := getNewRoomControlData()
	err := c.ShouldBindJSON(&roomcontral)
	if err != nil {
		c.String(400, "Invalid input")
	}
	roomlist, err := roomsvc.GetAllRoom()
	if err != nil {
		c.String(400, "Error")
	}
	length := len(roomlist)
	index := 0
	for ; index < roomcontral.PageIndex && length > roomcontral.PageSize; index, length = index+1, length-roomcontral.PageSize {
	}
	rawdata := roomlist[index*roomcontral.PageSize:]
	if len(rawdata) > roomcontral.PageSize {
		rawdata = rawdata[:roomcontral.PageSize]
	}
	c.JSON(200, rawdata)
}

//todo：增加读取发送者ID与房间Id的步骤
//@function: 房间用户发送信息
func MessageSend(c *gin.Context) {
	message := Message{}
	err := c.ShouldBindJSON(&message)
	if err != nil {
		c.String(400, "Invalid input")
	}

	userid := 1
	roomid := 1
	err = messagesvc.Send(userid, roomid, message.Id, message.Text)
	if err != nil {
		c.String(400, "Invalid input")
	}
	c.String(200, "successful operation")
}

//todo：增加读取发送者ID与房间Id的步骤
//@function: 房间接收消息
func MessageRetrieve(c *gin.Context) {
	messagecontral := getNewMessageControlData()
	err := c.ShouldBindJSON(&messagecontral)
	if err != nil {
		c.String(400, "Invalid input")
	}
	roomid := 1
	messagelist, err := messagesvc.Retrieve(roomid)
	if err != nil {
		c.String(400, "Invalid input")
	}
	length := len(messagelist)
	index := -1
	for ; index > messagecontral.PageIndex && length > messagecontral.PageSize; index, length = index-1, length-messagecontral.PageSize {
	}
	rawdata := messagelist[(-index-1)*messagecontral.PageSize:]
	if len(rawdata) > messagecontral.PageSize {
		rawdata = rawdata[:messagecontral.PageSize]
	}
	type Message struct {
		ID        string `json:"id"`
		Text      string `json:"text"`
		Timestamp string `json:"timestamp"`
	}
	data := []Message{}
	for _, v := range rawdata {
		data = append(data, Message{
			ID:        v.Messageid,
			Text:      v.Data,
			Timestamp: v.Time.Format("2006-01-02 15:04:05"),
		})
	}
	c.JSON(200, data)
}

func CreateUser(c *gin.Context) {
	user := User{}
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.String(400, "wrong info")
	}
	err = usersvc.CreateUser(user.Username, user.Firstname, user.Lastname, user.Phone, user.Email, user.Password)
	if err != nil {
		c.String(400, "wrong info")
	}
	c.String(200, "successful operation")
}

//todo
func Login(c *gin.Context) {

}

func UserInfo(c *gin.Context) {
	username := c.Param("username")
	type user struct {
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Email     string `json:"email"`
		Phone     string `json:"phone"`
	}
	u, err := usersvc.Userinfo(username)
	if err != nil {
		c.String(400, "Invalid username supplied")
	}
	data := user{
		FirstName: u.Firstname,
		LastName:  u.Lastname,
		Email:     u.Email,
		Phone:     u.Phone,
	}
	c.JSON(200, data)
}

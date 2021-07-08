package transport

import (
	"chatroom/src/service"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
	"time"
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

type Client struct {
	Username string `json:"username"`
	UserID   int    `json:"userid"`
	RoomID   int    `json:"roomid"`
	jwt.StandardClaims
}

var sec = []byte("Aajkghjuidbhakjbfghjiv")

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

func authentication(token string) (Client, error) {
	client := Client{}
	getToken, _ := jwt.ParseWithClaims(token, &client, func(token *jwt.Token) (interface{}, error) {
		return sec, nil
	})
	if getToken.Valid {
		return client, nil
	} else {
		return Client{}, errors.New("Not Authentication")
	}
}

func createToken(client Client) string {
	client.ExpiresAt = time.Now().Add(time.Minute * 10).Unix()
	token_obj := jwt.NewWithClaims(jwt.SigningMethodHS256, client)
	token, _ := token_obj.SignedString(sec)
	return token
}

func updateToken(token string) string {
	client, _ := authentication(token)
	client.ExpiresAt = time.Now().Add(time.Minute * 10).Unix()
	token = createToken(client)
	return token
}

//增加获取建房者信息的步骤
func Room(c *gin.Context) {
	room := Roominfo{}
	token := c.GetHeader("authorization")
	client, err := authentication(token)
	if err != nil {
		c.String(400, "Invalid input")
	}
	err = c.ShouldBindJSON(&room)
	if err != nil {
		c.String(http.StatusBadRequest, "Invalid input")
	}
	user := client.Username
	id, err := roomsvc.CreateRoom(room.Name, user)
	if err != nil {
		c.String(http.StatusBadRequest, "Invalid input")
	}
	token = updateToken(token)
	c.Header("authorization", token)
	c.String(200, strconv.Itoa(id))
}

//增加获取进入房间者用户的步骤
func RoomEnter(c *gin.Context) {
	token := c.GetHeader("authorization")
	client, err := authentication(token)
	if err != nil {
		c.String(400, "Invalid Room ID")
	}
	id := c.Param("roomid")
	user := client.Username
	Id, _ := strconv.Atoi(id)
	err = roomsvc.EnterRoom(Id, user)
	if err != nil {
		c.String(400, "Invalid Room ID")
	}
	client.RoomID = Id
	token = createToken(client)
	c.Header("authorization", token)
	c.String(200, "Enter the Room")
}

//增加获取离开房间者用户的步骤
func RoomLeave(c *gin.Context) {
	token := c.GetHeader("authorization")
	client, err := authentication(token)
	if err != nil {
		c.String(400, "Invalid input")
	}
	err = roomsvc.LeaveRoom(client.Username)
	if err != nil {
		log.Println(err)
		c.String(400, "Error")
	}
	client.RoomID = -1
	token = createToken(client)
	c.Header("authorization", token)
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

//@function: 房间用户发送信息
func MessageSend(c *gin.Context) {
	token := c.GetHeader("authorization")
	client, err := authentication(token)
	if err != nil {
		c.String(400, "Invalid input")
	}
	message := Message{}
	err = c.ShouldBindJSON(&message)
	if err != nil {
		c.String(400, "Invalid input")
	}

	userid := client.UserID
	roomid := client.RoomID
	err = messagesvc.Send(userid, roomid, message.Id, message.Text)
	if err != nil {
		c.String(400, "Invalid input")
	}
	token = updateToken(token)
	c.Header("authorization", token)
	c.String(200, "successful operation")
}

//@function: 房间接收消息
func MessageRetrieve(c *gin.Context) {
	token := c.GetHeader("authorization")
	client, err := authentication(token)
	if err != nil {
		c.String(400, "Invalid input")
	}
	roomid := client.RoomID
	messagecontral := getNewMessageControlData()
	err = c.ShouldBindJSON(&messagecontral)
	if err != nil {
		c.String(400, "Invalid input")
	}
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
	token = updateToken(token)
	c.Header("authorization", token)
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

//@function: 用户登陆
func Login(c *gin.Context) {
	username, password := c.Query("username"), c.Query("password")
	user, err := usersvc.Login(username, password)
	if err != nil {
		c.String(400, "Invalid username or password.")
	}
	client := Client{
		Username: user.Name,
		UserID:   user.ID,
		RoomID:   -1,
	}
	token := createToken(client)
	c.Header("authorization", token)
	c.String(200, "token")
}

//@function: 获取用户信息
func UserInfo(c *gin.Context) {
	username := c.Param("username")
	type user struct {
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Email     string `json:"email"`
		Phone     string `json:"phone"`
	}
	u, err := usersvc.UserInfo(username)
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

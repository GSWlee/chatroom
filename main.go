package main

import (
	"chatroom/src/transport"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.POST("/room", transport.Room)
	r.PUT("/room/:roomid/enter", transport.RoomEnter)
	r.PUT("/roomLeave", transport.RoomLeave)
	r.GET("/room/:roomid", transport.RoomName)
	r.GET("/room/:roomid/users", transport.RoomUsers)
	r.POST("/roomList", transport.RoomList)

	r.POST("/message/send", transport.MessageSend)
	r.POST("/message/retrieve", transport.MessageRetrieve)

	r.POST("/user", transport.CreateUser)
	r.GET("/userLogin", transport.Login)
	r.GET("/user/:username", transport.UserInfo)

	r.Run()
}

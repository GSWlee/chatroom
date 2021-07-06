package main

import (
	_ "chatroom/src/dataorm"
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
	r.Run()
}

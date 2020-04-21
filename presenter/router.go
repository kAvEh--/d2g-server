package presenter

import (
	"fmt"
	"github.com/gin-gonic/gin"
	socketio "github.com/googollee/go-socket.io"
	"log"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	gin.SetMode(gin.DebugMode)
	r.MaxMultipartMemory = 50 << 20 // 50 MB

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	//Version 1 APIs
	// APIs without authentication
	v1 := r.Group("/api/v1")
	{
		v1.GET("/login")
	}

	return r
}

var D2GServer *socketio.Server

func SetupSocket() *socketio.Server {
	server, err := socketio.NewServer(nil)
	if err != nil {
		log.Fatal(err)
	}
	D2GServer = server
	server.OnConnect("/", func(s socketio.Conn) error {
		s.SetContext("")
		fmt.Println("connected:", s.ID())
		return nil
	})
	server.OnEvent("/", "sign-in", login)
	server.OnEvent("/", "find-match", findMatch)
	server.OnEvent("/", "ready", ready)
	server.OnEvent("/", "after", after)
	server.OnError("/", func(s socketio.Conn, e error) {
		fmt.Println("meet error:", e)
	})
	server.OnDisconnect("/", func(s socketio.Conn, reason string) {
		fmt.Println("closed", reason)
	})

	return server
}

package presenter

import (
	"bodybuilding/src/lib/common"
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

func SetupSocket() *socketio.Server {
	server, err := socketio.NewServer(nil)
	if err != nil {
		log.Fatal(err)
	}
	server.OnConnect("/", func(s socketio.Conn) error {
		s.SetContext("")
		fmt.Println("connected:", s.ID())
		return nil
	})
	server.OnEvent("/", "sign-in", func(s socketio.Conn, msg common.JSON) {
		fmt.Println("notice:", msg)
		//s.Emit("reply", "have "+msg)
	})
	server.OnError("/", func(s socketio.Conn, e error) {
		fmt.Println("meet error:", e)
	})
	server.OnDisconnect("/", func(s socketio.Conn, reason string) {
		fmt.Println("closed", reason)
	})
	go server.Serve()
	defer server.Close()

	return server
}

package presenter

import (
	"bodybuilding/src/lib/common"
	"fmt"
	socketio "github.com/googollee/go-socket.io"
)

func login(s socketio.Conn, msg common.JSON) {
	fmt.Println("-------->>", msg)
}

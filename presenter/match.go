package presenter

import (
	"bodybuilding/src/lib/common"
	"fmt"
	socketio "github.com/googollee/go-socket.io"
)

func findMatch(s socketio.Conn, msg common.JSON) {
	fmt.Println("Find Match-------->>", msg)
	//var search model.Search
	userId := msg["playerId"].(int)
	level := msg["level"].(int)
	gameType := msg["type"].(int)
	fmt.Println(userId, ":::", level, ":::", gameType)
}

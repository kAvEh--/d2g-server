package presenter

import (
	"bodybuilding/src/lib/common"
	"d2g-server/model"
	"encoding/json"
	"fmt"
	socketio "github.com/googollee/go-socket.io"
	"strconv"
)

func findMatch(s socketio.Conn, msg common.JSON) {
	fmt.Println("Find Match-------->>", msg)
	//var search model.Search
	userId := msg["playerId"].(string)
	lvl, _ := strconv.Atoi(msg["level"].(string))
	gameType, _ := strconv.Atoi(msg["type"].(string))
	fmt.Println(userId, ":::", lvl, ":::", gameType)

	var room model.Room

	room.P1Formation = 1
	room.P2Formation = 2

	D2GServer.JoinRoom("", userId, s)
	fmt.Println(D2GServer.Rooms(""))

	b, err := json.Marshal(room)
	if err != nil {
		fmt.Println(err)
		return
	}
	D2GServer.BroadcastToRoom("", userId, "match-found", string(b))
}

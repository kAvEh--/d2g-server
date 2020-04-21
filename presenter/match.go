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

	fmt.Println("rooms:::: ", D2GServer.Rooms(""))

	b, err := json.Marshal(room)
	if err != nil {
		fmt.Println(err)
		return
	}
	s.Emit("match-found", string(b))
}

func ready(s socketio.Conn, msg common.JSON) {
	userId := msg["playerId"].(string)
	fmt.Println("Player is Ready-------->>", msg, userId)
	var room model.Room

	room.P1Formation = 1
	room.P2Formation = 2
	b, err := json.Marshal(room)
	if err != nil {
		fmt.Println(err)
		return
	}
	s.Emit("startTheFuckinMatch", string(b))
}

func after(s socketio.Conn, msg common.JSON) {
	fmt.Println("Player sends after-------->>", msg)
}

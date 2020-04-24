package presenter

import (
	"bodybuilding/src/lib/common"
	"d2g-server/model"
	"encoding/json"
	"fmt"
	socketio "github.com/googollee/go-socket.io"
	"strconv"
	"strings"
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

var flag bool

func after(s socketio.Conn, msg common.JSON) {
	//fmt.Println("Player sends after-------->>", msg)
	//ball := msg["ballPosition"].([]interface {})
	// convert map to json
	//t1 := `{"ballPosition":["0x0.0p0","0x0.0p0"],"keeperPosition":["0x1.bp2","0x0.0p0"],"playerId":"00000000-0000-0000-0000-000000000000","playersPosition":{"1":["-0x1.0p2","0x1.2p1"],"2":["0x0.0p0","0x1.2p1"],"3":["0x1.0p4","0x1.2p3"],"4":["0x1.0p2","0x1.2p1"],"5":["0x1.0p4","0x1.2p3"]},"playersStamina":{"1":"0x1.0p0","2":"0x1.0p0","3":"0x1.0p0","4":"0x1.0p0","5":"0x1.0p0"}}`
	t2 := `{"bPos":["0x0.0p0","0x0.0p0"],"kPo":["0x1.fecp8","0x0.0p0"],"pId":"10000000-0000-0000-0000-000000000000","pPos":{"1":["-0x1.0p2","-0x1.2p1"],"2":["0x0.0p0","-0x1.2p1"],"3":["0x1.0p4","0x1.2p3"],"4":["0x1.0p2","-0x1.2p1"],"5":["0x1.0p4","0x1.2p3"]},"mSta":{"1":"0x1.0p0","2":"0x1.0p0","3":"0x1.0p0","4":"0x1.0p0","5":"0x1.0p0"}}`
	//fmt.Println(t1)
	//t1 := map[ballPosition:[-0x1.9428f6p2 0x0.0p0] keeperPosition:[0x1.bp2 0x0.0p0] playerId:00000000-0000-0000-0000-000000000000 playersPosition:map[1:[-0x1.9428f6p2 0x1.a147aep1] 2:[-0x1.9428f6p2 0x1.a147aep1] 3:[-0x1.9428f6p2 0x1.a147aep1] 4:[-0x1.9428f6p2 0x1.a147aep1] 5:[-0x1.9428f6p2 0x1.a147aep1]] playersStamina:map[1:0x1.0p0 2:0x1.0p0 3:0x1.0p0 4:0x1.0p0 5:0x1.0p0]]
	jsonString, _ := json.Marshal(msg)
	fmt.Println(string(jsonString))
	tt := strings.Replace(string(jsonString), "pPos", "delete", -1)
	tt = strings.Replace(tt, "p2Pos", "pPos", -1)
	tt = strings.Replace(tt, "00000000-0000-0000-0000-000000000000", "1000000-0000-0000-0000-000000000000", -1)
	fmt.Println(">>>>>>>>", tt)
	if flag {
		s.Emit("after", tt)
	} else {
		flag = true
		s.Emit("after", t2)
	}
}

func before(s socketio.Conn, msg common.JSON) {
	fmt.Println("Player sends before-------->>", msg)
	jsonString, _ := json.Marshal(msg)
	fmt.Println(string(jsonString))
	t1 := `{"direction":{"arrow":["0x1.cp-4","0x1.cc0002p1"],"player":2},"playerId":"10000000-0000-0000-0000-000000000000"}`
	s.Emit("before", t1)
}

func goal(s socketio.Conn, msg common.JSON) {
	fmt.Println("goal-------->>", msg)
}

func matchEnd(s socketio.Conn, msg common.JSON) {
	fmt.Println("end match-------->>", msg)
}

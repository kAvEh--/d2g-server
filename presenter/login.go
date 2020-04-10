package presenter

import (
	"bodybuilding/src/lib/common"
	"d2g-server/model"
	"encoding/json"
	"fmt"
	socketio "github.com/googollee/go-socket.io"
)

func login(s socketio.Conn, msg common.JSON) {
	fmt.Println("-------->>", msg)
	var user model.User
	var player model.Player
	//user.ID = uuid.UUID{}
	user.DailyCoin = true
	player.ID = 1
	player.WinCount = 10
	player.GameCount = 20
	player.WinRate = 50
	player.CleanSheet = 15
	player.Formation = 2
	player.Shirt = 1
	player.Name = "kAvEh"
	player.AvatarID = 3
	player.UserName = "kAvEh"
	player.Goals = 100
	player.WinInRaw = 5
	player.Coin = 2000
	player.Lineup = "B"
	var level model.Level
	level.Level = 6
	level.XP = 500
	player.Level = &level
	user.Player = &player
	//user.Player.Level.Level = 6
	//user.Player.Level.XP = 500

	b, err := json.Marshal(user)
	if err != nil {
		fmt.Println(err)
		return
	}
	s.Emit("loggedInPlayer", string(b))
}

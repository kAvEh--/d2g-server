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
	player.Lineup = "B"
	player.Shirt = 1
	player.Name = "kAvEh"
	player.AvatarID = 3
	player.UserName = "kAvEh"
	player.Goals = 100
	player.WinInRaw = 5
	player.Coin = 2000
	var level model.Level
	level.Level = 6
	level.XP = 500
	player.Level = &level
	var ach model.Achievement
	ach.WinInRaw = 2
	ach.CleanSheet = 1
	player.Achievement = &ach
	var m1 model.Member
	var m2 model.Member
	var m3 model.Member
	var m4 model.Member
	var m5 model.Member
	m1.ShirtNum = 1
	m1.Size = 1
	m1.Speed = 0
	m1.Stamina = 2
	m2.ShirtNum = 2
	m2.Size = 2
	m2.Speed = 0
	m2.Stamina = 2
	m3.ShirtNum = 3
	m3.Size = 1
	m3.Speed = 1
	m3.Stamina = 2
	m4.ShirtNum = 4
	m4.Size = 1
	m4.Speed = 0
	m4.Stamina = 1
	m5.ShirtNum = 5
	m5.Size = 0
	m5.Speed = 0
	m5.Stamina = 2
	var tmp []model.Member
	tmp = append(tmp, m1, m2, m3, m4, m5)
	player.Players = &tmp
	shirts := make([]model.Shirt, 24)
	for i, shirt := range shirts {
		shirt.ShirtID = i
		shirt.HasShirt = i%4 == 0
		shirts[i] = shirt
	}
	player.Shirts = &shirts

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

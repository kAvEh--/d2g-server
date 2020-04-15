package model

type Room struct {
	IDv2
	RoomNum int  `json:"room_num"`
	Player1 User `json:"player_1"`
	Player2 User `json:"player_2"`
}

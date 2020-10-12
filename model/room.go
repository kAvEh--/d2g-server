package model

import uuid "github.com/satori/go.uuid"

type Room struct {
	IDv2
	Turn      uuid.UUID   `json:"turn"`
	Level     int         `json:"level"`
	RoomState *RoomState  `json:"room_state"`
	P1        *Player     `json:"p_1"`
	P2        *Player     `json:"p_2"`
	P1Stat    *PlayerStat `json:"p1_stat"`
	P2Stat    *PlayerStat `json:"p2_stat"`
}

type PlayerStat struct {
	X              float32               `json:"x"`
	Y              float32               `json:"y"`
	IsSync         bool                  `json:"is_sync"`
	PlayerShooting int                   `json:"player_shooting"`
	IsReady        bool                  `json:"is_ready"`
	PlayerPosition []map[string][]string `json:"playersPosition"`
	Stamina        map[string]string     `json:"playersStamina"`
}

type RoomState struct {
	LastTouch      int      `json:"last_touch"`
	RoundNum       int      `json:"round_num"`
	GameState      int      `json:"game_state"`
	IsP1First      bool     `json:"is_p1_first"`
	GoalToWin      int      `json:"goal_to_win"`
	IsMatchReady   bool     `json:"is_match_ready"`
	IsWinner       bool     `json:"is_winner"`
	HalfStatus     int      `json:"half_status"`
	GaolerPosition int      `json:"gaoler_position"`
	GameType       int      `json:"game_type"`
	P1Goals        int      `json:"p1_goals"`
	P2Goals        int      `json:"p2_goals"`
	BallPosition   []string `json:"ballPosition"`
	KeeperPosition []string `json:"keeperPosition"`
}

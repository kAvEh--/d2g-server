package model

import uuid "github.com/satori/go.uuid"

type Room struct {
	IDv2
	RoomNum          int       `json:"room_num"`
	Turn             uuid.UUID `json:"turn"`
	Level            int       `json:"level"`
	P1ID             uuid.UUID `json:"p1ID"`
	P2ID             uuid.UUID `json:"p2ID"`
	P1Name           string    `json:"p1_name"`
	P2Name           string    `json:"p2_name"`
	P1Level          int       `json:"p1_level"`
	P2Level          int       `json:"p2_level"`
	P1Xp             int       `json:"p1_xp"`
	P2Xp             int       `json:"p2_xp"`
	P1Shirt          int       `json:"p1_shirt"`
	P2Shirt          int       `json:"p2_shirt"`
	P1Avatar         int       `json:"p1_avatar"`
	P2Avatar         int       `json:"p2_avatar"`
	P1WinRate        float32   `json:"p1_win_rate"`
	P2WinRate        float32   `json:"p2_win_rate"`
	P1Formation      int       `json:"p1_formation"`
	P2Formation      int       `json:"p2_formation"`
	LastTouch        int       `json:"last_touch"`
	P1Goals          int       `json:"p1_goals"`
	P2Goals          int       `json:"p2_goals"`
	RoundNum         int       `json:"round_num"`
	GameState        int       `json:"game_state"`
	IsP1First        bool      `json:"is_p1_first"`
	GoalToWin        int       `json:"goal_to_win"`
	IsP1Ready        bool      `json:"is_p1_ready"`
	IsP2Ready        bool      `json:"is_p2_ready"`
	P1PlayerShooting int       `json:"p1_player_shooting"`
	P2PlayerShooting int       `json:"p2_player_shooting"`
	P1X              float32   `json:"p1_x"`
	P1Y              float32   `json:"p1_y"`
	P2X              float32   `json:"p2_x"`
	P2Y              float32   `json:"p2_y"`
	P1DataSync       bool      `json:"p1_data_sync"`
	P2DataSync       bool      `json:"p2_data_sync"`
	IsMatchReady     bool      `json:"is_match_ready"`
	IsWinner         bool      `json:"is_winner"`
	HalfStatus       int       `json:"half_status"`
	GaolerPosition   int       `json:"gaoler_position"`
	GameType         int       `json:"game_type"`
}

type GameState struct {
	PlayerID       uuid.UUID             `json:"playerId"`
	BallPosition   []string              `json:"ballPosition"`
	KeeperPosition []string              `json:"keeperPosition"`
	PlayerPosition []map[string][]string `json:"playersPosition"`
	Stamina        map[string]string     `json:"playersStamina"`
}

package model

import (
	uuid "github.com/satori/go.uuid"
	"time"
)

type User struct {
	IDv1
	LastDailyCoin *time.Time `json:"-"`
	DailyCoin     bool       `gorm:"-"json:"dailyCoin"`
	Player        *Player    `json:"player"`
}

type Player struct {
	IDv2
	UserID      uuid.UUID    `json:"user_id"`
	WinCount    int          `json:"winCount"`
	GameCount   int          `json:"gameCount"`
	WinRate     float32      `json:"winRate"`
	CleanSheet  int          `json:"cleanSheet"`
	Formation   int          `json:"formation"`
	Shirt       int          `json:"shirt"`
	Name        string       `json:"nickname"`
	AvatarID    int          `json:"avatarId"`
	UserName    string       `json:"username"`
	Goals       int          `json:"goals"`
	WinInRaw    int          `json:"winInaRow"`
	Coin        int          `json:"coin"`
	Level       *Level       `json:"level"`
	Achievement *Achievement `json:"achievements"`
	Players     *[]Member    `json:"players"`
	Lineup      string       `json:"lineup"`
}

type Level struct {
	IDv2
	PlayerID int `json:"player_id"`
	Level    int `json:"lvl"`
	XP       int `json:"xp"`
}

type Achievement struct {
	IDv2
	PlayerID   int `json:"player_id"`
	Goal       int `json:"goal"`
	CleanSheet int `json:"cleanSheet"`
	Win        int `json:"win"`
	WinInRaw   int `json:"winInaRow"`
}

type Member struct {
	IDv2
	PlayerID int `json:"player_id"`
	Stamina  int `json:"stamina"`
	Size     int `json:"size"`
	Speed    int `json:"speed"`
}

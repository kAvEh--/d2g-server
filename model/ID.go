package model

import (
	uuid "github.com/satori/go.uuid"
	"time"
)

type IDv1 struct {
	ID        uuid.UUID `gorm:"primary_key" sql:"type:uuid;default:uuid_generate_v4()"json:"id"`
	CreatedAt *time.Time
	UpdatedAt *time.Time
	DeletedAt *time.Time `sql:"index"json:"-"`
}

type IDv2 struct {
	ID        int `gorm:"not null;auto_increment;primary_key"json:"id"`
	CreatedAt *time.Time
	UpdatedAt *time.Time
	DeletedAt *time.Time `sql:"index"json:"-"`
}

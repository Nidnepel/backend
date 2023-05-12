package entity

import (
	"time"
)

type Session struct {
	Id        int
	ProjectId int `db:"project_id"`
	UserId    int `db:"user_id"`
	Keylog    string
	Screens   [][]byte
	Start     time.Time
	Finish    time.Time
}

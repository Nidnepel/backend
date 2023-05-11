package entity

import "github.com/jackc/pgx/pgtype"

type Session struct {
	id        int
	projectId int
	userId    int
	keylog    string
	//	screens map[bitmap]bool
	start  pgtype.Timestamp
	finish pgtype.Timestamp
}

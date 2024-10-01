// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package db

import (
	"database/sql"
)

type ChatRoom struct {
	ID                string
	Name              string
	LastMessageSentAt sql.NullString
	CreatedAt         string
	UpdatedAt         string
}

type Message struct {
	ID         int64
	ChatRoomID string
	UserID     int64
	Username   string
	Message    string
	CreatedAt  string
	UpdatedAt  string
}

type User struct {
	ID        int64
	Name      string
	CreatedAt string
	UpdatedAt string
	SessionID string
}

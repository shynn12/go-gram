package models

import "time"

type MessageDTO struct {
	UserID int       `json:"user_id"`
	Body   string    `json:"body"`
	ChatID int       `json:"chat_id"`
	Time   time.Time `json:"time"`
}

type Error struct {
	Text string `json:"err_text"`
}

package models

import "time"

type Status string

const (
	Pendign Status = "pending"
	Sent    Status = "sent"
	Failed  Status = "failed"
)

type Notification struct {
	ID        string
	UserID    int64
	Channel   string
	Message   string
	Status    Status
	CreatedAt time.Time
	sentAt    *time.Time
}

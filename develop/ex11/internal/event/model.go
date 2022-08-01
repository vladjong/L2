package event

import "time"

type Event struct {
	ID   string    `json:"user_id"`
	Date time.Time `json:"date"`
}

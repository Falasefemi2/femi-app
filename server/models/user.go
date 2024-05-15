package models

import "time"

type User struct {
	ID        int       `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	CreateAt  time.Time `json:"create_at"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type Channel struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	CreatedBy int       `json:"createdBy"` // User ID of the creator
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type Message struct {
	ID        int       `json:"id"`
	Content   string    `json:"content"`
	ChannelID int       `json:"channelId"`
	SentBy    int       `json:"sentBy"` // User ID of the sender
	SentAt    time.Time `json:"sentAt"`
}

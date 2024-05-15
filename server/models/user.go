package models

import (
	"fmt"
	"time"

	"github.com/falasefemi2/chat-app/db"
)

type User struct {
	ID        int       `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
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

func CreateUser(user *User) error {
	query := `
	INSERT INTO users (username, email, password)
	VALUES (?, ?, ?)
	`

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return fmt.Errorf("failed to prepare statement: %v", err)
	}

	defer stmt.Close()

	result, err := stmt.Exec(user.Username, user.Email, user.Password)

	if err != nil {
		return fmt.Errorf("failed to execute statement: %w", err)
	}

	id, err := result.LastInsertId()

	if err != nil {
		return fmt.Errorf("failed to get last insert id: %w", err)
	}

	user.ID = int(id)
	return nil
}

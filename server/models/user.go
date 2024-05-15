package models

import (
	"fmt"
	"time"

	"github.com/falasefemi2/chat-app/db"
	"github.com/falasefemi2/chat-app/utils"
)

type User struct {
	ID        int       `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"create_at"`
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
	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return fmt.Errorf("failed to hash password: %w", err)
	}

	query := `
        INSERT INTO users (username, email, password)
        VALUES (?, ?, ?)
    `

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return fmt.Errorf("failed to prepare statement: %v", err)
	}
	defer stmt.Close()

	result, err := stmt.Exec(user.Username, user.Email, hashedPassword)
	if err != nil {
		return fmt.Errorf("failed to execute statement: %w", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return fmt.Errorf("failed to get last insert id: %w", err)
	}

	user.ID = int(id)
	user.Password = hashedPassword // Store the hashed password

	// Retrieve the created_at and updated_at values from the database
	rows, err := db.DB.Query("SELECT created_at, updated_at FROM users WHERE id = ?", id)
	if err != nil {
		return fmt.Errorf("failed to retrieve timestamp values: %w", err)
	}
	defer rows.Close()

	if rows.Next() {
		var createdAt, updatedAt time.Time
		if err := rows.Scan(&createdAt, &updatedAt); err != nil {
			return fmt.Errorf("failed to scan timestamp values: %w", err)
		}
		user.CreatedAt = createdAt
		user.UpdatedAt = updatedAt
	}

	return nil
}

func DeleteAllUsers() error {
	query := "DELETE FROM users"

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return fmt.Errorf("failed to prepare statement: %v", err)
	}
	defer stmt.Close()

	_, err = stmt.Exec()
	if err != nil {
		return fmt.Errorf("failed to execute statement: %w", err)
	}

	return nil
}

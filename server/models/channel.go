package models

import (
	"fmt"
	"time"

	"github.com/falasefemi2/chat-app/db"
)

func CreateChannel(name string, createdBy int) (*Channel, error) {
	query := `
	INSERT INTO channels (name, created_by, created_at, updated_at)
	VALUES (?,?,?,?)
	`

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return nil, fmt.Errorf("failed to prepare statement: %w", err)
	}
	defer stmt.Close()

	now := time.Now()

	result, err := stmt.Exec(name, createdBy, now, now)

	if err != nil {
		return nil, fmt.Errorf("failed to get last insert ID: %w", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, fmt.Errorf("failed to get last insert ID: %w", err)
	}

	channel := &Channel{
		ID:        int(id),
		Name:      name,
		CreatedBy: createdBy,
		CreatedAt: now,
		UpdatedAt: now,
	}

	return channel, nil

}

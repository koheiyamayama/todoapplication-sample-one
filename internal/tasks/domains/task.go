package tasks

import (
	"fmt"
	"time"
	"todoapplication-sample-one/bff/internal/libs"
	users "todoapplication-sample-one/bff/internal/users/domains"
)

type (
	task struct {
		id        ID
		title     Title
		body      Body
		userID    UserID
		createdAt CreatedAt
	}

	Task interface{}

	ID        string
	Title     string
	Body      string
	UserID    users.ID
	CreatedAt time.Time
)

func NewTask(title Title, body Body, userID UserID) (Task, error) {
	return newTask(title, body, userID)
}

func newTask(title Title, body Body, userID UserID) (Task, error) {
	id, err := libs.GenerateUserID(20)
	if err != nil {
		return nil, fmt.Errorf("domains.task: %w", err)
	}

	return task{
		id:        ID(id),
		title:     title,
		body:      body,
		userID:    userID,
		createdAt: CreatedAt(time.Now()),
	}, nil
}

package users

import (
	"errors"
	"fmt"
	"todoapplication-sample-one/bff/internal/libs"
)

type (
	user struct {
		id   ID
		name Name
	}

	ID   string
	Name string

	User interface {
	}
)

func NewUser(name Name) (User, error) {
	id, err := libs.GenerateUserID(20)

	if err != nil {
		return nil, err
	}

	return newUser(ID(id), name)
}

func newUser(id ID, name Name) (User, error) {
	if _, err := id.Valid(); err != nil {
		return nil, fmt.Errorf("domain.newUser: failed to create user: %w", err)
	}
	return user{
		id:   id,
		name: name,
	}, nil
}

func (id ID) Valid() (bool, error) {
	ret := len(id) == 20

	if ret {
		return ret, nil
	} else {
		return ret, errors.New("ID is invalid")
	}
}

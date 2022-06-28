package users

import (
	"crypto/rand"
	"errors"
	"fmt"
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
	id, err := generateUserID(20)

	if err != nil {
		return nil, err
	}

	return newUser(id, name)
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

func generateUserID(digit uint32) (ID, error) {
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	b := make([]byte, digit)
	if _, err := rand.Read(b); err != nil {
		return ID(""), errors.New("unexpected error")
	}

	var result string
	for _, v := range b {
		result += string(letters[int(v)%len(letters)])
	}

	return ID(result), nil
}

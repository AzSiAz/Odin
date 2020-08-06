// Package user contains everythink related to user (struct, services, etc)
package user

import (
	"fmt"
	"time"

	gonanoid "github.com/matoous/go-nanoid"
)

type User struct {
	ID        uint64    `json:"id" storm:"id,increment"`
	CreatedAt time.Time `json:"createdAt" storm:"index"`
	UpdatedAt time.Time `json:"updatedAt" storm:"index"`
	Username  string    `json:"username" storm:"unique"`
	Email     string    `json:"email" storm:"unique"`
	Password  string    `json:"-"`
}

func NewUser(username, email, hash string) (*User, error) {
	id, err := gonanoid.Nanoid(4)
	if err != nil {
		return nil, err
	}

	return &User{
		Username:  fmt.Sprintf("%s#%s", username, id),
		Email:     email,
		Password:  hash,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}, err
}

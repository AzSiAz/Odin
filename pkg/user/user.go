// Package user contains everythink related to user (struct, services, etc)
package user

import "time"

type UserLogin struct {
	UserID    string
	LoginDate time.Time
	Env       string
	Type      string
}

package user

import (
	"odin/pkg/database"

	"github.com/sirupsen/logrus"
)

var log = logrus.New()

type UserService struct {
	db database.Database
}

func NewUserService(db database.Database) *UserService {
	return &UserService{db}
}

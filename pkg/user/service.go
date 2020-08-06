package user

import (
	"fmt"
	"odin/pkg/database"

	"github.com/asdine/storm"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

var log = logrus.New()

type UserService struct {
	db *database.StormDB
}

func NewUserService(db *database.StormDB) *UserService {
	db.DB.Init(&User{})

	return &UserService{db}
}

func (us *UserService) hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func (us *UserService) CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func (us *UserService) CreateUser(username, email, password string) (*User, error) {
	log.WithFields(logrus.Fields{
		"email":    email,
		"username": username,
	}).Info("trying to create user")

	db := us.db.DB

	if _, err := us.GetUserByEmail(email); err != storm.ErrNotFound {
		return nil, fmt.Errorf("user already exist or: %w", err)
	}

	hash, err := us.hashPassword(password)
	if err != nil {
		return nil, err
	}

	user, err := NewUser(username, email, hash)
	if err != nil {
		return nil, err
	}

	if err = db.Save(user); err != nil {
		return nil, err
	}

	return user, nil
}

func (us *UserService) GetUserByEmail(email string) (*User, error) {
	log.WithField("email", email).Info("trying to get user from email")
	db := us.db.DB

	var user User
	err := db.One("Email", email, &user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (us *UserService) GetUserByUsername(username string) (*User, error) {
	log.WithField("username", username).Info("trying to get user from username")
	db := us.db.DB

	var user User
	err := db.One("Username", username, &user)
	if err != nil {
		log.WithField("username", username).Error(err)
		return nil, err
	}

	return &user, nil
}

func (us *UserService) GetUserByID(ID uint64) (*User, error) {
	log.WithField("id", ID).Info("trying to get user from id")
	db := us.db.DB

	var user User
	err := db.One("ID", ID, &user)
	if err != nil {
		log.WithField("id", ID).Error(err)
		return nil, err
	}

	return &user, nil
}

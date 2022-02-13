package model

import (
	"errors"
  "gorm.io/gorm"
)

// User Struct of an user
type User struct {
	gorm.Model
	ID   int    `json:"id" gorm:"primaryKey"`
	Name string `json:"name"`
}

// User Struct of users
type Users []User

// NewUser Constructor of an user
func NewUser(name string) (*User, error) {
	if name == "" {
		return nil, errors.New("名前は必須です。")
	}

	user := &User{
		Name: name,
	}

	return user, nil
}

// SetUser Setter of an User
// This function used to update the user.
func (user *User) SetUser(name string) error {
	if name == "" {
		return errors.New("名前は必須です。")
	}

	user.Name = name

	return nil
}

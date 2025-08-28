package dto

import (
	"errors"
	"strings"
)

type UserDto struct {
	Id       uint   `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (user *UserDto) validateRequired() error {
	if user.Name == "" {
		return errors.New("name is required")
	}

	if user.Email == "" {
		return errors.New("email is required")
	}

	if user.Password == "" {
		return errors.New("password is required")
	}

	return nil
}

func (user *UserDto) format() {
	user.Name = strings.TrimSpace(user.Name)
	user.Email = strings.TrimSpace(user.Email)
}

func (user *UserDto) Prepare() error {
	if err := user.validateRequired(); err != nil {
		return err
	}

	user.format()
	return nil
}

type UserUpdateDto struct {
	Id    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func (user *UserUpdateDto) validateRequired() error {
	if user.Name == "" {
		return errors.New("name is required")
	}

	if user.Email == "" {
		return errors.New("email is required")
	}

	return nil
}

func (user *UserUpdateDto) format() {
	user.Name = strings.TrimSpace(user.Name)
	user.Email = strings.TrimSpace(user.Email)
}

func (user *UserUpdateDto) Prepare() error {
	if err := user.validateRequired(); err != nil {
		return err
	}

	user.format()
	return nil
}

package models

import (
	dto "main/src/dtos"
	"main/src/utils"
	"time"
)

type User struct {
	ID        int       `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Slug      string    `json:"slug,omitempty"`
	Email     string    `json:"email,omitempty"`
	Password  string    `json:"password,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
}

func (userModel *User) ConvertFromDto(userDto dto.UserDto) {
	userModel.Name = userDto.Name
	userModel.Email = userDto.Email
	userModel.Password = userDto.Password
	userModel.Slug = utils.Slugify(userDto.Name)
	userModel.CreatedAt = time.Now()
}

package models

import (
	"log"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"mai_lab/rpc"
)

type CreateUserDTO struct {
	Name     string `json:"name"`
	Email    string `json:"email,omitempty"`
	Mobile   string `json:"mobile"`
	Password string `json:"password"`
}

type UpdateUserDTO struct {
	ID          uuid.UUID `json:"uuid,omitempty"`
	Name        string    `json:"name,omitempty"`
	Email       string    `json:"email,omitempty"`
	Mobile      string    `json:"mobile,omitempty"`
	Password    string    `json:"password,omitempty"`
	OldPassword string    `json:"old_password,omitempty"`
	NewPassword string    `json:"new_password,omitempty"`
}

func (dto *CreateUserDTO) NewUser() User {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(dto.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatalln(err)
	}
	return User{
		Name:         dto.Name,
		Email:        dto.Email,
		Mobile:       dto.Mobile,
		PasswordHash: string(hashedPassword),
	}
}

func (dto *UpdateUserDTO) UpdateUser(u *User) {

	if len(dto.Name) > 0 {
		u.Name = dto.Name
	}
	if len(dto.Email) > 0 {
		u.Email = dto.Email
	}
	if len(dto.Mobile) > 0 {
		u.Mobile = dto.Mobile
	}
	if len(dto.NewPassword) > 0 {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(dto.Password), bcrypt.DefaultCost)
		if err != nil {
			log.Fatalln(err)
		}
		u.PasswordHash = string(hashedPassword)
	}
}

func TwirpFromUser(user *User) *rpc.User {
	return &rpc.User{
		Id:       user.ID.String(),
		Name:     user.Name,
		Email:    user.Email,
		Mobile:   user.Mobile,
		Password: "",
	}
}

func TwirpFromUsers(user []User) []*rpc.User {
	var users []*rpc.User
	for _, u := range user {
		users = append(users, TwirpFromUser(&u))
	}
	return users
}

package model

import (
	"github.com/google/uuid"
	"mai_lab/pkg"
)

type CreateUserDTO struct {
	Name     string `json:"name"`
	Email    string `json:"email" ,omitempty`
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
	return User{
		ID:           uuid.New(),
		Name:         dto.Name,
		Email:        dto.Email,
		Mobile:       dto.Mobile,
		PasswordHash: pkg.GeneratePasswordHash(dto.Password),
	}
}

func (dto *UpdateUserDTO) UpdateUser() User {
	return User{
		ID:           dto.ID,
		Name:         dto.Name,
		Email:        dto.Email,
		Mobile:       dto.Mobile,
		PasswordHash: dto.Password,
	}
}

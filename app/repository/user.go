package repository

import (
	"fmt"
	"github.com/google/uuid"
)

type User struct {
	ID   uuid.UUID
	Name string
}

type UserStore interface { //separating the layer from the business logic
	Create(u User) (*uuid.UUID, error)
	Read(uid uuid.UUID) (*User, error)
}

type Users struct {
	ustore UserStore
}

// initialization function(first)
func NewUsers(ustore UserStore) *Users {
	return &Users{
		ustore: ustore,
	}
}

func (us *Users) Create(u User) (*User, error) {
	id, err := us.ustore.Create(u)
	if err != nil {
		return nil, fmt.Errorf("create user error: %w", err)
	}
	u.ID = *id //if there is no error, we have a valid pointer. We return a pointer to the user.
	return &u, nil
}

// reading a profile by a unique identifier
func (us *Users) Read(uid uuid.UUID) (*User, error) {
	u, err := us.ustore.Read(uid)
	if err != nil {
		return nil, fmt.Errorf("read user error: %w", err)
	}
	return u, nil
}

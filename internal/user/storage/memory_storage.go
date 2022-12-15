package storage

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"log"
	"mai_lab/internal/user/model"
)

// an alias that can be used anywhere in place of a type name
type usersMap map[uuid.UUID]model.User

func NewStorage() Storage {
	return &usersMap{}
}

// the function of verifying the existence of a user with this key
func (u usersMap) checkUserExist(key uuid.UUID) bool {
	if _, ok := u[key]; ok {
		return true
	}
	return false
}

func (u usersMap) Create(ctx context.Context, user model.User) error {
	if !u.checkUserExist(user.ID) {
		u[user.ID] = user
		log.Println("create user ", user.Name)
		return nil
	}
	return fmt.Errorf("failed to create user")
}

func (u usersMap) GetUser(ctx context.Context, key string) (model.User, error) {
	id, err := uuid.Parse(key)
	if err != nil {
		return model.User{}, err
	}

	if exist := u.checkUserExist(id); exist {
		return u[id], nil
	} else {
		return model.User{}, fmt.Errorf("failed to find user by id: %s", id)
	}
}

func (u usersMap) GetAll(ctx context.Context) ([]model.User, error) {
	all := make([]model.User, 0, len(u))
	for _, value := range u {
		all = append(all, value)
	}
	return all, nil
}

func (u usersMap) Update(ctx context.Context, user model.User) error {
	if u.checkUserExist(user.ID) {
		u[user.ID] = user
		return nil
	}
	return fmt.Errorf("failed to update, user does not exist")
}

func (u usersMap) Delete(ctx context.Context, id uuid.UUID) error {
	if u.checkUserExist(id) {
		delete(u, id)
		return nil
	}
	return fmt.Errorf("failed to delete, user does not exist")
}

package repository

import (
	"errors"
	"go_services_lab/models"
	"strconv"

	"github.com/patrickmn/go-cache"
)

type UserCache struct {
	c *cache.Cache
}

func NewUserCache(c *cache.Cache) *UserCache {
	return &UserCache{c: c}
}

func (r *UserCache) getCount() (int, error) {
	ret, f := r.c.Get("countUser")
	if !f {
		return 0, errors.New("Unable to get number of users.")
	}
	return ret.(int), nil
}

func (r *UserCache) getByLogin(login string) error {
	count, f := r.getCount()
	if f != nil {
		return f
	}
	for i := 1; i <= count; i++ {
		user, fl := r.c.Get("user" + strconv.Itoa(i))
		if fl {
			if login == user.(*models.User).Login {
				return errors.New("User with this login exist.")
			}
		}
	}
	return nil
}

func (r *UserCache) Get(id int) (models.User, error) {
	user, fl := r.c.Get("user" + strconv.Itoa(id))
	if !fl {
		return models.User{}, errors.New("Unable to get user.")
	}
	return models.User{user.(*models.User).ID, user.(*models.User).Name, user.(*models.User).Login, user.(*models.User).Password}, nil
}

func (r *UserCache) Create(user models.User) (int, error) {
	curr_id, f := r.getCount()
	if f != nil {
		return 0, f
	}
	curr_id += 1
	fl := r.getByLogin(user.Login)
	if fl != nil {
		return 0, fl
	}
	r.c.Set("user"+strconv.Itoa(curr_id), &models.User{curr_id, user.Name, user.Login, user.Password}, cache.DefaultExpiration)
	r.c.Increment("countUser", 1)
	return curr_id, nil
}

func (r *UserCache) GetAll() ([]models.User, error) {
	var retList []models.User
	curr_id, f := r.getCount()
	if f != nil {
		return retList, f
	}
	for i := 1; i <= curr_id; i++ {
		user, f := r.c.Get("user" + strconv.Itoa(i))
		if f {
			retList = append(retList, models.User{user.(*models.User).ID, user.(*models.User).Name, user.(*models.User).Login, user.(*models.User).Password})
		}
	}
	return retList, nil
}

func (r *UserCache) Delete(id int) (int, error) {
	r.c.Delete("user" + strconv.Itoa(id))
	return id, nil
}

package repository

import (
	"go_services_lab/models"
	"testing"

	"github.com/patrickmn/go-cache"
)

func TestRepositoryUserCache(t *testing.T) {
	t.Run("if there are no users, then array is empty", func(t *testing.T) {
		cache := NewUserCache(cache.New(cache.NoExpiration, cache.NoExpiration))
		result, _ := cache.GetAll()

		if len(result) != 0 {
			t.Error("failed: array is not empty")
		}
	})

	t.Run("if there are users, array will be returned", func(t *testing.T) {
		c := cache.New(cache.NoExpiration, cache.NoExpiration)
		c.Set("countUser", 1, cache.NoExpiration)
		c.Set("user1", &models.User{ID: 1, Name: "test", Login: "test", Password: "test"}, cache.NoExpiration)
		cache := NewUserCache(c)

		result, _ := cache.GetAll()

		if len(result) == 0 {
			t.Error("failed: array is empty")
		}

		if result[0].ID != 1 || result[0].Login != "test" || result[0].Name != "test" || result[0].Password != "test" {
			t.Error("failed: incorrect user data")
		}
	})

	t.Run("if there are no users, then id == 0", func(t *testing.T) {
		cache := NewUserCache(cache.New(cache.NoExpiration, cache.NoExpiration))
		result, _ := cache.getCount()

		if result != 0 {
			t.Errorf("failed: incorrect result %d", result)
		}
	})

	t.Run("if there is a user login, an error will be returned", func(t *testing.T) {
		c := cache.New(cache.NoExpiration, cache.NoExpiration)
		c.Set("countUser", 1, cache.NoExpiration)
		c.Set("user1", &models.User{ID: 1, Name: "test1", Login: "test2", Password: "test3"}, cache.NoExpiration)
		cache := NewUserCache(c)

		result := cache.getByLogin("test2")

		if result == nil {
			t.Error("failed: incorrect result")
		}
	})

	t.Run("if there is no user login, then nil will be returned", func(t *testing.T) {
		c := cache.New(cache.NoExpiration, cache.NoExpiration)
		c.Set("countUser", 1, cache.NoExpiration)
		cache := NewUserCache(c)

		result := cache.getByLogin("test2")

		if result != nil {
			t.Error("failed: incorrect result")
		}
	})

	t.Run("if you create a user, his id will be returned", func(t *testing.T) {
		c := cache.New(cache.NoExpiration, cache.NoExpiration)
		c.Set("countUser", 0, cache.NoExpiration)
		cache := NewUserCache(c)

		result, _ := cache.Create(models.User{ID: -1, Name: "test", Login: "test", Password: "test"})

		if result != 1 {
			t.Error("failed: incorrect result")
		}
	})

	t.Run("if the 'countUser' field does not exist, the user will not be created", func(t *testing.T) {
		cache := NewUserCache(cache.New(cache.NoExpiration, cache.NoExpiration))
		result, _ := cache.Create(models.User{ID: -1, Name: "test", Login: "test", Password: "test"})

		if result != 0 {
			t.Error("failed: incorrect result")
		}
	})

	t.Run("if you delete a user, his id will be returned", func(t *testing.T) {
		c := cache.New(cache.NoExpiration, cache.NoExpiration)
		c.Set("countUser", 3, cache.NoExpiration)
		c.Set("user1", &models.User{ID: 1, Name: "test", Login: "test1", Password: "test"}, cache.NoExpiration)
		c.Set("user2", &models.User{ID: 2, Name: "test", Login: "test2", Password: "test"}, cache.NoExpiration)
		c.Set("user3", &models.User{ID: 3, Name: "test", Login: "test3", Password: "test"}, cache.NoExpiration)
		cache := NewUserCache(c)

		result, _ := cache.Delete(2)

		if result != 2 {
			t.Errorf("failed: incorrect result \"%d\"", result)
		}
	})
}

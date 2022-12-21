package repository

import (
	"go_services_lab/models"
	"testing"

	"github.com/patrickmn/go-cache"
)

func TestRepositoryProductCache(t *testing.T) {
	t.Run("if there are no products, then the array is empty", func(t *testing.T) {
		cache := NewProductCache(cache.New(cache.NoExpiration, cache.NoExpiration))
		result, _ := cache.GetAll()

		if len(result) != 0 {
			t.Error("failed: array is not empty")
		}
	})

	t.Run("if there are products, an array will be returned", func(t *testing.T) {
		c := cache.New(cache.NoExpiration, cache.NoExpiration)
		c.Set("countProduct", 1, cache.NoExpiration)
		c.Set("product1", &models.Product{ID: 1, Name: "test", Price: 111}, cache.NoExpiration)
		cache := NewProductCache(c)

		result, _ := cache.GetAll()

		if len(result) == 0 {
			t.Error("failed: array is empty")
		}
	})

	t.Run("if there are no products, then id == 0", func(t *testing.T) {
		cache := NewProductCache(cache.New(cache.NoExpiration, cache.NoExpiration))
		result, _ := cache.getCount()

		if result != 0 {
			t.Errorf("failed: incorrect result %d", result)
		}
	})

	t.Run("if there is a product name, an error will be returned", func(t *testing.T) {
		c := cache.New(cache.NoExpiration, cache.NoExpiration)
		c.Set("countProduct", 1, cache.NoExpiration)
		c.Set("product1", &models.Product{ID: 1, Name: "test2", Price: 111}, cache.NoExpiration)
		cache := NewProductCache(c)

		result := cache.getByName("test2")

		if result == nil {
			t.Error("failed: incorrect result")
		}
	})

	t.Run("if there is no product name, then nil will be returned", func(t *testing.T) {
		c := cache.New(cache.NoExpiration, cache.NoExpiration)
		c.Set("countProduct", 1, cache.NoExpiration)
		cache := NewProductCache(c)

		result := cache.getByName("test2")

		if result != nil {
			t.Error("failed: incorrect result")
		}
	})

	t.Run("if there are products, then the last one will return", func(t *testing.T) {
		c := cache.New(cache.NoExpiration, cache.NoExpiration)
		c.Set("countProduct", 2, cache.NoExpiration)
		c.Set("product1", &models.Product{ID: 1, Name: "test1", Price: 111}, cache.NoExpiration)
		c.Set("product2", &models.Product{ID: 2, Name: "test2", Price: 222}, cache.NoExpiration)
		cache := NewProductCache(c)

		result, _ := cache.LastOne()

		if result.ID != 2 {
			t.Error("failed: incorrect result")
		}
	})
}

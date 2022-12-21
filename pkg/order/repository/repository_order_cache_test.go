package repository

import (
	"go_services_lab/models"
	"testing"

	"github.com/patrickmn/go-cache"
)

func TestRepositoryOrderCache(t *testing.T) {
	t.Run("if there are no orders, then the array is empty", func(t *testing.T) {
		cache := NewOrderCache(cache.New(cache.NoExpiration, cache.NoExpiration))
		result, _ := cache.GetAll()

		if len(result) != 0 {
			t.Error("failed: array is not empty")
		}
	})

	t.Run("if there are orders, an array will be returned", func(t *testing.T) {
		c := cache.New(cache.NoExpiration, cache.NoExpiration)
		c.Set("countOrder", 1, cache.NoExpiration)
		c.Set("order1", &models.Order{ID: 1, UserID: 1, Store: []models.Store{{Count: 100, Product: models.Product{ID: 1, Name: "kek", Price: 77}}}}, cache.NoExpiration)
		cache := NewOrderCache(c)

		result, _ := cache.GetAll()

		if len(result) == 0 {
			t.Error("failed: array is empty")
		}
	})

	t.Run("if there are no orders, then id == 0", func(t *testing.T) {
		cache := NewOrderCache(cache.New(cache.NoExpiration, cache.NoExpiration))
		result, _ := cache.getCount()

		if result != 0 {
			t.Errorf("failed: incorrect result %d", result)
		}
	})

	t.Run("if there are products, it will return the total cost of the order", func(t *testing.T) {
		c := cache.New(cache.NoExpiration, cache.NoExpiration)
		c.Set("countOrder", 1, cache.NoExpiration)
		c.Set("order1", &models.Order{ID: 1, UserID: 1, Store: []models.Store{}}, cache.NoExpiration)
		cache := NewOrderCache(c)

		result, _ := cache.Amount(1)

		if result != 0 {
			t.Errorf("failed: incorrect result %f", result)
		}
	})

	t.Run("if you create a order, his id will be returned", func(t *testing.T) {
		c := cache.New(cache.NoExpiration, cache.NoExpiration)
		c.Set("countOrder", 0, cache.NoExpiration)
		c.Set("countProduct", 0, cache.NoExpiration)
		c.Set("product1", &models.Product{ID: 1, Name: "kek", Price: 100}, cache.NoExpiration)

		cache := NewOrderCache(c)

		result, _ := cache.Create(2, map[int]int{1: 1})

		if result != 1 {
			t.Errorf("failed: incorrect result %d", result)
		}
	})

	t.Run("if you delete a products, his id will be returned", func(t *testing.T) {
		c := cache.New(cache.NoExpiration, cache.NoExpiration)
		c.Set("countOrder", 1, cache.NoExpiration)
		c.Set("order1", &models.Order{ID: 1, UserID: 1, Store: []models.Store{{Count: 100, Product: models.Product{ID: 1, Name: "kek", Price: 77}}}}, cache.NoExpiration)
		cache := NewOrderCache(c)

		result, _ := cache.Delete(1)

		if result != 1 {
			t.Errorf("failed: incorrect result \"%d\"", result)
		}
	})
}

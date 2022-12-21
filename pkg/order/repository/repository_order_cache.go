package repository

import (
	"errors"
	"go_services_lab/models"
	"strconv"

	"github.com/patrickmn/go-cache"
)

type OrderCache struct {
	c *cache.Cache
}

func NewOrderCache(c *cache.Cache) *OrderCache {
	return &OrderCache{c: c}
}

func (r *OrderCache) getCount() (int, error) {
	ret, f := r.c.Get("countOrder")
	if !f {
		return 0, errors.New("Unable to get number of orders")
	}
	return ret.(int), nil
}

func (r *OrderCache) Delete(id int) (int, error) {
	r.c.Delete("order" + strconv.Itoa(id))
	return id, nil
}

func (r *OrderCache) Amount(id int) (float32, error) {
	var amount float32
	order, fl := r.c.Get("order" + strconv.Itoa(id))
	if !fl {
		return 0, errors.New("Unable to get order.")
	}
	foo := order.(*models.Order)
	for _, pos := range foo.Store {
		product, _ := r.c.Get("product" + strconv.Itoa(pos.ID))
		amount += product.(*models.Product).Price * float32(pos.Count)
	}
	return amount, nil
}

func (r *OrderCache) Get(id int) (models.Order, error) {
	order, fl := r.c.Get("order" + strconv.Itoa(id))
	if !fl {
		return models.Order{}, errors.New("Unable to get order.")
	}
	return models.Order{order.(*models.Order).ID, order.(*models.Order).UserID, order.(*models.Order).Store}, nil
}

func (r *OrderCache) GetAll() ([]models.Order, error) {
	var retList []models.Order
	curr_id, f := r.getCount()
	if f != nil {
		return retList, f
	}
	for i := 1; i <= curr_id; i++ {
		order, f := r.c.Get("order" + strconv.Itoa(i))
		if f {
			retList = append(retList, models.Order{order.(*models.Order).ID, order.(*models.Order).UserID, order.(*models.Order).Store})
		}
	}
	return retList, nil
}

func (r *OrderCache) Create(user_id int, products map[int]int) (int, error) {
	var store models.Stores
	curr_id, f := r.getCount()
	curr_id += 1
	if f != nil {
		return 0, f
	}
	for key, val := range products {
		product, f := r.c.Get("product" + strconv.Itoa(key))
		if !f {
			return key, errors.New("This product doesn't exist.")
		}
		store = append(store, models.Store{models.Product{product.(*models.Product).ID, product.(*models.Product).Name, product.(*models.Product).Price}, val})
	}
	r.c.Set("order"+strconv.Itoa(curr_id), &models.Order{curr_id, user_id, store}, cache.DefaultExpiration)
	r.c.Increment("countOrder", 1)
	return curr_id, nil
}

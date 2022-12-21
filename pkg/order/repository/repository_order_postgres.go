package repository

import (
	"go_services_lab/models"

	"github.com/jmoiron/sqlx"
)

type BdOrderAnswer struct {
	Id      int `db:"id"`
	User_id int `db:"user_id"`
}

type OrderPostgres struct {
	db *sqlx.DB
}

func NewOrderPostgres(db *sqlx.DB) *OrderPostgres {
	return &OrderPostgres{db: db}
}

func (r *OrderPostgres) Delete(id int) (int, error) {
	_, err := r.db.Exec("DELETE FROM orders WHERE id = $1", id)
	return id, err
}

func (r *OrderPostgres) Get(id int) (order models.Order, err error) {
	var store models.Store
	var stores models.Stores
	var user_id int

	rows, err := r.db.Queryx("SELECT p.id, p.name, p.price, s.count FROM store s JOIN products p ON s.product_id = p.id WHERE s.order_id=$1", id)
	for rows.Next() {
		err = rows.StructScan(&store)
		if err == nil {
			stores = append(stores, store)
		}
	}

	if len(stores) != 0 {
		err = r.db.Get(&user_id, "SELECT user_id from orders WHERE id = $1", id)
		return models.Order{id, user_id, stores}, nil
	}

	return
}

func (r *OrderPostgres) Amount(id int) (amount float32, err error) {
	err = r.db.Get(&amount, "SELECT SUM(p.price*s.count) FROM products p JOIN store s ON p.id = s.product_id WHERE s.order_id = $1", id)
	return
}

func (r *OrderPostgres) GetAll() (retList []models.Order, err error) {
	var order BdOrderAnswer

	rows, err := r.db.Queryx("SELECT * FROM orders")
	for rows.Next() {
		err = rows.StructScan(&order)
		if err == nil {
			retOne, _ := r.Get(order.Id)
			retList = append(retList, retOne)
		}
	}

	return retList, nil
}

func (r *OrderPostgres) Create(user_id int, products map[int]int) (order_id int, err error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	row := r.db.QueryRow("INSERT INTO orders(user_id) VALUES ($1) RETURNING id", user_id)
	if err := row.Scan(&order_id); err != nil {
		tx.Rollback()
		return 0, err
	}

	for key, val := range products {
		_, err = r.db.Exec("INSERT INTO store(order_id, product_id, count) VALUES ($1, $2, $3)", order_id, key, val)
		if err != nil {
			tx.Rollback()
			return 0, err
		}
	}

	return order_id, tx.Commit()
}

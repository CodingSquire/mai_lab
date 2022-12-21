package repository

import (
	"fmt"
	"go_services_lab/models"

	"github.com/jmoiron/sqlx"
)

type UserPostgres struct {
	db *sqlx.DB
}

func NewUserPostgres(db *sqlx.DB) *UserPostgres {
	return &UserPostgres{db: db}
}

func (r *UserPostgres) Get(id int) (user models.User, err error) {
	err = r.db.Get(&user, "SELECT * FROM users WHERE id=$1", id)
	return
}

func (r *UserPostgres) Create(user models.User) (id int, err error) {
	query := fmt.Sprintf("INSERT INTO users (name, login, password) VALUES ($1, $2, $3) RETURNING id")
	row := r.db.QueryRow(query, user.Name, user.Login, user.Password)

	if err = row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *UserPostgres) GetAll() (users []models.User, err error) {
	var user models.User

	rows, err := r.db.Queryx("SELECT * FROM users")
	for rows.Next() {
		err = rows.StructScan(&user)
		if err == nil {
			users = append(users, user)
		}
	}

	return
}

func (r *UserPostgres) Delete(id int) (int, error) {
	_, err := r.db.Exec("DELETE FROM users WHERE id = $1", id)
	return id, err
}

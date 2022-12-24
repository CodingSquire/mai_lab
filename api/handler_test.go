package handler

import (
	user "mai_lab/app/repository"
	store_users "mai_lab/db"
	postgres "mai_lab/postgres"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestRouter_CreateUser(t *testing.T) {
	db, err := postgres.NewPostgresDB(postgres.Config{
		Host:     "postgres_container",
		Port:     "5432",
		Username: "postgres",
		Password: "qwerty",
		DBName:   "postgres",
		SSLMode:  "disable",
	})

	if err != nil {
		t.Errorf("No accees to database: %s", err.Error())
	}

	ust := store_users.NewUsers(db)
	us := user.NewUsers(ust)
	rt := NewRouter(us)
	h := http.HandlerFunc(rt.CreateUser).ServeHTTP
	w := &httptest.ResponseRecorder{}
	r := httptest.NewRequest("POST", "/create", strings.NewReader(`{"name":"user5671", "email": "oleg@kovinev.ru", "Phone":"9167743904"}`))
	h(w, r)
	if w.Code != http.StatusCreated {
		t.Error("status wrong")
	}
}

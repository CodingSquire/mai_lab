package handler

import (
	"database/sql"
	"encoding/json"
	"errors"
	"github.com/google/uuid"
	user "mai_lab/app/repository"
	"net/http"
)

type Router struct {
	*http.ServeMux
	us *user.Users
}

func NewRouter(us *user.Users) *Router {
	r := &Router{
		ServeMux: http.NewServeMux(),
		us:       us,
	}
	r.HandleFunc("/create", http.HandlerFunc(r.CreateUser).ServeHTTP) //ServeHTTP - http handler
	r.HandleFunc("/read", http.HandlerFunc(r.ReadUser).ServeHTTP)
	return r
}

type User struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"` //teg for json
}

func (rt *Router) CreateUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost { //in json, method post
		http.Error(w, "bad method", http.StatusMethodNotAllowed)
		return
	}

	u := User{}
	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	bu := user.User{
		Name: u.Name,
	}

	nbu, err := rt.us.Create(bu)
	if err != nil {
		http.Error(w, "error when creating", http.StatusInternalServerError)
		return
	}
	//for a return
	w.WriteHeader(http.StatusCreated) //code 201

	_ = json.NewEncoder(w).Encode(
		User{
			ID:   nbu.ID,
			Name: nbu.Name,
		},
	)
}

// read?uid=...
func (rt *Router) ReadUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "bad method", http.StatusMethodNotAllowed)
		return
	}

	suid := r.URL.Query().Get("uid") //parsing
	if suid == "" {                  //uid null = error
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}
	uid, err := uuid.Parse(suid) //error format
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if (uid == uuid.UUID{}) { //null ?
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	nbu, err := rt.us.Read(uid) //read uuid
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			http.Error(w, "not found", http.StatusNotFound)
		} else {
			http.Error(w, "error when reading", http.StatusInternalServerError)
		}
		return
	}

	_ = json.NewEncoder(w).Encode(
		User{
			ID:   nbu.ID,
			Name: nbu.Name,
		},
	)
}

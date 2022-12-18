package handler

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/google/uuid"
	user "mai_lab/app/repository"
	"net/http"
)

type Router struct { //внешний адаптер принимающий внешние запросы
	//external adapter accepting external requests
	*http.ServeMux
	us *user.Users
}

func NewRouter(us *user.Users) *Router {
	r := &Router{
		ServeMux: http.NewServeMux(),
		us:       us,
	}
	r.HandleFunc("/create", http.HandlerFunc(r.CreateUser).ServeHTTP) //принимать функции
	//accept functions
	r.HandleFunc("/read", http.HandlerFunc(r.ReadUser).ServeHTTP)
	r.HandleFunc("/delete", http.HandlerFunc(r.DeleteUser).ServeHTTP)
	r.HandleFunc("/search", http.HandlerFunc(r.SearchUser).ServeHTTP)
	return r
}

type User struct { // получение от клиента
	// receiving from the client
	ID    uuid.UUID `json:"id"`
	Name  string    `json:"name"`
	Email string    `json:"email"`
	Phone string    `json:"phone"`
}

// curl  -X POST -d '{"name":"user5671", "email": "oleg@kovinev.ru", "Phone":"9167743904"}' localhost:8000/create
func (rt *Router) CreateUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost { //проверка метода
		//method check
		http.Error(w, "bad method", http.StatusMethodNotAllowed)
		return
	}

	u := User{}
	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		http.Error(w, "bad request", http.StatusBadRequest) //несоответствие формату
		//format mismatch
		return
	}

	bu := user.User{
		Name:  u.Name,
		Email: u.Email,
		Phone: u.Phone,
	}

	nbu, err := rt.us.Create(bu)
	if err != nil {
		http.Error(w, "error when creating", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	//если всё ок
	//if everything is ok
	_ = json.NewEncoder(w).Encode(
		User{ //вернулся объект. Нет json. В обратную сторону уже идёт
			//object returned. No json. Already going backwards
			ID:    nbu.ID,
			Name:  nbu.Name,
			Email: nbu.Email,
			Phone: nbu.Phone,
		},
	)
}

// read?uuid=...
func (rt *Router) ReadUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "bad method", http.StatusMethodNotAllowed)
		return
	}

	suid := r.URL.Query().Get("uuid")
	if suid == "" {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}
	uid, err := uuid.Parse(suid)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if (uid == uuid.UUID{}) {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	nbu, err := rt.us.Read(uid)
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
			ID:    nbu.ID,
			Name:  nbu.Name,
			Email: nbu.Email,
			Phone: nbu.Phone,
		},
	)
}

// delete?uuid=
func (rt *Router) DeleteUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "bad method", http.StatusMethodNotAllowed)
		return
	}

	suid := r.URL.Query().Get("uuid")
	if suid == "" {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}
	uid, err := uuid.Parse(suid)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if (uid == uuid.UUID{}) {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	nbu, err := rt.us.Delete(uid)
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
			ID:    nbu.ID,
			Name:  nbu.Name,
			Email: nbu.Email,
			Phone: nbu.Phone,
		},
	)
}

// /search?name=...
func (rt *Router) SearchUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "bad method", http.StatusMethodNotAllowed)
		return
	}

	q := r.URL.Query().Get("name")
	if q == "" {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	ch, err := rt.us.SearchUsers(q)
	if err != nil {
		http.Error(w, "error when reading", http.StatusInternalServerError)
		return
	}

	enc := json.NewEncoder(w)

	first := true
	fmt.Fprintf(w, "[")
	defer fmt.Fprintf(w, "]")

	for {
		select {
		case <-r.Context().Done():
			return
		case u, ok := <-ch:
			if !ok {
				return
			}
			if first {
				first = false
			} else {
				fmt.Fprintf(w, ",")
			}
			_ = enc.Encode(
				User{
					ID:    u.ID,
					Name:  u.Name,
					Email: u.Email,
					Phone: u.Phone,
				},
			)
			w.(http.Flusher).Flush()
		}
	}
}

package user

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/julienschmidt/httprouter"
	"log"
	"mai_lab/internal/apperror"
	"mai_lab/internal/handlers"
	"net/http"
)

const (
	usersURL = "/users"
	userURL  = "/users/:uuid"
)

type handler struct {
	service Service
}

func NewHandler(s Service) handlers.Handler {
	return &handler{service: s}
}

func (h *handler) Register(router *httprouter.Router) {
	router.HandlerFunc(http.MethodGet, usersURL, apperror.Middleware(h.GetList))
	router.HandlerFunc(http.MethodPost, usersURL, apperror.Middleware(h.CreateUser))
	router.HandlerFunc(http.MethodGet, userURL, apperror.Middleware(h.GetUserByID))
	router.HandlerFunc(http.MethodPatch, userURL, apperror.Middleware(h.PartiallyUpdateUser))
	router.HandlerFunc(http.MethodDelete, userURL, apperror.Middleware(h.DeleteUser))

}

func (h *handler) GetList(w http.ResponseWriter, r *http.Request) error {
	w.Header().Set("Content-Type", "application/json")

	//TODO RESPONSE
	all, err := h.service.GetAllUsers(r.Context())
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return err
	}

	allBytes, err := json.Marshal(all)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return err
	}

	w.WriteHeader(http.StatusOK)
	w.Write(allBytes)

	return nil
}

func (h *handler) GetUserByID(w http.ResponseWriter, r *http.Request) error {
	log.Println("GET USER")
	w.Header().Set("Content-Type", "application/json")

	params := r.Context().Value(httprouter.ParamsKey).(httprouter.Params)
	userUUID := params.ByName("uuid")
	userID, err := uuid.Parse(userUUID)

	log.Println(userUUID)
	log.Println(userID)
	if err != nil {
		return apperror.BadRequestError("123bad request")
	}
	user, err := h.service.GetUserByID(r.Context(), userID)
	if err != nil {
		return err
	}

	log.Println("marshal user")

	bytes, err := json.Marshal(user)
	if err != nil {
		return fmt.Errorf("failed to marshall user. error: %w", err)
	}

	w.WriteHeader(http.StatusOK)
	w.Write(bytes)

	return nil
}

func (h *handler) CreateUser(w http.ResponseWriter, r *http.Request) error {
	log.Println("CREATE USER")
	w.Header().Set("Content-Type", "application/json")
	log.Println("decode create user dto")

	var dto CreateUserDTO
	defer r.Body.Close()
	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		return apperror.BadRequestError("invalid JSON scheme")
	}

	if err := h.service.CreateUser(r.Context(), dto); err != nil {
		return err
	}

	w.WriteHeader(http.StatusCreated)
	//TODO RESPONSE ??
	json.NewEncoder(w).Encode(dto)

	return nil
}

func (h *handler) PartiallyUpdateUser(w http.ResponseWriter, r *http.Request) error {
	log.Println("PARTIALLY UPDATE USER")
	w.Header().Set("Content-Type", "application/json")

	params := r.Context().Value(httprouter.ParamsKey).(httprouter.Params)
	userUUID := params.ByName("uuid")
	userID, err := uuid.Parse(userUUID)
	if err != nil {
		return apperror.BadRequestError("bad request")
	}
	log.Println("decode update user dto")

	var updDTO UpdateUserDTO
	defer r.Body.Close()

	if err := json.NewDecoder(r.Body).Decode(&updDTO); err != nil {
		return apperror.BadRequestError("invalid JSON scheme")
	}
	updDTO.ID = userID

	if err := h.service.UpdateUser(r.Context(), updDTO); err != nil {
		return err
	}

	w.WriteHeader(http.StatusNoContent)

	return nil

	return nil
}

func (h *handler) DeleteUser(w http.ResponseWriter, r *http.Request) error {
	log.Println("DELETE USER")
	w.Header().Set("Content-Type", "application/json")

	log.Println("get uuid from context")
	params := r.Context().Value(httprouter.ParamsKey).(httprouter.Params)
	userUUID := params.ByName("uuid")
	userID, err := uuid.Parse(userUUID)
	if err != nil {
		return apperror.BadRequestError("bad request")
	}
	err = h.service.DeleteUser(r.Context(), userID)
	if err != nil {
		return err
	}

	w.WriteHeader(http.StatusNoContent)

	return nil
}

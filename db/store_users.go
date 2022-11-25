package store_users

import (
	"database/sql"
	"github.com/google/uuid"
	user "mai_lab/app/repository"
)

var _ user.UserStore = &Users{} //type matching. Variable interface
//stupidity protection. The design matches the interface. A set of methods corresponds to an interface

type Users struct {
	//	sync.Mutex                         //protection. Miscellaneous external requests
	m map[uuid.UUID]user.User //map by uuid. //ID, Name
}

func NewUsers() *Users {
	return &Users{
		m: make(map[uuid.UUID]user.User), //map initialization
	}
}

func (us *Users) Create(u user.User) (*uuid.UUID, error) {
	uid := uuid.New()
	u.ID = uid
	us.m[u.ID] = u
	return &uid, nil
}

func (us *Users) Read(uid uuid.UUID) (*user.User, error) {
	u, ok := us.m[uid]
	if ok {
		return &u, nil
	}
	return nil, sql.ErrNoRows
}

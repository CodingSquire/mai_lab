package storage

import (
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgconn"
	"log"
	"mai_lab/internal/user/model"
	"strings"
)

type storagePG struct {
	client PostgreSQLClient
}

func NewPostgreStorage(client PostgreSQLClient) Storage {
	return &storagePG{
		client: client,
	}
}

const (
	scheme      = "public"
	table       = "product"
	tableScheme = scheme + "." + table
)

func formatQuery(q string) string {
	return strings.ReplaceAll(strings.ReplaceAll(q, "\t", ""), "\n", " ")
}

func (s *storagePG) Create(ctx context.Context, user model.User) error {
	q := `
		INSERT INTO users 
		    (name,email,mobile,password_hash) 
		VALUES 
		       ($1, $2,$3,$4) 
		RETURNING id
	`
	log.Printf("SQL Query: %s", formatQuery(q))
	var ID int
	row := s.client.QueryRow(ctx, q, user.Name, user.Mobile, user.Email, user.PasswordHash)
	if err := row.Scan(&ID); err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			pgErr = err.(*pgconn.PgError)
			newErr := fmt.Errorf(fmt.Sprintf("SQL Error: %s, Detail: %s, Where: %s, Code: %s, SQLState: %s", pgErr.Message, pgErr.Detail, pgErr.Where, pgErr.Code, pgErr.SQLState()))
			log.Println(newErr)
			return newErr
		}
		return err
	}

	return nil
}

func (s *storagePG) GetUser(ctx context.Context, id string) (model.User, error) {
	q := `
		SELECT  id, name, email, mobile FROM public.users WHERE id = $1
	`
	log.Printf(fmt.Sprintf("SQL Query: %s", formatQuery(q)))

	var u model.User
	row := s.client.QueryRow(ctx, q, id)
	err := row.Scan(&u.ID, &u.Name, &u.Email, &u.Mobile)
	if err != nil {
		return model.User{}, err
	}

	return u, nil
}
func (s *storagePG) GetAll(ctx context.Context) ([]model.User, error) {
	q := `
		SELECT id, name, email, mobile FROM public.users;
	`
	log.Printf(fmt.Sprintf("SQL Query: %s", formatQuery(q)))

	rows, err := s.client.Query(ctx, q)
	if err != nil {
		return nil, err
	}

	users := make([]model.User, 0)

	for rows.Next() {
		var u model.User

		err = rows.Scan(&u.ID, &u.Name, &u.Email, &u.Mobile)
		if err != nil {
			return nil, err
		}

		users = append(users, u)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func (s *storagePG) Update(ctx context.Context, user model.User) error {
	//TODO implement me
	panic("implement me")
}
func (s *storagePG) Delete(ctx context.Context, id uuid.UUID) error {
	//TODO implement me
	panic("implement me")
}

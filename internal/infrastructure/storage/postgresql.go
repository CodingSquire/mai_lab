package storage

import (
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgconn"
	"log"
	"mai_lab/internal/domain/models"
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

func (s *storagePG) Create(ctx context.Context, user models.User) error {
	q := `
		INSERT INTO users_service.public.users
		    (name,email,mobile,password_hash) 
		VALUES 
		       ($1,$2,$3,$4) 
		RETURNING id
	`
	log.Printf("SQL Query: %s", formatQuery(q))
	var ID uuid.UUID
	row := s.client.QueryRow(ctx, q, user.Name, user.Mobile, user.Email, user.PasswordHash)
	if err := row.Scan(&ID); err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			pgErr = err.(*pgconn.PgError)
			newErr := fmt.Errorf(fmt.Sprintf("SQL Error: %s, Detail: %s, Where: %s, Code: %s, SQLState: %s",
				pgErr.Message, pgErr.Detail, pgErr.Where, pgErr.Code, pgErr.SQLState()))

			log.Println(newErr)
			return newErr
		}
		log.Println(err)
		return err
	}

	log.Println("ID: ", ID)
	return nil
}

func (s *storagePG) GetUser(ctx context.Context, id uuid.UUID) (models.User, error) {
	q := `
		SELECT  id, name, email, mobile FROM users_service.public.users WHERE id = $1
	`
	log.Printf(fmt.Sprintf("SQL Query: %s", formatQuery(q)))

	var u models.User
	row := s.client.QueryRow(ctx, q, id)
	err := row.Scan(&u.ID, &u.Name, &u.Email, &u.Mobile)
	if err != nil {
		return models.User{}, err
	}

	return u, nil
}

func (s *storagePG) GetAll(ctx context.Context) ([]models.User, error) {
	q := `
		SELECT id, name, email, mobile FROM users_service.public.users;
	`
	log.Printf(fmt.Sprintf("SQL Query: %s", formatQuery(q)))

	rows, err := s.client.Query(ctx, q)
	if err != nil {
		return nil, err
	}

	users := make([]models.User, 0)

	for rows.Next() {
		var u models.User

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

func (s *storagePG) Update(ctx context.Context, user models.User) error {
	//TODO implement me
	panic("implement me")
}

func (s *storagePG) Delete(ctx context.Context, id uuid.UUID) error {
	q := `
		DELETE FROM users_service.public.users WHERE id = $1
	`
	log.Printf(fmt.Sprintf("SQL Query: %s", formatQuery(q)))

	var recordID uuid.UUID
	err := s.client.QueryRow(ctx, q, id).Scan(recordID)
	if err != nil {
		return err
	}
	return nil
}

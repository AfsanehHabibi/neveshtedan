package postgres

import (
	"context"

	"github.com/AfsanehHabibi/neveshtedan/graph/model"
	"github.com/AfsanehHabibi/neveshtedan/pkg/repository"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"golang.org/x/crypto/bcrypt"
)

type PostgresUserRepository struct {
	con *pgxpool.Pool
}

func NewUserRepository(con *pgxpool.Pool) repository.UserRepository {
	return &PostgresUserRepository{con: con}
}

func (r *PostgresUserRepository) Add(ctx context.Context, user model.NewUser) (int, error) {
	query := `
	INSERT INTO users (username, password)
	VALUES ($1, $2)
	RETURNING id
	`
	hashedPassword, _ := HashPassword(user.Password)
	var id int
	err := r.con.QueryRow(ctx, query, user.Username, hashedPassword).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (r *PostgresUserRepository) GetById(ctx context.Context, id int) (*model.User, error) {
	query := `
	SELECT id, username
	FROM users
	WHERE id = $1
	`
	var user model.User
	err := r.con.QueryRow(ctx, query, id).Scan(&user.ID, &user.Username)
	if err == pgx.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *PostgresUserRepository) GetIdIfExists(ctx context.Context, username string, password string) (*int, error) {
	query := `
	SELECT id, password
	FROM users
	WHERE username = $1
	`
	var id int
	var hashedPassword string
	err := r.con.QueryRow(ctx, query, username).Scan(&id, &hashedPassword)
	if err == pgx.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	if CheckPasswordHash(password, hashedPassword) {
		return &id, nil
	}
	return nil, nil
}

func (r *PostgresUserRepository) Clear(ctx context.Context) error {
	_, err := r.con.Exec(context.Background(), "TRUNCATE TABLE users;")
	return err
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

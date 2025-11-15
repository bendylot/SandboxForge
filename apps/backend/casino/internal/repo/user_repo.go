package repo

import (
	"context"
	"errors"

	"casino.sandboxforge.tech/casino-api/internal/models"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepo struct {
	DB *pgxpool.Pool
}

func NewUserRepo(db *pgxpool.Pool) *UserRepo {
	return &UserRepo{DB: db}
}

func (r *UserRepo) Create(ctx context.Context, login, hashed string) (int64, error) {
	var id int64
	err := r.DB.QueryRow(ctx,
		`INSERT INTO users(login, password) VALUES($1,$2) RETURNING id`,
		login, hashed,
	).Scan(&id)
	return id, err
}

// func (r *UserRepo) ByEmail(ctx context.Context, email string) (*models.User, error) {
// 	row := r.DB.QueryRow(ctx, `SELECT id, email, password, created_at FROM users WHERE email=$1`, email)
// 	var u models.User
// 	if err := row.Scan(&u.ID, &u.Email, &u.Password, &u.CreatedAt); err != nil {
// 		if errors.Is(err, pgx.ErrNoRows) {
// 			return nil, nil
// 		}
// 		return nil, err
// 	}
// 	return &u, nil
// }

func (r *UserRepo) ByLogin(ctx context.Context, login string) (*models.User, error) {
	row := r.DB.QueryRow(ctx, `SELECT id, login, password, created_at FROM users WHERE login=$1`, login)
	var u models.User
	if err := row.Scan(&u.ID, &u.Login, &u.Password, &u.CreatedAt); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &u, nil
}

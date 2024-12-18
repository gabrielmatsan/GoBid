package services

import (
	"context"
	"errors"

	"github.com/gabrielmatsan/GoBid/internal/store/pgstore"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
	"golang.org/x/crypto/bcrypt"
)

var ErrDuplicatedEmailOrUsername = errors.New("Duplicated email or username")

type UserService struct {
	pool    *pgxpool.Pool
	queries *pgstore.Queries
}

func NewUserService(pool *pgxpool.Pool) UserService {
	return UserService{
		pool:    pool,
		queries: pgstore.New(pool),
	}
}

func (s *UserService) CreateUser(ctx context.Context, userName, email, password, bio string) (uuid.UUID, error) {

	hash, err := bcrypt.GenerateFromPassword([]byte(password), 12)

	if err != nil {
		return uuid.UUID{}, err
	}

	args := pgstore.CreateUserParams{
		Username:     userName,
		Email:        email,
		PasswordHash: hash,
		Bio:          bio,
	}

	id, err := s.queries.CreateUser(ctx, args)

	if err != nil {
		var pgErr *pgconn.PgError

		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return uuid.UUID{}, ErrDuplicatedEmailOrUsername
		}

		return uuid.UUID{}, err
	}
	return id, nil
}

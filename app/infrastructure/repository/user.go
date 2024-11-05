package repository

import (
	"database/sql"
	"errors"

	"github.com/KakinokiKanta/go-dev-template/domain"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) domain.UserRepository {
	return UserRepository{
		db: db,
	}
}


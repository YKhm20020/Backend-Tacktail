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

func (repo UserRepository) Create(user domain.User) (domain.User, error) {
	query := `
		INSERT INTO users (id, name, password) VALUES ($1, $2, $3);
	`
	_, err := repo.db.Exec(query, user.ID(), user.Name(), user.Password())

	if err != nil {
		return domain.User{}, err
	}

	return user, nil
}

func (repo UserRepository) FindByName(name string) (domain.User, error) {
	query := `
		SELECT id, name, password FROM users WHERE name=$1;
	`

	row := repo.db.QueryRow(query, name)

	if err := row.Err(); err != nil {
		return domain.User{}, errors.New("でーたべーすから値取れず")
	}

	var dbUser dbUser

	err := row.Scan(&dbUser.id, &dbUser.name, &dbUser.password)

	if err != nil {
		return domain.User{}, errors.New("DBユーザー型に置き換えれず")
	}

	user := domain.NewUser(dbUser.id, dbUser.name, dbUser.password)

	return user, nil
}

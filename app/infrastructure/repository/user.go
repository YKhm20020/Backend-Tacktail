package repository

import (
	"database/sql"
	"errors"

	"github.com/YKhm20020/Backend-Tacktail/domain"
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
		INSERT INTO users (id, name, password, story) VALUES ($1, $2, $3, $4);
	`
	_, err := repo.db.Exec(query, user.ID(), user.Name(), user.Password(), user.Story())

	if err != nil {
		return domain.User{}, err
	}

	return user, nil
}

func (repo UserRepository) UpdateStory(id string) error {
	query := `
		UPDATE users SET story = story + 1 WHERE id = $1;
	`

	_, err := repo.db.Exec(query, id)

	if err != nil {
		return err
	}

	return nil
}

func (repo UserRepository) FindByName(name string) (domain.User, error) {
	query := `
		SELECT id, name, password, story FROM users WHERE name=$1;
	`

	var dbUser dbUser
	err := repo.db.QueryRow(query, name).Scan(&dbUser.id, &dbUser.name, &dbUser.password, &dbUser.story)

	if err != nil {
		if err == sql.ErrNoRows {
			return domain.User{}, errors.New("not found user")
		}
		return domain.User{}, err
	}

	user := domain.ReUser(dbUser.id, dbUser.name, dbUser.password, dbUser.story)

	return user, nil
}

func (repo UserRepository) FindByID(id string) (domain.User, error) {
	query := `
		SELECT id, name, password, story FROM users WHERE id=$1;
	`

	var dbUser dbUser
	err := repo.db.QueryRow(query, id).Scan(&dbUser.id, &dbUser.name, &dbUser.password, &dbUser.story)

	if err != nil {
		if err == sql.ErrNoRows {
			return domain.User{}, errors.New("not found user")
		}
		return domain.User{}, err
	}

	user := domain.ReUser(dbUser.id, dbUser.name, dbUser.password, dbUser.story)

	return user, nil
}

package mysql

import (
	"context"
	"database/sql"

	"github.com/YngviWarrior/microservice-user/infra/database"
	"github.com/YngviWarrior/microservice-user/infra/database/mysql/repositorydto"
)

type userRepository struct {
	Db database.DatabaseInterface
}

type UserRepositoryInterface interface {
	Create(*repositorydto.InputUser) (bool, error)
	GetUserByEmail(string) (repositorydto.OutputUser, error)
}

func NewUserRepository(db database.DatabaseInterface) UserRepositoryInterface {
	return &userRepository{
		Db: db,
	}
}

func (u *userRepository) Create(in *repositorydto.InputUser) (bool, error) {
	tx, err := u.Db.CreateConnection().BeginTx(context.Background(), &sql.TxOptions{})
	if err != nil {
		return false, err
	}

	stmt, err := tx.Prepare(`
		INSERT INTO user (
			name,
			email,
			active
		) VALUES (?, ?, ?);
	`)

	if err != nil {
		return false, err
	}

	defer stmt.Close()

	_, err = stmt.Exec(
		in.Name,
		in.Email,
		in.Active,
	)

	if err != nil {
		return false, err
	}

	err = tx.Commit()
	if err != nil {
		return false, err
	}

	return true, nil
}

func (u *userRepository) GetUserByEmail(email string) (user repositorydto.OutputUser, err error) {
	tx, err := u.Db.CreateConnection().BeginTx(context.Background(), &sql.TxOptions{})
	if err != nil {
		return
	}

	err = tx.QueryRow(`
		SELECT user, name, email, active
		FROM user
		WHERE email = ?
	`, email,
	).Scan(&user.User, &user.Name, &user.Email, &user.Active)

	switch {
	case err == sql.ErrNoRows:
	case err != nil:
		return
	}

	err = tx.Commit()
	if err != nil {
		return
	}

	return
}

package mysql

import (
	"context"
	"database/sql"

	"github.com/YngviWarrior/microservice-user/infra/database"
	"github.com/YngviWarrior/microservice-user/infra/database/mysql/repositorydto"
)

type userStrategyRepository struct {
	Db database.DatabaseInterface
}

type UserStrategyRepositoryInterface interface {
	Create(*repositorydto.InputUserStrategy) (bool, error)
	ListByUser(uint64) ([]*repositorydto.OutputUserStrategy, error)
}

func NewUserStrategyRepository(db database.DatabaseInterface) UserStrategyRepositoryInterface {
	return &userStrategyRepository{
		Db: db,
	}
}

func (u *userStrategyRepository) Create(in *repositorydto.InputUserStrategy) (bool, error) {
	tx, err := u.Db.CreateConnection().BeginTx(context.Background(), &sql.TxOptions{})
	if err != nil {
		return false, err
	}

	stmt, err := tx.Prepare(`
		INSERT INTO user_strategy(
			user,
			trade_config
		) VALUES (?, ?)
	`)

	if err != nil {
		return false, err
	}

	defer stmt.Close()

	_, err = stmt.Exec(
		in.User,
		in.TradeConfig,
	)

	if err != nil {
		return false, err
	}

	return true, nil
}

func (u *userStrategyRepository) ListByUser(user uint64) (out []*repositorydto.OutputUserStrategy, err error) {
	tx, err := u.Db.CreateConnection().BeginTx(context.Background(), &sql.TxOptions{})
	if err != nil {
		return
	}

	stmt, err := tx.Prepare(`
		SELECT exchange, name
		FROM user_trade_config
		WHERE user = ?
	`)

	if err != nil {
		return
	}

	defer stmt.Close()

	res, err := stmt.Query(user)

	if err != nil {
		return
	}

	defer res.Close()

	for res.Next() {
		var u repositorydto.OutputUserStrategy

		err = res.Scan(&u.UserStrategy, &u.User, &u.TradeConfig)

		if err != nil {
			return
		}

		out = append(out, &u)
	}

	tx.Commit()

	return
}

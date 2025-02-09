package usecase

import (
	"github.com/YngviWarrior/microservice-user/infra/database/mysql"
	usecasesdto "github.com/YngviWarrior/microservice-user/usecase/usecases_dto"
)

type usecase struct {
	UserRepo     mysql.UserRepositoryInterface
	UserStrategy mysql.UserStrategyRepositoryInterface
}

type UseCaseInterface interface {
	CreateUser(in *usecasesdto.InputCreateUser) (out *usecasesdto.OutputCreateUser, err error)
	GetUserByEmail(in *usecasesdto.InputGetUserByEmail) (out *usecasesdto.OutputGetUserByEmail, err error)
}

func NewUseCase(
	userRepo mysql.UserRepositoryInterface,
	userStrategy mysql.UserStrategyRepositoryInterface,
) UseCaseInterface {
	return &usecase{
		UserRepo:     userRepo,
		UserStrategy: userStrategy,
	}
}

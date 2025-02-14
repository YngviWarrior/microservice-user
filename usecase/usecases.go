package usecase

import (
	"github.com/YngviWarrior/microservice-user/infra/database/mysql"
	usecasesdto "github.com/YngviWarrior/microservice-user/usecase/usecases_dto"
)

type usecase struct {
	UserRepo mysql.UserRepositoryInterface
}

type UseCaseInterface interface {
	CreateUser(in *usecasesdto.InputCreateUser) (out *usecasesdto.OutputCreateUser, err error)
	GetUserByEmail(in *usecasesdto.InputGetUserByEmail) (out *usecasesdto.OutputGetUserByEmail, err error)
	GetUserById(in *usecasesdto.InputGetUserById) (out *usecasesdto.OutputGetUserById, err error)
}

func NewUseCase(
	userRepo mysql.UserRepositoryInterface,
) UseCaseInterface {
	return &usecase{
		UserRepo: userRepo,
	}
}

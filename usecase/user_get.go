package usecase

import (
	"errors"

	"github.com/YngviWarrior/microservice-user/infra/database/mysql"
	"github.com/YngviWarrior/microservice-user/infra/database/mysql/repositorydto"
	usecasesdto "github.com/YngviWarrior/microservice-user/usecase/usecases_dto"
)

type usecaseGetUserById struct {
	UserRepo mysql.UserRepositoryInterface
}

type UseCaseGetUserByIdInterface interface {
	GetUserById(in *usecasesdto.InputGetUserById) (out *usecasesdto.OutputGetUserById, err error)
}

func NewGetUserByIdUseCase(
	userRepo mysql.UserRepositoryInterface,
) UseCaseGetUserByIdInterface {
	return &usecaseGetUserById{
		UserRepo: userRepo,
	}
}

func (u *usecaseGetUserById) GetUserById(in *usecasesdto.InputGetUserById) (out *usecasesdto.OutputGetUserById, err error) {
	out = &usecasesdto.OutputGetUserById{}
	user, err := u.UserRepo.GetUserById(in.Id)
	if err != nil {
		return
	}

	if (user == repositorydto.OutputUser{}) {
		return nil, errors.New("user dont exists")
	}

	out.Active = user.Active
	out.Email = user.Email
	out.User = user.User
	out.Name = user.Name

	return
}

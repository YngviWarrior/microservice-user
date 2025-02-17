package usecase

import (
	"errors"

	"github.com/YngviWarrior/microservice-user/infra/database/mysql"
	"github.com/YngviWarrior/microservice-user/infra/database/mysql/repositorydto"
	usecasesdto "github.com/YngviWarrior/microservice-user/usecase/usecases_dto"
)

type usecaseGetUserByEmail struct {
	UserRepo mysql.UserRepositoryInterface
}

type UseCaseGetUserByEmailInterface interface {
	GetUserByEmail(in *usecasesdto.InputGetUserByEmail) (out *usecasesdto.OutputGetUserByEmail, err error)
}

func NewGetUserByEmailUseCase(
	userRepo mysql.UserRepositoryInterface,
) UseCaseGetUserByEmailInterface {
	return &usecaseGetUserByEmail{
		UserRepo: userRepo,
	}
}

func (u *usecaseGetUserByEmail) GetUserByEmail(in *usecasesdto.InputGetUserByEmail) (out *usecasesdto.OutputGetUserByEmail, err error) {
	out = &usecasesdto.OutputGetUserByEmail{}
	user, err := u.UserRepo.GetUserByEmail(in.Email)
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

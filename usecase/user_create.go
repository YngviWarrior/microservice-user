package usecase

import (
	"errors"
	"fmt"

	"github.com/YngviWarrior/microservice-user/infra/database/mysql"
	"github.com/YngviWarrior/microservice-user/infra/database/mysql/repositorydto"
	usecasesdto "github.com/YngviWarrior/microservice-user/usecase/usecases_dto"
)

type usecaseCreateUser struct {
	UserRepo mysql.UserRepositoryInterface
}

type UseCaseCreateUserInterface interface {
	CreateUser(in *usecasesdto.InputCreateUser) (out *usecasesdto.OutputCreateUser, err error)
}

func NewCreateUserUseCase(
	userRepo mysql.UserRepositoryInterface,
) UseCaseCreateUserInterface {
	return &usecaseCreateUser{
		UserRepo: userRepo,
	}
}

func (u *usecaseCreateUser) CreateUser(in *usecasesdto.InputCreateUser) (out *usecasesdto.OutputCreateUser, err error) {
	user, err := u.UserRepo.GetUserByEmail(in.Email)
	if err != nil {
		return
	}

	if (user != repositorydto.OutputUser{}) {
		return nil, errors.New("user already exists")
	}

	success, err := u.UserRepo.Create(&repositorydto.InputUser{
		Name:   in.Name,
		Email:  in.Email,
		Active: in.Active,
	})

	if err != nil {
		return
	}

	if !success {
		err = errors.New("failed to create user")
		return
	}

	user, err = u.UserRepo.GetUserByEmail(in.Email)
	if err != nil {
		return
	}

	fmt.Println(user)
	out = &usecasesdto.OutputCreateUser{
		User:   user.User,
		Name:   user.Name,
		Email:  user.Email,
		Active: user.Active,
	}

	return
}

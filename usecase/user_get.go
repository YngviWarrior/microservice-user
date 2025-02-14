package usecase

import (
	"errors"

	"github.com/YngviWarrior/microservice-user/infra/database/mysql/repositorydto"
	usecasesdto "github.com/YngviWarrior/microservice-user/usecase/usecases_dto"
)

func (u *usecase) GetUserById(in *usecasesdto.InputGetUserById) (out *usecasesdto.OutputGetUserById, err error) {
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

func (u *usecase) GetUserByEmail(in *usecasesdto.InputGetUserByEmail) (out *usecasesdto.OutputGetUserByEmail, err error) {
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

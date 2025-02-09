package usecase

import (
	"errors"

	"github.com/YngviWarrior/microservice-user/infra/database/mysql/repositorydto"
	usecasesdto "github.com/YngviWarrior/microservice-user/usecase/usecases_dto"
)

func (u *usecase) GetUserByEmail(in *usecasesdto.InputGetUserByEmail) (out *usecasesdto.OutputGetUserByEmail, err error) {
	user, err := u.UserRepo.GetUserByEmail(in.Email)
	if err != nil {
		return
	}

	if (user == repositorydto.OutputUser{}) {
		return nil, errors.New("user dont exists")
	}

	return
}

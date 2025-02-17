package grpcserver

import (
	"context"
	"errors"

	"github.com/YngviWarrior/microservice-user/infra/database/mysql"
	"github.com/YngviWarrior/microservice-user/infra/grpcServer/proto/pb"
	"github.com/YngviWarrior/microservice-user/infra/utils"
	"github.com/YngviWarrior/microservice-user/usecase"
	usecasesdto "github.com/YngviWarrior/microservice-user/usecase/usecases_dto"
)

func (g *grpcServer) GetUserById(ctx context.Context, in *pb.GetUserByIdRequest) (out *pb.UserResponse, err error) {
	userRepo := mysql.NewUserRepository(g.Db)
	userUseCase := usecase.NewGetUserByIdUseCase(userRepo)

	response, err := userUseCase.GetUserById(&usecasesdto.InputGetUserById{
		Id: in.GetId(),
	})
	if err != nil {
		return
	}

	out = &pb.UserResponse{
		User: []*pb.User{},
	}
	out.User = append(out.User, &pb.User{
		Name:   response.Name,
		Email:  response.Email,
		Active: response.Active,
		User:   response.User,
	})

	return
}

func (g *grpcServer) GetUserByEmail(ctx context.Context, in *pb.GetUserByEmailRequest) (out *pb.UserResponse, err error) {
	userRepo := mysql.NewUserRepository(g.Db)
	userUseCase := usecase.NewGetUserByEmailUseCase(userRepo)

	response, err := userUseCase.GetUserByEmail(&usecasesdto.InputGetUserByEmail{
		Email: in.GetEmail(),
	})
	if err != nil {
		return
	}

	out = &pb.UserResponse{
		User: []*pb.User{},
	}
	out.User = append(out.User, &pb.User{
		Name:   response.Name,
		Email:  response.Email,
		Active: response.Active,
		User:   response.User,
	})

	return
}

func (g *grpcServer) CreateUser(ctx context.Context, in *pb.CreateUserRequest) (out *pb.UserResponse, err error) {
	if !utils.IsValidEmail(in.GetEmail()) {
		return nil, errors.New("invalid email")
	}

	userRepo := mysql.NewUserRepository(g.Db)
	userUseCase := usecase.NewCreateUserUseCase(userRepo)
	response, err := userUseCase.CreateUser(&usecasesdto.InputCreateUser{
		Name:   in.GetName(),
		Email:  in.GetEmail(),
		Active: in.GetActive(),
	})

	if err != nil {
		return
	}

	out = &pb.UserResponse{
		User: []*pb.User{},
	}
	out.User = append(out.User, &pb.User{
		Name:   response.Name,
		Email:  response.Email,
		Active: response.Active,
		User:   response.User,
	})

	return
}

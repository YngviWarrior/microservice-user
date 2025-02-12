package grpcserver

import (
	"context"

	"github.com/YngviWarrior/microservice-user/infra/grpcServer/proto/pb"
)

func (g *grpcServer) CreateUserStrategy(ctx context.Context, in *pb.CreateUserStrategyRequest) (out *pb.UserStrategyResponse, err error) {

	// userStrategyRepo := mysql.NewUserStrategyRepository(g.Db)
	// userStrategy := usecase.NewUseCase(nil, userStrategyRepo)

	return
}

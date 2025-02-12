package grpcserver

import (
	"github.com/YngviWarrior/microservice-user/infra/database"
	"github.com/YngviWarrior/microservice-user/infra/grpcServer/proto/pb"
)

type grpcServer struct {
	pb.UnimplementedUserServiceServer
	Db database.DatabaseInterface
}

// mustEmbedUnimplementedUserServiceServer implements pb.UserServiceServer.
func (g grpcServer) mustEmbedUnimplementedUserServiceServer() {
	panic("unimplemented")
}

func NewGrpcServer(db database.DatabaseInterface) *grpcServer {
	return &grpcServer{
		Db: db,
	}
}

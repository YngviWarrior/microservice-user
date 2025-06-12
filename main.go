package main

import (
	"log"
	"net"
	"os"
	"time"

	"github.com/YngviWarrior/microservice-user/infra/database"
	grpcserver "github.com/YngviWarrior/microservice-user/infra/grpcServer"
	"github.com/YngviWarrior/microservice-user/infra/grpcServer/proto/pb"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal(".env file is missing")
	}

	file, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)

	switch os.Getenv("ENVIROMENT") {
	case "local":
		log.SetOutput(os.Stdout)
		log.SetOutput(file)

	case "testnet", "demo":
		log.SetOutput(os.Stdout)

	case "server":
		loc, _ := time.LoadLocation("America/Sao_Paulo")
		time.Local = loc

		log.SetOutput(file)
	}

	if err != nil {
		log.Printf("%v", err)
	}

	database := database.NewDatabase()

	userService := grpcserver.NewGrpcServer(database)

	grpcServer := grpc.NewServer()
	pb.RegisterUserServiceServer(grpcServer, userService)
	reflection.Register(grpcServer)

	lis, err := net.Listen("tcp", os.Getenv("PORT"))
	if err != nil {
		panic(err)
	}

	log.Println("Running at port ", os.Getenv("PORT"))
	if err := grpcServer.Serve(lis); err != nil {
		panic(err)
	}

}

package main

import (
	"log"
	"os"
	"time"

	"github.com/YngviWarrior/microservice-user/controller"
	"github.com/YngviWarrior/microservice-user/infra/database"
	"github.com/YngviWarrior/microservice-user/infra/database/mysql"
	"github.com/YngviWarrior/microservice-user/infra/server"
	"github.com/YngviWarrior/microservice-user/usecase"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal(".env file is missing")
	}

	file, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)

	database := database.NewDatabase()

	userRepo := mysql.NewUserRepository(database)
	userStrategyRepo := mysql.NewUserStrategyRepository(database)

	// should return all usecases
	allUseCases := usecase.NewUseCase(
		userRepo,
		userStrategyRepo,
	)
	// should return all controllers
	controllers := controller.NewController(allUseCases)

	switch os.Getenv("ENVIROMENT") {
	case "local":
		log.SetOutput(os.Stdout)
		log.SetOutput(file)

	case "testnet":
		log.SetOutput(os.Stdout)

	case "server":
		loc, _ := time.LoadLocation("America/Sao_Paulo")
		time.Local = loc

		log.SetOutput(file)
	}

	if err != nil {
		log.Printf("%v", err)
	}

	server.NewServer().InitServer(controllers)
}

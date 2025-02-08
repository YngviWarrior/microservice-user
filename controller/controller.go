package controller

import (
	"encoding/json"
	"log"
	"net/http"

	validator "github.com/YngviWarrior/data-validator"
	"github.com/YngviWarrior/microservice-user/usecase"
)

type controller struct {
	Usecase usecase.UseCaseInterface
}

type ControllerInterface interface {
	CreateUser(w http.ResponseWriter, r *http.Request)
}

type outputControllerDto struct {
	Status  int64    `json:"status,omitempty"`
	Message string   `json:"message,omitempty"`
	Data    any      `json:"data,omitempty"`
	Errors  []string `json:"errors,omitempty"`
}

func NewController(usecase usecase.UseCaseInterface) ControllerInterface {
	return &controller{
		Usecase: usecase,
	}
}

func (c *controller) InputValidation(w http.ResponseWriter, input any) bool {
	var send outputControllerDto
	validatorHandler := validator.NewValidator()
	errors := validatorHandler.Validation(input)

	if len(errors) > 0 {
		send.Status = 0
		send.Errors = errors

		resp, err := json.Marshal(send)

		if err != nil {
			log.Panicf("CIV02: %s", err)
		}

		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(resp))
		return false
	}

	return true
}

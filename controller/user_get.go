package controller

import (
	"errors"
	"net/http"

	controllerdto "github.com/YngviWarrior/microservice-user/controller/controller_dto"
	usecasesdto "github.com/YngviWarrior/microservice-user/usecase/usecases_dto"
	"github.com/gorilla/mux"
)

func (c *controller) GetUserByEmail(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var send outputControllerDto
	var err error

	vars := mux.Vars(r)            // Captura os parâmetros da URL
	email, exists := vars["email"] // Pega o valor de "id"

	if !exists || email == "" {
		w.WriteHeader(http.StatusInternalServerError)
		err = errors.New("id is mandatory")
		send.Status = 0
		send.Errors = append(send.Errors, err.Error())

		c.FormatResponse(w, send)
	}

	output, err := c.Usecase.GetUserByEmail(&usecasesdto.InputGetUserByEmail{
		Email: email,
	})

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		send.Status = 0
		send.Errors = append(send.Errors, err.Error())
	} else {
		send.Status = 1
		send.Message = "Success"
		send.Data = controllerdto.OutputGetUserByEmail{
			User:   output.User,
			Name:   output.Name,
			Email:  output.Email,
			Active: output.Active,
		}
	}

	c.FormatResponse(w, send)
}

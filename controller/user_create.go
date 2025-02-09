package controller

import (
	"encoding/json"
	"net/http"

	controllerdto "github.com/YngviWarrior/microservice-user/controller/controller_dto"
	usecasesdto "github.com/YngviWarrior/microservice-user/usecase/usecases_dto"
)

func (c *controller) CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var send outputControllerDto
	var input controllerdto.InputCreateUser

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)

		send.Errors = append(send.Errors, "invalid primitive type")
		c.FormatResponse(w, send)
		return
	}

	if !c.InputValidation(w, input) {
		return
	}

	if input.Active == nil {
		active := false
		input.Active = &active
	}

	output, err := c.Usecase.CreateUser(&usecasesdto.InputCreateUser{
		Name:   *input.Name,
		Email:  *input.Email,
		Active: *input.Active,
	})

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		send.Status = 0
		send.Errors = append(send.Errors, err.Error())
	} else {
		send.Status = 1
		send.Message = "Success"
		send.Data = controllerdto.OutputCreateUser{
			User:   output.User,
			Name:   output.Name,
			Email:  output.Email,
			Active: output.Active,
		}
	}

	c.FormatResponse(w, send)
}

package controllers

import (
	"net/http"

	"github.com/alexisbcz/panache/internal/database/repositories"
)

type ForgotPasswordController struct {
	usersRepository repositories.UsersRepository
}

func NewForgotPasswordController(usersRepository repositories.UsersRepository) *ForgotPasswordController {
	return &ForgotPasswordController{
		usersRepository: usersRepository,
	}
}

func (c *ForgotPasswordController) Show(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (c *ForgotPasswordController) Handle(w http.ResponseWriter, r *http.Request) error {
	return nil
}

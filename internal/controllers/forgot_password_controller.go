package controllers

import (
	"net/http"

	"github.com/alexisbcz/panache/internal/database/repositories"
	"github.com/alexisbcz/panache/internal/views/pages"
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
	return pages.ForgotPasswordPage().Render(w)
}

func (c *ForgotPasswordController) Handle(w http.ResponseWriter, r *http.Request) error {
	return nil
}

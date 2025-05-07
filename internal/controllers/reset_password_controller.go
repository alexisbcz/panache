package controllers

import (
	"net/http"

	"github.com/alexisbcz/panache/internal/database/repositories"
	"github.com/alexisbcz/panache/internal/views/pages"
)

type ResetPasswordController struct {
	usersRepository repositories.UsersRepository
}

func NewResetPasswordController(usersRepository repositories.UsersRepository) *ResetPasswordController {
	return &ResetPasswordController{
		usersRepository: usersRepository,
	}
}

func (c *ResetPasswordController) Show(w http.ResponseWriter, r *http.Request) error {
	return pages.ResetPasswordPage().Render(w)
}

func (c *ResetPasswordController) Handle(w http.ResponseWriter, r *http.Request) error {
	return nil
}

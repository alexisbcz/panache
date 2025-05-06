package controllers

import (
	"net/http"

	"github.com/alexisbcz/panache/internal/database/repositories"
	"github.com/alexisbcz/panache/internal/views/pages"
)

type SignInController struct {
	usersRepository repositories.UsersRepository
}

func NewSignInController(usersRepository repositories.UsersRepository) *SignInController {
	return &SignInController{
		usersRepository: usersRepository,
	}
}

func (c *SignInController) Show(w http.ResponseWriter, r *http.Request) error {
	return pages.SignInPage().Render(w)
}

func (c *SignInController) Handle(w http.ResponseWriter, r *http.Request) error {
	return nil
}

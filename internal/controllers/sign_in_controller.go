package controllers

import "net/http"

type SignInController struct{}

func NewSignInController() *SignInController {
	return &SignInController{}
}

func (c *SignInController) Show(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (c *SignInController) Handle(w http.ResponseWriter, r *http.Request) error {
	return nil
}

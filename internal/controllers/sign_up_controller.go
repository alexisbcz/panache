package controllers

import "net/http"

type SignUpController struct{}

func NewSignUpController() *SignUpController {
	return &SignUpController{}
}

func (c *SignUpController) Show(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (c *SignUpController) Handle(w http.ResponseWriter, r *http.Request) error {
	return nil
}

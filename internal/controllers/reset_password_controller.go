package controllers

import "net/http"

type ResetPasswordController struct{}

func NewResetPasswordController() *ResetPasswordController {
	return &ResetPasswordController{}
}

func (c *ResetPasswordController) Show(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (c *ResetPasswordController) Handle(w http.ResponseWriter, r *http.Request) error {
	return nil
}

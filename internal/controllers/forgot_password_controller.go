package controllers

import "net/http"

type ForgotPasswordController struct{}

func NewForgotPasswordController() *ForgotPasswordController {
	return &ForgotPasswordController{}
}

func (c *ForgotPasswordController) Show(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (c *ForgotPasswordController) Handle(w http.ResponseWriter, r *http.Request) error {
	return nil
}

package router

import (
	"fmt"
	"net/http"

	"github.com/alexisbcz/panache/internal/controllers"
)

func New() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "404 Error — this page does not exist", http.StatusNotFound)
	})
	mux.HandleFunc("GET /{$}", makeHandler(func(w http.ResponseWriter, r *http.Request) error {
		fmt.Fprint(w, "hello, world")
		return nil
	}))

	signUpController := controllers.NewSignUpController()
	mux.HandleFunc("GET /auth/sign_up/{$}", makeHandler(signUpController.Show))
	mux.HandleFunc("POST /auth/sign_up/{$}", makeHandler(signUpController.Handle))

	signInController := controllers.NewSignInController()
	mux.HandleFunc("GET /auth/sign_in/{$}", makeHandler(signInController.Show))
	mux.HandleFunc("POST /auth/sign_in/{$}", makeHandler(signInController.Handle))

	forgotPassword := controllers.NewForgotPasswordController()
	mux.HandleFunc("GET /auth/forgot_password/{$}", makeHandler(forgotPassword.Show))
	mux.HandleFunc("POST /auth/forgot_password/{$}", makeHandler(forgotPassword.Handle))

	resetPassword := controllers.NewResetPasswordController()
	mux.HandleFunc("GET /auth/reset_password/{$}", makeHandler(resetPassword.Show))
	mux.HandleFunc("POST /auth/reset_password/{$}", makeHandler(resetPassword.Handle))

	return mux
}

type httpHandler func(http.ResponseWriter, *http.Request)
type httpHandlerWithErr func(http.ResponseWriter, *http.Request) error

func makeHandler(handler httpHandlerWithErr) httpHandler {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := handler(w, r); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

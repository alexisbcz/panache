package router

import (
	"fmt"
	"io/fs"
	"net/http"

	"github.com/alexisbcz/panache/internal/controllers"
	"github.com/alexisbcz/panache/internal/database/repositories"
	"github.com/alexisbcz/panache/internal/public"
	"github.com/jackc/pgx/v5/pgxpool"
)

func New(dbpool *pgxpool.Pool) *http.ServeMux {
	mux := http.NewServeMux()

	// Use fs.Sub to strip the "assets" prefix from the embedded FS
	assetsFS, err := fs.Sub(public.FS, "assets")
	if err != nil {
		panic(fmt.Errorf("failed to sub fs: %w", err))
	}

	// Serve embedded files under /assets/
	mux.Handle("/", http.StripPrefix("/", http.FileServer(http.FS(assetsFS))))

	mux.HandleFunc("GET /{$}", makeHandler(func(w http.ResponseWriter, r *http.Request) error {
		_, err := fmt.Fprint(w, "this is my home page...")
		if err != nil {
			return err
		}

		// ...

		return nil
	}))

	usersRepository := repositories.NewUsersRepository(dbpool)

	signUpController := controllers.NewSignUpController(usersRepository)
	mux.HandleFunc("GET /auth/sign_up/{$}", makeHandler(signUpController.Show))
	mux.HandleFunc("POST /auth/sign_up/{$}", makeHandler(signUpController.Handle))

	signInController := controllers.NewSignInController(usersRepository)
	mux.HandleFunc("GET /auth/sign_in/{$}", makeHandler(signInController.Show))
	mux.HandleFunc("POST /auth/sign_in/{$}", makeHandler(signInController.Handle))

	forgotPasswordController := controllers.NewForgotPasswordController(usersRepository)
	mux.HandleFunc("GET /auth/forgot_password/{$}", makeHandler(forgotPasswordController.Show))
	mux.HandleFunc("POST /auth/forgot_password/{$}", makeHandler(forgotPasswordController.Handle))

	resetPasswordController := controllers.NewResetPasswordController(usersRepository)
	mux.HandleFunc("GET /auth/reset_password/{$}", makeHandler(resetPasswordController.Show))
	mux.HandleFunc("POST /auth/reset_password/{$}", makeHandler(resetPasswordController.Handle))

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

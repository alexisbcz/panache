package router

import (
	"io/fs"
	"net/http"

	"github.com/alexisbcz/panache/internal/controllers"
	"github.com/alexisbcz/panache/internal/database/repositories"
	"github.com/alexisbcz/panache/internal/public"
	"github.com/alexisbcz/x/exit"
	"github.com/jackc/pgx/v5/pgxpool"
)

func New(dbpool *pgxpool.Pool) *http.ServeMux {
	mux := http.NewServeMux()

	assetsFS, err := fs.Sub(public.FS, "assets")
	if err != nil {
		exit.WithErr(err)
	}
	mux.Handle("/", http.StripPrefix("/", http.FileServer(http.FS(assetsFS))))

	usersRepository := repositories.NewUsersRepository(dbpool)

	mux.HandleFunc("GET /{$}", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/sign-up/", http.StatusSeeOther)
	})

	mux.HandleFunc("/hotreload", hotReloadHandler)

	signUpController := controllers.NewSignUpController(usersRepository)
	mux.HandleFunc("GET /sign-up/{$}", makeHandler(signUpController.Show))
	mux.HandleFunc("POST /sign-up/{$}", makeHandler(signUpController.Handle))

	signInController := controllers.NewSignInController(usersRepository)
	mux.HandleFunc("GET /sign-in/{$}", makeHandler(signInController.Show))
	mux.HandleFunc("POST /sign-in/{$}", makeHandler(signInController.Handle))

	forgotPasswordController := controllers.NewForgotPasswordController(usersRepository)
	mux.HandleFunc("GET /forgot-password/{$}", makeHandler(forgotPasswordController.Show))
	mux.HandleFunc("POST /forgot-password/{$}", makeHandler(forgotPasswordController.Handle))

	resetPasswordController := controllers.NewResetPasswordController(usersRepository)
	mux.HandleFunc("GET /reset-password/{$}", makeHandler(resetPasswordController.Show))
	mux.HandleFunc("POST /reset-password/{$}", makeHandler(resetPasswordController.Handle))

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

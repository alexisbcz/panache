package main

import (
	"context"
	"log/slog"
	"net/http"

	"github.com/alexisbcz/panache/internal/router"
	"github.com/alexisbcz/x/env"
	"github.com/alexisbcz/x/exit"
	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	dbpool, err := pgxpool.New(context.Background(), env.GetVar("DATABASE_URL", "postgres://panache:panache@localhost/panache?sslmode=disable"))
	if err != nil {
		exit.WithErr(err)
	}
	defer dbpool.Close()

	if err := dbpool.Ping(context.Background()); err != nil {
		exit.WithErr(err)
	}

	mux := router.New(dbpool)
	addr := env.GetVar("HTTP_ADDR", ":3000")
	slog.Info("http server listening for requests", "addr", addr)
	if err := http.ListenAndServe(addr, mux); err != nil {
		exit.WithErr(err)
	}
}

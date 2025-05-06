package main

import (
	"context"
	"time"

	"github.com/alexisbcz/panache/internal/database/seeders"
	"github.com/alexisbcz/x/env"
	"github.com/alexisbcz/x/exit"
	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	dbpool, err := pgxpool.New(ctx, env.GetVar("DATABASE_URL", "postgres://panache:panache@localhost/panache?sslmode=disable"))
	if err != nil {
		exit.WithErr(err)
	}
	defer dbpool.Close()

	seeders := []seeders.Seeder{
		seeders.NewUserSeeder(dbpool),
	}

	for _, seeder := range seeders {
		if err := seeder.Seed(ctx); err != nil {
			exit.WithErr(err)
		}
	}
}

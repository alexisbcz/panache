package main

import (
	"context"
	"time"

	"github.com/alexisbcz/panache/internal/database/models"
	"github.com/alexisbcz/panache/internal/database/repositories"
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

	users := []*models.User{
		{Email: "socrates@athens.gr", Password: "i_know_that_i_know_nothing"},
		{Email: "plato@academy.gr", Password: "the_unexamined_life_is_not_worth_living"},
		{Email: "aristotle@lyceum.gr", Password: "man_is_by_nature_a_political_animal"},
		{Email: "descartes@cogito.fr", Password: "i_think_therefore_i_am"},
		{Email: "kant@konigsberg.de", Password: "dare_to_know"},
		{Email: "nietzsche@zarathustra.de", Password: "god_is_dead"},
		{Email: "hegel@dialectic.de", Password: "the_owl_of_minerva"},
		{Email: "marx@capital.de", Password: "workers_of_the_world_unite"},
		{Email: "foucault@discipline.fr", Password: "knowledge_is_power"},
		{Email: "derrida@deconstruction.fr", Password: "there_is_nothing_outside_the_text"},
		{Email: "camus@absurd.fr", Password: "one_must_imagine_sisyphus_happy"},
		{Email: "sartre@existential.fr", Password: "existence_precedes_essence"},
		{Email: "de_beauvoir@feminism.fr", Password: "one_is_not_born_a_woman"},
		{Email: "wittgenstein@language.uk", Password: "whereof_one_cannot_speak"},
		{Email: "russell@logic.uk", Password: "the_philosophy_of_logical_atomism"},
		{Email: "hume@empiricism.uk", Password: "reason_is_the_slave_of_the_passions"},
		{Email: "locke@tabula.uk", Password: "all_ideas_come_from_sensation"},
		{Email: "hobbes@leviathan.uk", Password: "life_is_nasty_brutish_and_short"},
		{Email: "spinoza@ethics.nl", Password: "god_or_nature"},
		{Email: "kierkegaard@faith.dk", Password: "leap_of_faith"},
	}

	usersRepository := repositories.NewUsersRepository(dbpool)

	for _, user := range users {
		if err := usersRepository.Store(ctx, user); err != nil {
			exit.WithErr(err)
		}
	}
}

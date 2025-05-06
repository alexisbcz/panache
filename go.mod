module github.com/alexisbcz/panache

go 1.24.2

require (
	github.com/alexisbcz/x v0.1.0
	github.com/jackc/pgx/v5 v5.7.4
	maragu.dev/gomponents v1.1.0
)

replace github.com/alexisbcz/x => ../x

require (
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgservicefile v0.0.0-20240606120523-5a60cdf6a761 // indirect
	github.com/jackc/puddle/v2 v2.2.2 // indirect
	golang.org/x/crypto v0.31.0 // indirect
	golang.org/x/sync v0.10.0 // indirect
	golang.org/x/text v0.21.0 // indirect
)

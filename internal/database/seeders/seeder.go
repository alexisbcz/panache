package seeders

import "context"

type Seeder interface {
	Seed(ctx context.Context) error
}

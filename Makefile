.PHONY: setup db-up db-down migration-up migration-down migration-fresh seed clean

setup:
	@echo "🔧 Installing tern for database migrations..."
	go install github.com/jackc/tern/v2@latest
	@echo "📦 Installing dependencies..."
	go mod tidy

db-up:
	@echo "🐳 Starting database container..."
	docker compose -f ./zarf/compose.dev.yml up -d

db-down:
	@echo "🛑 Stopping database container..."
	docker compose -f ./zarf/compose.dev.yml down

migration-up:
	@echo "⬆️  Running migrations..."
	TERN_MIGRATIONS="./internal/database/migrations" tern -c ./zarf/tern.dev.conf migrate

migration-down:
	@echo "⬇️  Rolling back migrations..."
	TERN_MIGRATIONS="./internal/database/migrations" tern -c ./zarf/tern.dev.conf migrate --destination 0

migration-fresh: migration-down migration-up
	@echo "🔄 Database freshed!"

seed:
	@echo "🌱 Seeding database with philosophers..."
	go run ./cmd/seed
	@echo "✨ Database successfully seeded!"

clean: db-down
	@echo "🧹 Development environment has been cleaned up!"

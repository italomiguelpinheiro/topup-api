# Variáveis
DB_URL=postgres://postgres:1234@localhost:5432/postgres?sslmode=disable
MIGRATIONS_DIR=./migrations

# -------- DOCKER --------
up:
	docker-compose up -d

down:
	docker-compose down

logs:
	docker-compose logs -f

ps:
	docker-compose ps

# -------- MIGRATIONS (Goose) --------
migrate-up:
	goose -dir $(MIGRATIONS_DIR) postgres "$(DB_URL)" up

migrate-down:
	goose -dir $(MIGRATIONS_DIR) postgres "$(DB_URL)" down

migrate-status:
	goose -dir $(MIGRATIONS_DIR) postgres "$(DB_URL)" status

# -------- API --------
run:
	go run ./cmd/main.go


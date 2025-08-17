MIGRATE       = docker compose run --rm migrate
MIGRATIONS_DIR = /migrations

migrate-up:
	$(MIGRATE) up

migrate-down:
	$(MIGRATE) down 1

migrate-force:
	$(MIGRATE) force $(version)

# create a new migration file: make migrate-create name=create_users_table
migrate-create:
	$(MIGRATE) create -ext sql -dir $(MIGRATIONS_DIR) -seq $(name)

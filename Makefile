
migrate-create:
	docker compose run --rm migrate create -seq -ext sql -dir /migrations $(name)

migrate-up:
	docker compose run --rm migrate -path=/migrations -database $$DB_URL up

migrate-down:
	docker compose run --rm migrate -path=/migrations -database $$DB_URL down 1

migrate-version:
	docker compose run --rm migrate -path=/migrations -database $$DB_URL version

migrate-force:
	docker compose run --rm migrate -path=/migrations -database $$DB_URL force $(v)

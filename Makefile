
migrate-up:
	goose -dir="internal/db/migrations" postgres postgresql://tssdr:tssdr@localhost:5432/booking up

migrate-down:
	goose -dir="internal/db/migrations" postgres postgresql://tssdr:tssdr@localhost:5432/booking down

generate-sql:
	sqlc generate

create_db:
	docker exec -it postgres psql --u tssdr -c 'create database booking'

run-air:
	@if [ "$(OS)" = "windows" ]; then \
		air -c .air.windows.toml; \
	elif [ "$(OS)" = "linux" ]; then \
		air -c .air.unix.toml; \
	else \
		echo "Usage: make run-air OS=windows|linux"; \
	fi
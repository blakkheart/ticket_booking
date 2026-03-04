
migrate-up:
	goose -dir="internal/db/migrations" postgres postgresql://wowsp:wowsp@localhost:5432/booking up

migrate-down:
	goose -dir="internal/db/migrations" postgres postgresql://wowsp:wowsp@localhost:5432/booking down

generate-sql:
	sqlc generate

create_db:
	docker exec -it postgres psql --u wowsp -c 'create database booking'

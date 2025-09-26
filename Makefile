
migrate-up:
	goose -dir="repository/migrations" postgres postgresql://wowsp:wowsp@localhost:5432/booking up

migrate-down:
	goose -dir="repository/migrations" postgres postgresql://wowsp:wowsp@localhost:5432/booking down

generate-sql:
	sqlc generate

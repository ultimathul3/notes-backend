PSQL_URL=postgres://root:root@localhost:5432/postgres?sslmode=disable

run:
	go run cmd/app/main.go

up_migrate:
	migrate -path ./migrations -database '$(PSQL_URL)' up

down_migrate:
	migrate -path ./migrations -database '$(PSQL_URL)' down -all

drop_migrate:
	migrate -path ./migrations -database '$(PSQL_URL)' drop -f
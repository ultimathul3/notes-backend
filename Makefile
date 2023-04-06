PSQL_URL=postgres://root:root@localhost:5432/postgres?sslmode=disable

run:
	go run cmd/app/main.go

generate:
	go generate ./...

test:
	go test -v ./internal/...

up_migrate:
	migrate -path ./migrations -database '$(PSQL_URL)' up

down_migrate:
	migrate -path ./migrations -database '$(PSQL_URL)' down -all

drop_migrate:
	migrate -path ./migrations -database '$(PSQL_URL)' drop -f

PSQL_URL=postgres://root:root@localhost:5432/postgres?sslmode=disable

run:
	go run cmd/app/main.go

generate:
	go generate ./...

test:
	go test -v ./internal/... && go test -v ./pkg/...

swag:
	swag init --parseDependency -g cmd/app/main.go

up_migrate:
	migrate -path ./migrations -database '$(PSQL_URL)' up

down_migrate:
	migrate -path ./migrations -database '$(PSQL_URL)' down -all

drop_migrate:
	migrate -path ./migrations -database '$(PSQL_URL)' drop -f

lines_count:
	find ./internal ./pkg -name '*.go' -not -path "*/mocks/*" -type f -print0 | xargs -0 cat | wc -l

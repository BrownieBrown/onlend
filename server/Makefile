include .env

build:
	@go build -o bin/onlend

run: build
	@./bin/onlend

test:
	@go test -v ./...

server:
	@go run ./cmd/main.go

userMock:
	mockgen -source=pkg/models/User.go -destination=mocks/mock_user.go -package=mocks

accountMock:
	mockgen -source=pkg/models/Account.go -destination=mocks/mock_account.go -package=mocks

postgresInit:
	docker run --name $(POSTGRES_NAME) -p $(POSTGRES_PORT):5432 -e POSTGRES_USER=$(POSTGRES_USER) -e POSTGRES_PASSWORD=$(POSTGRES_PASSWORD) -d postgres:16-alpine

postgres:
	docker exec -it onlend psql -U $(POSTGRES_USER) -d onlend

createDb:
	docker exec -it onlend createdb --username=$(POSTGRES_USER) --owner=$(POSTGRES_USER) $(POSTGRES_NAME)

dropDb:
	docker exec -it onlend dropdb --username=$(POSTGRES_USER) $(POSTGRES_NAME)

migrateUp:
	migrate -path internal/database/migrations -database "postgresql://$(POSTGRES_USER):$(POSTGRES_PASSWORD)@localhost:$(POSTGRES_PORT)/$(POSTGRES_NAME)?sslmode=disable" -verbose up

migrateDown:
	migrate -path internal/database/migrations -database "postgresql://$(POSTGRES_USER):$(POSTGRES_PASSWORD)@localhost:$(POSTGRES_PORT)/$(POSTGRES_NAME)?sslmode=disable" -verbose down

.PHONY: postgresinit postgres createdb dropdb migrateup migratedown

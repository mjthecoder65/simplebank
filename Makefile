simple_bank_db:
	@echo "Starting Simple Bank Database...."
	docker-compose up -d 

createdb:
	@echo "Created simple_bank database..."
	docker exec -it simple-bank-db --username=root --owner=root simple_bank

dropdb:
	@echo "Dropping simple_bank database..."
	docker exec -it dropdb simple_bank

migrate_up:
	@echo "Migrate Up..."
	migrate -path db/migration -database 'postgresql://admin:y7jHf&DNWG15@localhost:5030/main?sslmode=disable' -verbose up

migrate_down:
	@echo "Migrate down..."
	migrate -path db/migration -database 'postgresql://admin:y7jHf&DNWG15@localhost:5030/main?sslmode=disable' -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go

proto:
	rm -f pb/*.go
	protoc --proto_path=proto --go_out=pb --go_opt=paths=source_relative \
		--go-grpc_out=pb --go-grpc_opt=paths=source_relative \
		proto/*.proto

.PHONY: simple_bank_db createdb dropdb migrate_down migrate_up sql test server proto
DB_URL=postgresql://admin:admin@localhost:5432/postgres?sslmode=disable

migrate-up:
	migrate -path db/migration -database "$(DB_URL)" -verbose up
migrate-down:
	migrate -path db/migration -database "$(DB_URL)" -verbose down
run-test:
	go test -v ./...
run-server:
	go run main.go
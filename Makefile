DB_URL=postgresql://admin:admin@localhost:5432/simple_bank?sslmode=disable

migrate-up:
	migrate -path db/migration -database "$(DB_URL)" -verbose up
run-test:
	go test -v ./...
run-server:
	go run main.go
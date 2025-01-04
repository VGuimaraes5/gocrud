build:
	go build ./cmd/gocrud-server

run: 
	go run ./cmd/gocrud-server

import:
	go mod tidy

docker-run:
	docker-compose up -d --build

docker-stop:
	docker-compose down

docker-run-db:
	docker-compose up -d postgres-db
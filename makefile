run:
	go run ./cmd/api

test:
	go test ./...

swagger:
	swag init -g cmd/api/main.go -o docs

docker-build:
	docker build -t your-app .

docker-up:
	docker-compose up -d

docker-down:
	docker-compose down

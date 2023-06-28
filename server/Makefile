up:
		docker compose up --build
down:
		docker compose down
build-dev-db:
		docker build -t discord-mongo mongo
run-dev-db:
		docker run -d --name discord-mongo-container -p 27017:27017 discord-mongo
dev:
		go run cmd/main.go

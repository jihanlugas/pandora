server:
	swag init
	go run main.go server

up:
	go run main.go db up

down:
	go run main.go db down

seed:
	go run main.go db seed

reset-db:
	go run main.go db reset

region-up:
	go run main.go db region-up

region-down:
	go run main.go db region-down

region-seed:
	go run main.go db region-seed

region-reset-db:
	go run main.go db region-reset

build :
	go build -o ./bin/main main.go
	cp .env ./bin/.env

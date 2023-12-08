.PHONY: build
build:
	go build -o ./internal/db_project/microservices/redis_api/bin/redis_api.exe -v ./internal/db_project/microservices/redis_api/main.go
	go build -o ./internal/db_project/microservices/postgre_api/bin/postgre_api.exe -v ./internal/db_project/microservices/postgre_api/main.go
.PHONY: test
run: build
	./internal/db_project/microservices/redis_api/bin/redis_api.exe
	./internal/db_project/microservices/postgre_api/bin/postgre_api.exe

.DEFAULT_GOAL := build
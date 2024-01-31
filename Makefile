.PHONY: build
.PHONY build:
	@go build -o bin/server

run: build
	@./bin/server

test:
	@go test -v ./...

dcr:
	@docker-compose -f ./build/docker/docker-compose.yml restart

dcu:
	@docker-compose -f ./build/docker/docker-compose.yml up --build -d --remove-orphans
	make dl

dc-down:
	@docker-compose -f ./build/docker/docker-compose.yml down --remove-orphans

dcw:
	@docker-compose -f ./build/docker/docker-compose.yml up --build catalog -d
	make dl

dcs:
	@docker-compose -f ./build/docker/docker-compose.yml stop

dl:
	@docker logs catalog --follow

watch:
	@watchexec -e go,env ./watch.sh

ssh:
	@docker exec -it catalog /bin/bash

migrate:
	@docker exec -it catalog sh -c 'migrate -path=/migrations -database $$MYSQL_DSN up'

flush-redis:
	@docker exec -it redis redis-cli FLUSHALL
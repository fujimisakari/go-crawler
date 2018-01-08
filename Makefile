cmd = '' # for migrate task

.PHONY: help build server migrate

help:
	@echo "Usage:"
	@echo "  make build                      # go build"
	@echo "  make server                     # run server"
	@echo "  make migrate cmd=[up|down]      # execute migration."

build:
	go build -o ./bin/go-crawler

server:
	./bin/go-crawler server

migrate:
	migrate -path=./backend/infrastracture/migrations -database 'mysql://$(DB_USER):$(DB_PASSWORD)@tcp($(DATABASE_URL):3306)/$(DB_NAME)' $(cmd)

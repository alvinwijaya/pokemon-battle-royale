ifneq (,$(wildcard ./.env))
    include .env
    export
endif

migrate_up:
	migrate -path db/migration -database "$(DB_CONNECTION)://$(DB_USER):$(DB_PASSWORD)@tcp($(DB_HOST):$(DB_PORT))/$(DB_NAME)" -verbose up

migrate_down:
	migrate -path db/migration -database "$(DB_CONNECTION)://$(DB_USER):$(DB_PASSWORD)@tcp($(DB_HOST):$(DB_PORT))/$(DB_NAME)" -verbose down
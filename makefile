PSQL_USER=root
PSQL_PASSWORD=secret
POSTGRES_DOCKER_NAME=postgres:12-alpine
DB_NAME=ticktr

postgres:
	docker run --name postgres12 -p 5432:5432 -e POSTGRES_USER=$(PSQL_USER) -e POSTGRES_PASSWORD=$(SQL_PASSWORD) -d $(POSTGRES_DOCKER_NAME)

createdb:
	docker exec -it postgres12 createdb --username=root --owner=root $(DB_NAME) 
	
dropdb:
	docker exec -it postgres12 dropdb $(DB_NAME) 

dbshell:
	docker exec -it postgres12 psql

migrateup:
	migrate -path internal/db/migration -database "postgresql://root:secret@localhost:5432/$(DB_NAME)?sslmode=disable" -verbose up

migratedown:
	migrate -path internal/db/migration -database "postgresql://root:secret@localhost:5432/$(DB_NAME)?sslmode=disable" -verbose down

sqlc:
	sqlc generate
PSQL_USER=root
PSQL_PASSWORD=secret
POSTGRES_DOCKER_NAME=postgres:12-alpine

postgres:
	docker run --name postgres12 -p 5432:5432 -e POSTGRES_USER=$(PSQL_USER) -e POSTGRES_PASSWORD=$(SQL_PASSWORD) -d $(POSTGRES_DOCKER_NAME)


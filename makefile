# Environment variables:
# POSTGRES_NAME: The name of the PostgreSQL database to create
# POSTGRES_USER: The username for the PostgreSQL database
# POSTGRES_PASSWORD: The password for the PostgreSQL database
# POSTGRES_HOST: The IP address for the PostgreSQL container

# Load environment variables from the specified env file or default to dev.env
# User must provide an ENV variable, like 'make ENV="prod" <target>'.
ENV ?= dev
include .env/$(ENV).env

# This target runs a PostgreSQL container widockth specified environment variables
run_postgres:
	docker run \
	-d \
	--name ${POSTGRES_NAME}_${ENV}_postgres17 \
	-e POSTGRES_USER=${POSTGRES_USER} \
	-e POSTGRES_PASSWORD=${POSTGRES_PASSWORD} \
	-p ${POSTGRES_PORT}:5432 \
	-h ${POSTGRES_HOST} \
	-v data_${POSTGRES_NAME}_${ENV}_postgres17:/var/lib/postgresql/data \
	--restart unless-stopped \
	--health-cmd="pg_isready -U ${POSTGRES_USER}" \
	--health-interval=10s \
	--health-timeout=5s \
	--health-retries=5 \
	--network=bridge \
	postgres:17-alpine
.PHONY: postgres

# This target creates a PostgreSQL database using the specified username and 
# owner.
create_db:
	docker exec -it ${POSTGRES_NAME}_${ENV}_postgres17 createdb \
	--username=${POSTGRES_USER} \
	--owner=${POSTGRES_USER} \
	${POSTGRES_NAME}
.PHONY: createdb

# This target drops the specified PostgreSQL database using the provided name.
drop_db:
	docker exec -it ${POSTGRES_NAME}_${ENV}_postgres17 dropdb ${POSTGRES_NAME}
.PHONY: dropdb

# This target adds a new migration file to the database schema.
# User must provide a NAME variable, like 'add_migration NAME=migration_name'.
add_migration:
	migrate create -ext sql -dir db/migration -seq $(NAME)
.PHONY: add_migration

# This target applies database migrations to the PostgreSQL database.
migrate_up:
	migrate \
	-path db/migration \
	-database "postgresql://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${POSTGRES_HOST}:${POSTGRES_PORT}/${POSTGRES_NAME}?sslmode=disable" \
	-verbose up || { echo "Migration failed"; exit 1; }
.PHONY: migrate_up

# This target rolls back the last applied database migration.
migrate_down:
	migrate \
	-path db/migration \
	-database "postgresql://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${POSTGRES_HOST}:${POSTGRES_PORT}/${POSTGRES_NAME}?sslmode=disable" \
	-verbose down || { echo "Migration failed"; exit 1; }
.PHONY: migrate_down

sqlc:
	sqlc generate
.PHONY: sqlc

test:
	go test -v -cover ./...
.PHONY: test

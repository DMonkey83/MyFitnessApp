DB_URL=postgresql://root:secret@localhost:5433/workout?sslmode=disable

network:
	cd workout-be && docker network create fitness-network

postgres:
	cd workout-be && docker run --name postgres --network fitness-network -p 5433:5433 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:15.4-alpine3.18

mysql:
	cd workout-be && docker run --name mysql8 -p 3306:3306  -e MYSQL_ROOT_PASSWORD=secret -d mysql:8

createdb:
	cd workout-be && docker exec -it postgres createdb --username=root --owner=root workout

dropdb:
	cd workout-be && docker exec -it postgres dropdb workout

migrateup:
	migrate -path workout-be/db/migration -database "$(DB_URL)" -verbose up

migratedown:
	migrate -path db/migration -database "$(DB_URL)" -verbose down
	
migrateuplocal:
	migrate -path workout-be/db/migration -database "postgresql://evilnis:Lon19ska83@localhost:5432/workout?sslmode=disable" -verbose up

migratedownlocal:
	migrate -path db/migration -database "postgresql://evilnis:Lon19ska83@localhost:5432/workout?sslmode=disable" -verbose down

sqlc:
	cd workout-be && sqlc generate

test: 
	cd workout-be && go test -v -cover ./...

server:
	cd workout-be && go run main.go

dbdocs:
	cd workout-be && dbdocs build loc doc/db.dbml

dbschema:
	cd workout-be && dbml2sql --postgres - doc/schema.sql doc/db.dbml

mock:
	cd workout-be && mockgen -package mockdb -destination db/mock/store.go github.com/DMonkey83/MyFitnessApp/workout-be/db/sqlc Store

.PHONY: migrateup migratedown sqlc test dbdocs dbschema

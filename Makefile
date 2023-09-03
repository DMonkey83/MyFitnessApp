DB_URL=postgresql://root:secret@localhost:5433/workout?sslmode=disable

network:
	docker network create fitness-network

postgres:
	docker run --name postgres --network fitness-network -p 5433:5433 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:15.4-alpine3.18

mysql:
	docker run --name mysql8 -p 3306:3306  -e MYSQL_ROOT_PASSWORD=secret -d mysql:8

createdb:
	docker exec -it postgres createdb --username=root --owner=root workout

dropdb:
	docker exec -it postgres dropdb workout

migrateup:
	migrate -path workout-be/db/migration -database "$(DB_URL)" -verbose up

migratedown:
	migrate -path db/migration -database "$(DB_URL)" -verbose down

migrateup1:
	migrate -path workout-be/db/migration -database "postgresql://root:NRbg8foygGuvOTzzzbqg@workout.car9zaosrys5.eu-west-2.rds.amazonaws.com:5432/workout" -verbose up

migratedown1:
	migrate -path db/migration -database "postgresql://root:NRbg8foygGuvOTzzzbqg@workout.car9zaosrys5.eu-west-2.rds.amazonaws.com:5432/workout" -verbose down
	
migrateuplocal:
	migrate -path workout-be/db/migration -database "postgresql://evilnis:Lon19ska83@localhost:5432/workout?sslmode=disable" -verbose up

migratedownlocal:
	migrate -path db/migration -database "postgresql://evilnis:Lon19ska83@localhost:5432/workout?sslmode=disable" -verbose down

sqlc:
	sqlc generate

test: 
	go test -v -cover ./...

server:
	go run main.go

dbdocs:
	dbdocs build loc doc/db.dbml

dbschema:
	dbml2sql --postgres - doc/schema.sql doc/db.dbml

mock:
	mockgen -package mockdb -destination db/mock/store.go github.com/DMonkey83/MyFitnessApp/workout-be/db/sqlc Store

.PHONY: migrateuplocal migratedown sqlc test dbdocs dbschema
		NRbg8foygGuvOTzzzbqg

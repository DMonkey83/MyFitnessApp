migrateup:
	migrate -path workout-be/db/migration -database "postgresql://evilnis:Lon19ska83@localhost:5432/workout?sslmode=disable" -verbose up

migratedown:
	migrate -path db/schema -database "postgresql://evilnis:Lon19ska83@localhost:5432/workout?sslmode=disable" -verbose down

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

.PHONY: migrateup migratedown sqlc test dbdocs dbschema

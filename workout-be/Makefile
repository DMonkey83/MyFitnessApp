migrateup:
	migrate -path workout-be/db/migration -database "postgresql://evilnis:Lon19ska83@localhost:5432/workout?sslmode=disable" -verbose up

migratedown:
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

.PHONY: migrateup migratedown sqlc test dbdocs dbschema

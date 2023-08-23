migrateup:
	migrate -path db/migration -database "postgresql://evilnis:Lon19ska83@localhost:5432/workout?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://evilnis:Lon19ska83@localhost:5432/workout?sslmode=disable" -verbose down

sqlc:
	sqlc generate

test: 
	go test -v -cover ./...

.PHONY: migrateup migratedown sqlc test

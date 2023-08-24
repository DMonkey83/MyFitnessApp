migrateup:
	migrate -path workout-be/db/migration -database "postgresql://evilnis:Lon19ska83@localhost:5432/workout?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://evilnis:Lon19ska83@localhost:5432/workout?sslmode=disable" -verbose down

sqlc:
	cd workout-be && sqlc generate

test: 
	cd workout-be && go test -v -cover ./...

.PHONY: migrateup migratedown sqlc test

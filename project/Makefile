current_dir = $(shell echo %cd%)

postgresup:
	docker run --name project7-postgres -e POSTGRES_PASSWORD=royce -e POSTGRES_USER=royce -e POSTGRES_DB=stock_info -p 5432:5432 -v D:\Workspace\JackRabbit\GO\project7\project\db-data\postgres:/var/lib/postgresql/data/ -d postgres:14.2

postgresrm:
	docker stop project7-postgres
	docker rm project7-postgres
	rm -Recurse -Force .\db-data\postgres

createdb:
	docker exec -it project7-postgres  createdb --username=royce --owner=royce stock_info

dropdb:
	docker exec -it project7-postgres  dropdb --username=royce stock_info

sqlc:
	docker run --rm -v $(current_dir):/src -w /src sqlc/sqlc generate

migrateup:
	migrate -path db/migrations/ -database "postgres://royce:royce@localhost:5432/stock_info?sslmode=disable" --verbose up

migratedown:
	migrate -path db/migrations/ -database "postgres://royce:royce@localhost:5432/stock_info?sslmode=disable" --verbose down


test:
	go test -v -cover ./...

.PHONY: postgresup postgresrm createdb dropdb test
 
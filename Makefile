current_dir = $(shell echo %cd%)

postgresup:
	docker run --name project7-postgres --network stockinfo-network -e POSTGRES_PASSWORD=royce -e POSTGRES_USER=royce -e POSTGRES_DB=stock_info -p 5432:5432 -v D:\Workspace\JackRabbit\GO\project7\project\db-data\postgres:/var/lib/postgresql/data/ -d postgres:14.2

postgresrm:
	docker stop project7-postgres
	docker rm project7-postgres
	rm -Recurse -Force .\db-data\postgres

createdb:
	docker exec -it project7-postgres  createdb --username=royce --owner=royce stock_info

dropdb:
	docker exec -it project7-postgres  dropdb --username=royce stock_info

sqlc:
	docker run --rm -v $(current_dir)/project:/src -w /src sqlc/sqlc generate

migrateup:
	migrate -path project/db/migrations/ -database "postgres://royce:royce@localhost:5432/stock_info?sslmode=disable" --verbose up

migratedown:
	migrate -path project/db/migrations/ -database "postgres://royce:royce@localhost:5432/stock_info?sslmode=disable" --verbose down

test:
	go test -v -cover ./project/...
	go test -v -cover ./shared/...
	go test -v -cover ./api/...
server:
	go run main.go

# mockgen需要依賴go.mod  你的執行指令目錄或父目錄必須包含go.mod,  所以無法在root 目錄執行  因為只有go.work
mock:
	mockgen -package mockdb -destination db/mock/store.go github.com/RoyceAzure/go-stockinfo-project/db/sqlc Store

.PHONY: postgresup postgresrm createdb dropdb test server mock
 
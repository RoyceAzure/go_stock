current_dir = $(shell echo %cd%)
DB_URL_LOCAL=postgres://royce:royce@localhost:5432/stock_info?sslmode=disable
DB_URL_AWS=postgres://royce:gqD2yhIOpUpuwK6IX6xz@stockinfo.cblayv8xneas.ap-northeast-1.rds.amazonaws.com:5432/stockinfo

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

awsmigrateup:
	migrate -path project/db/migrations/ -database $(DB_URL_AWS) --verbose up

migrateup:
	migrate -path project/db/migrations/ -database $(DB_URL_LOCAL) --verbose up

migratedown:
	migrate -path project/db/migrations/ -database $(DB_URL_LOCAL) --verbose down

test:
	go test -v -cover ./project/...
	go test -v -cover ./shared/...
	go test -v -cover ./api/...
server:
	go run main.go

# mockgen需要依賴go.mod  你的執行指令目錄或父目錄必須包含go.mod,  所以無法在root 目錄執行  因為只有go.work
mock:
	cd  project/
	mockgen -package mockdb -destination db/mock/store.go github.com/RoyceAzure/go-stockinfo-project/db/sqlc Store

db_docs:
	docker run --rm -v $(current_dir)/doc:/app/data -w /app/data node_docs dbdocs build ./db.dbml

db_schema:
	docker run --rm -v $(current_dir)/doc:/app/data -w /app/data node_docs dbml2sql --postgres -o schema.sql db.dbml


.PHONY: postgresup postgresrm createdb dropdb test server mock awsmigrateup db_docs db_schema
 
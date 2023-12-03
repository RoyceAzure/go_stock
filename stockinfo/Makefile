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
	docker run --rm -v $(current_dir)/project:/src -w /src sqlc/sqlc:latest generate

awsmigrateup:
	migrate -path project/db/migrations/ -database $(DB_URL_AWS) --verbose up

migrateup:
	migrate -path project/db/migrations/ -database $(DB_URL_LOCAL) --verbose up

migrateupto:
	migrate -path project/db/migrations/ -database $(DB_URL_LOCAL) --verbose up	to $(stage)

migratedownto:
	migrate -path project/db/migrations/ -database $(DB_URL_LOCAL) --verbose down to $(stage)

migrateup1:
	migrate -path project/db/migrations/ -database $(DB_URL_LOCAL) --verbose up	1

migrateudown1:
	migrate -path project/db/migrations/ -database $(DB_URL_LOCAL) --verbose down 1

migratedownto:
	migrate -path project/db/migrations/ -database $(DB_URL_LOCAL) --verbose down to $(stage)

migratedown:
	migrate -path project/db/migrations/ -database $(DB_URL_LOCAL) --verbose down

migrateforce:
	migrate -path project/db/migrations/ -database $(DB_URL_LOCAL) --verbose force $(stage)

new_migration:
	migrate create -ext sql -dir db/migrations -seq $(name)

redis:
	docker run --name redis -p 6379:6379 -d redis:7.2.2-alpine

test:
	go test -v -cover -short ./project/...
	go test -v -cover -short ./shared/...
	go test -v -cover -short ./api/...
server:
	go run main.go

# mock:
# 	mockgen -source="./distributor.go" -package=mockwk -destination="mock/distributor_mock.go"
# 	mockgen  . Store

db_docs:
	docker run --rm -v $(current_dir)/doc:/app/data -w /app/data node_docs dbdocs build ./db.dbml

db_schema:
	docker run --rm -v $(current_dir)/doc:/app/data -w /app/data node_docs dbml2sql --postgres -o schema.sql db.dbml

protoc:
	powershell -Command "Remove-Item -Path 'api/pb/*.go' -Force"
	powershell -Command "Remove-Item -Path 'doc/swagger/*.swagger.json' -Force"
	protoc   -I=C:/Users/royce/include -I=./proto --grpc-gateway_out api/pb \
	--proto_path=proto  --go_out=api/pb  --go_opt=paths=source_relative  --grpc-gateway_opt=paths=source_relative \
	--go-grpc_out=api/pb --go-grpc_opt=paths=source_relative --openapiv2_out doc/swagger --openapiv2_opt=allow_merge=true,merge_file_name=stock_info \
	proto/*.proto
	statik -src=./doc/swagger -dest=./doc -f

evans:
	docker run -it --rm -v $(current_dir):/mount:ro ghcr.io/ktr0731/evans:latest --host host.docker.internal --port 9090 -r repl

.PHONY: postgresup postgresrm createdb dropdb test server awsmigrateup db_docs db_schema protoc evans redis new_migration migrateforce migrateupto migratedownto
 
current_dir = $(shell echo %cd%)
STOCKINFO_DB_URL_LOCAL=postgres://royce:royce@localhost:5432/stock_info?sslmode=disable
SCHEDULER_DB_URL_LOCAL=postgres://royce:royce@localhost:5432/stock_info_scheduler?sslmode=disable
DISTRIBUTOR_DB_URL_LOCAL=postgres://royce:royce@localhost:5432/stock_info_distributor?sslmode=disable
DB_URL_AWS=postgres://royce:gqD2yhIOpUpuwK6IX6xz@stockinfo.cblayv8xneas.ap-northeast-1.rds.amazonaws.com:5432/stockinfo
migrateup:
	migrate -path stockinfo/repository/db/migrations/ -database $(STOCKINFO_DB_URL_LOCAL) --verbose up
	migrate -path scheduler/repository/migrations/ -database $(SCHEDULER_DB_URL_LOCAL) --verbose up
	migrate -path distributor/repository/db/migrations/ -database $(DISTRIBUTOR_DB_URL_LOCAL) --verbose up
test:
	go test -v -cover -short ./broker/...
	go test -v -cover -short ./distributor/...
	go test -v -cover -short ./logger/...
	go test -v -cover -short ./scheduler/...
	go test -v -cover -short ./stockinfo/...
applylocalk8s:
	kubectl apply -f ./local_k8s/
.PHONY:	migrateup test applylocalk8s
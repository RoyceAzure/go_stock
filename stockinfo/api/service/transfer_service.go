package service

import (
	"context"

	db "github.com/RoyceAzure/go-stockinfo/repository/db/sqlc"
	remote_repo "github.com/RoyceAzure/go-stockinfo/repository/remote_repo"
)

type ITransferService interface {
	StockTransfer(ctx context.Context, arg TransferStockServiceParams) error
}

type TransferService struct {
	store       db.Store
	schdulerDao remote_repo.SchdulerInfoDao
}

func NewTransferService(store db.Store) ITransferService {
	return &TransferService{
		store: store,
	}
}

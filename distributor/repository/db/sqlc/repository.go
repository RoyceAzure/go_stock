package repository

type DistributorDao interface {
	Querier
}

type SQLDistributorDao struct {
	*Queries
}

func NewSQLDistributorDao(db DBTX) DistributorDao {
	return &SQLDistributorDao{
		New(db),
	}
}

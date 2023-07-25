package postgres

import (
	"app/config"
	"app/storage"
	"context"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"
)

type store struct {
	db       *pgxpool.Pool
	product  *productRepo
	category *categoryRepo
	filial   *filialRepo
	coming   *comingRepo
	remainder *remainderRepo
	comingProducts *comingProductsRepo
}

func NewConnectionPostgres(cfg *config.Config) (storage.StorageI, error) {
	connect, err := pgxpool.ParseConfig(fmt.Sprintf(
		"host=%s user=%s dbname=%s password=%s port=%d sslmode=disable",
		cfg.PostgresHost,
		cfg.PostgresUser,
		cfg.PostgresDatabase,
		cfg.PostgresPassword,
		cfg.PostgresPort,
	))

	if err != nil {
		return nil, err
	}

	connect.MaxConns = cfg.PostgresMaxConnection

	pgxpool, err := pgxpool.ConnectConfig(context.Background(), connect)
	if err != nil {
		return nil, err
	}

	return &store{
		db: pgxpool,
	}, nil
}

func (s *store) Close() {
	s.db.Close()
}

func (s *store) Product() storage.ProductRepoI {

	if s.product == nil {
		s.product = NewProductRepo(s.db)
	}

	return s.product
}

func (s *store) Category() storage.CategoryRepoI {

	if s.category == nil {
		s.category = NewCategoryRepo(s.db)
	}

	return s.category
}

func (s *store) Filial() storage.FilialRepoI {

	if s.filial == nil {
		s.filial = NewFilialRepo(s.db)
	}

	return s.filial
}

func (s *store) Coming() storage.ComingRepoI {

	if s.coming == nil {
		s.coming = NewComingRepo(s.db)
	}

	return s.coming
}


func (s *store) Remainder() storage.RemainderRepoI {

	if s.remainder == nil {
		s.remainder = NewRemainderRepo(s.db)
	}

	return s.remainder
}


func (s *store) ComingProducts() storage.ComingProductsRepoI {

	if s.comingProducts == nil {
		s.comingProducts = NewComingProductsRepo(s.db)
	}

	return s.comingProducts
}
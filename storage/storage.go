package storage

import (
	"app/api/models"
	"context"
)

type StorageI interface {
	Close()
	Product() ProductRepoI
	Category() CategoryRepoI
	Filial() FilialRepoI
	Coming()  ComingRepoI
	Remainder() RemainderRepoI
	ComingProducts() ComingProductsRepoI
}

type CategoryRepoI interface {
	Create(context.Context, *models.CreateCategory) (string, error)
	GetByID(context.Context, *models.CategoryPrimaryKey) (*models.Category, error)
	GetList(context.Context, *models.CategoryGetListRequest) (*models.CategoryGetListResponse, error)
	Update(context.Context, *models.UpdateCategory) (int64, error)
	Delete(context.Context, *models.CategoryPrimaryKey) error
}

type ProductRepoI interface {
	Create(context.Context, *models.CreateProduct) (string, error)
	GetByID(context.Context, *models.ProductPrimaryKey) (*models.Product, error)
	// GetByBarcode(context.Context, *models.ProductBarcodeKey) (*models.Product, error)
	GetList(context.Context, *models.ProductGetListRequest) (*models.ProductGetListResponse, error)
	Update(context.Context, *models.UpdateProduct) (int64, error)
	Patch(context.Context, *models.PatchRequest) (int64, error)
	Delete(context.Context, *models.ProductPrimaryKey) error
}


type FilialRepoI interface {
	Create(context.Context, *models.CreateFilial) (string, error)
	GetByID(context.Context, *models.FilialPrimaryKey) (*models.Filial, error)
	GetList(context.Context, *models.FilialGetListRequest) (*models.FilialGetListResponse, error)
	Update(context.Context, *models.UpdateFilial) (int64, error)
	Patch(context.Context, *models.PatchRequest) (int64, error)
	Delete(context.Context, *models.FilialPrimaryKey) error
}

type ComingRepoI interface {
	Create(context.Context, *models.CreateComing) (string, error)
	GetByID(context.Context, *models.ComingPrimaryKey) (*models.Coming, error)
	GetList(context.Context, *models.ComingGetListRequest) (*models.ComingGetListResponse, error)
	Update(context.Context, *models.UpdateComing) (int64, error)
	Patch(context.Context, *models.PatchRequest) (int64, error)
	Delete(context.Context, *models.ComingPrimaryKey) error
}

type RemainderRepoI interface {
	Create(context.Context, *models.CreateRemainder) (string, error)
	GetByID(context.Context, *models.RemainderPrimaryKey) (*models.Remainder, error)
	GetList(context.Context, *models.RemainderGetListRequest) (*models.RemainderGetListResponse, error)
	Update(context.Context, *models.UpdateRemainder) (int64, error)
	Patch(context.Context, *models.PatchRequest) (int64, error)
	Delete(context.Context, *models.RemainderPrimaryKey) error
	AddProduct(ctx context.Context, req *models.RespProduct) (int64, error) }


type ComingProductsRepoI interface {
	Create(context.Context, *models.CreateComingProducts) (string, error)
	GetByID(context.Context, *models.ComingProductsPrimaryKey) (*models.ComingProducts, error)
	GetList(context.Context, *models.ComingProductsGetListRequest) (*models.ComingProductsGetListResponse, error)
	Update(context.Context, *models.UpdateComingProducts) (int64, error)
	Patch(context.Context, *models.PatchRequest) (int64, error)
	Delete(context.Context, *models.ComingProductsPrimaryKey) error
	GetByComingID(ctx context.Context, req *models.ComingPrimaryKey) (*models.ComingProducts, error)
}


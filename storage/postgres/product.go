package postgres

import (
	"app/api/models"
	"app/pkg/helper"
	"context"
	"database/sql"
	"errors"
	"fmt"

	uuid "github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
)

type productRepo struct {
	db *pgxpool.Pool
}

func NewProductRepo(db *pgxpool.Pool) *productRepo {
	return &productRepo{
		db: db,
	}
}

func (r *productRepo) Create(ctx context.Context, req *models.CreateProduct) (string, error) {

	trx, err := r.db.Begin(ctx)
	if err != nil {
		return "", nil
	}

	defer func() {
		if err != nil {
			trx.Rollback(ctx)
		} else {
			trx.Commit(ctx)
		}
	}()

	var (
		id    = uuid.New().String()
		query string
	)

	query = `
		INSERT INTO product(id, name, price, barcode, category_id, updated_at)
		VALUES ($1, $2, $3, $4, $5, NOW())
	`

	fmt.Printf("%+v", req)
	_, err = trx.Exec(ctx, query,
		id,
		req.Name,
		req.Price,
		req.Barcode,
		helper.NewNullString(req.CategoryId),
	)

	if err != nil {
		return "", err
	}

	return id, nil
}
// func (r *productRepo) GetByBarcode(ctx context.Context, req *models.ProductBarcodeKey) (*models.Product, error) {

// 	var (
// 		query string
// 		id         sql.NullString
// 		name       sql.NullString
// 		price      sql.NullFloat64
// 		categoryID sql.NullString
// 		barcode    sql.NullString
// 		createdAt  sql.NullString
// 		updatedAt  sql.NullString
// 	)

// 	query = `
// 		select 
// 			id,
// 			name,
// 			price,
// 			barcode,
// 			category_id,
// 			created_at,
// 			updated_at
// 		FROM product 
// 		WHERE barcode = $1`

// 	fmt.Println(req.Barcode)
// 	err := r.db.QueryRow(ctx, query, req.Barcode).Scan(
// 		&id,
// 		&name,
// 		&price,
// 		&barcode,
// 		&categoryID,
// 		&createdAt,
// 		&updatedAt,
// 	)

// 	if err != nil {
// 		return nil, err
// 	}
// 	return &models.Product{
// 		Id:         id.String,
// 		Name:       name.String,
// 		Price:      price.Float64,
// 		Barcode:    barcode.String,
// 		CategoryId: categoryID.String,
// 		CreatedAt:  createdAt.String,
// 		UpdatedAt:  updatedAt.String,
// 	}, nil
// }

func (r *productRepo) GetByID(ctx context.Context, req *models.ProductPrimaryKey) (*models.Product, error) {

	var (
		query string
		id         sql.NullString
		name       sql.NullString
		price      sql.NullFloat64
		categoryID sql.NullString
		barcode    sql.NullString
		createdAt  sql.NullString
		updatedAt  sql.NullString
	)

	query = `
		select 
			id,
			name,
			price,
			barcode,
			category_id,
			created_at,
			updated_at
		FROM product 
		WHERE id = $1`

	fmt.Println(req.Id)
	err := r.db.QueryRow(ctx, query, req.Id).Scan(
		&id,
		&name,
		&price,
		&barcode,
		&categoryID,
		&createdAt,
		&updatedAt,
	)

	if err != nil {
		return nil, err
	}
	return &models.Product{
		Id:         id.String,
		Name:       name.String,
		Price:      price.Float64,
		Barcode:    barcode.String,
		CategoryId: categoryID.String,
		CreatedAt:  createdAt.String,
		UpdatedAt:  updatedAt.String,
	}, nil
}

func (r *productRepo) GetList(ctx context.Context, req *models.ProductGetListRequest) (*models.ProductGetListResponse, error) {

	var (
		resp   = &models.ProductGetListResponse{}
		query  string
		where  = " WHERE TRUE"
		offset = " OFFSET 0"
		limit  = " LIMIT 10"
	)

	query = `
		SELECT
			COUNT(*) OVER(),
			id,
			name,
			price,
			barcode,
			category_id,
			created_at,
			updated_at
		FROM product
	`

	if req.Offset > 0 {
		offset = fmt.Sprintf(" OFFSET %d", req.Offset)
	}

	if req.Limit > 0 {
		limit = fmt.Sprintf(" LIMIT %d", req.Limit)
	}

	if req.SearchByName != "" {
		where += ` AND name ILIKE '%' || '` + req.SearchByName + `' || '%'`
	}

	if req.SearchByBarcode != "" {
		where += ` AND barcode ILIKE '%' || '` + req.SearchByBarcode + `' || '%'`
	}

	query += where + offset + limit

	fmt.Println(query)
	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var (
			id         sql.NullString
			name       sql.NullString
			price      sql.NullFloat64
			barcode    sql.NullString
			categoryID sql.NullString
			createdAt  sql.NullString
			updatedAt  sql.NullString
		)

		err := rows.Scan(
			&resp.Count,
			&id,
			&name,
			&price,
			&barcode,
			&categoryID,
			&createdAt,
			&updatedAt,
		)

		if err != nil {
			return nil, err
		}

		resp.Products = append(resp.Products, &models.Product{
			Id:         id.String,
			Name:       name.String,
			Price:      price.Float64,
			Barcode:    barcode.String,
			CategoryId: categoryID.String,
			CreatedAt:  createdAt.String,
			UpdatedAt:  updatedAt.String,
		})
	}

	return resp, nil
}

func (r *productRepo) Update(ctx context.Context, req *models.UpdateProduct) (int64, error) {

	trx, err := r.db.Begin(ctx)
	if err != nil {
		return 0, nil
	}

	defer func() {
		if err != nil {
			trx.Rollback(ctx)
		} else {
			trx.Commit(ctx)
		}
	}()

	var (
		query  string
		params map[string]interface{}
	)

	query = `
		UPDATE
			product
		SET
			name = :name,
			price = :price,
			barcode = :barcode,
			category_id = :category_id,
			updated_at = NOW()
		WHERE id = :id
	`

	params = map[string]interface{}{
		"id":          req.Id,
		"name":        req.Name,
		"price":       req.Price,
		"barcode":     req.Barcode,
		"category_id": helper.NewNullString(req.CategoryId),
	}

	query, args := helper.ReplaceQueryParams(query, params)

	result, err := trx.Exec(ctx, query, args...)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected(), nil
}

func (r *productRepo) Patch(ctx context.Context, req *models.PatchRequest) (int64, error) {

	var (
		query string
		set   string
	)

	if len(req.Fields) <= 0 {
		return 0, errors.New("no fields")
	}

	for key := range req.Fields {
		set += fmt.Sprintf(" %s = :%s, ", key, key)
	}

	query = `
		UPDATE
			product
		SET ` + set + ` updated_at = now()
		WHERE id = :id
	`

	req.Fields["id"] = req.ID

	query, args := helper.ReplaceQueryParams(query, req.Fields)
	result, err := r.db.Exec(ctx, query, args...)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected(), nil
}

func (r *productRepo) Delete(ctx context.Context, req *models.ProductPrimaryKey) error {

	_, err := r.db.Exec(ctx, "DELETE FROM product WHERE id = $1", req.Id)
	if err != nil {
		return err
	}

	return nil
}

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

type comingProductsRepo struct {
	db *pgxpool.Pool
}

func NewComingProductsRepo(db *pgxpool.Pool) *comingProductsRepo {
	return &comingProductsRepo{
		db: db,
	}
}

func (r *comingProductsRepo) Create(ctx context.Context, req *models.CreateComingProducts) (string, error) {

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
		id          = uuid.New().String()
		query       string
		productData models.Product
	)

	if err != nil {
		return "", err
	}

	fmt.Println("%+v", productData)
	query = `
		INSERT INTO coming_products(id,coming_id, category_id, name, barcode, amount, price,total_price, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, NOW())
	`

	fmt.Printf("%+v", req)
	_, err = trx.Exec(ctx, query,
		id,
		req.ComingID,
		req.CategoryID,
		req.Name,
		req.Barcode,
		req.Amount,
		req.Price,
		req.Price*float64(req.Amount),
	)

	if err != nil {
		return "", err
	}

	return id, nil
}

func (r *comingProductsRepo) GetByID(ctx context.Context, req *models.ComingProductsPrimaryKey) (*models.ComingProducts, error) {

	var (
		query       string
		id          sql.NullString
		comingID    sql.NullString
		category_id sql.NullString
		name        sql.NullString
		barcode     sql.NullString
		amount      sql.NullInt64
		price       sql.NullFloat64
		total_price sql.NullFloat64
		createdAt   sql.NullString
		updatedAt   sql.NullString
	)

	query = `
		select 
			id,
			coming_id,
			category_id,
			name,
			barcode,
			amount,
			price,
			total_price,
			created_at,
			updated_at
		FROM coming_products 
		WHERE id = $1`

	fmt.Println(req.Id)
	err := r.db.QueryRow(ctx, query, req.Id).Scan(
		&id,
		&comingID,
		&category_id,
		&name,
		&barcode,
		&amount,
		&price,
		&total_price,
		&createdAt,
		&updatedAt,
	)

	if err != nil {
		return nil, err
	}
	return &models.ComingProducts{
		Id:         id.String,
		ComingID:   comingID.String,
		CategoryID: category_id.String,
		Name:       name.String,
		Barcode:    barcode.String,
		Amount:     int(amount.Int64),
		Price:      price.Float64,
		TotalPrice: total_price.Float64,
		CreatedAt:  createdAt.String,
		UpdatedAt:  updatedAt.String,
	}, nil
}

func (r *comingProductsRepo) GetList(ctx context.Context, req *models.ComingProductsGetListRequest) (*models.ComingProductsGetListResponse, error) {

	var (
		resp   = &models.ComingProductsGetListResponse{}
		query  string
		where  = " WHERE TRUE"
		offset = " OFFSET 0"
		limit  = " LIMIT 10"
	)

	query = `
		SELECT
			COUNT(*) OVER(),
			id,
			coming_id,
			category_id,
			name,
			barcode,
			amount,
			price,
			total_price,
			created_at,
			updated_at
		FROM coming_products
	`

	if req.Offset > 0 {
		offset = fmt.Sprintf(" OFFSET %d", req.Offset)
	}

	if req.Limit > 0 {
		limit = fmt.Sprintf(" LIMIT %d", req.Limit)
	}

	if req.SearchByCategoryID != "" {
		where += ` AND category_id ILIKE '%' || '` + req.SearchByCategoryID + `' || '%'`
	}
	if req.SearchByCategoryID != "" {
		where += ` AND barcode ILIKE '%' || '` + req.SearchByBarcode + `' || '%'`
	}
	if req.SearchBYComingID != ""{
		where += ` AND coming_id ILIKE '%' || '` + req.SearchBYComingID + `' || '%'`
	}

	query += where + offset + limit

	fmt.Println(query)
	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	fmt.Println("\n\n\n\n")

	for rows.Next() {
		var (
			id          sql.NullString
			comingID    sql.NullString
			category_id sql.NullString
			name        sql.NullString
			barcode     sql.NullString
			amount      sql.NullInt64
			price       sql.NullFloat64
			total_price sql.NullFloat64
			createdAt   sql.NullString
			updatedAt   sql.NullString
		)

		err := rows.Scan(
			&resp.Count,
			&id,
			&comingID,
			&category_id,
			&name,
			&barcode,
			&amount,
			&price,
			&total_price,
			&createdAt,
			&updatedAt,
		)

		if err != nil {
			return nil, err
		}

		resp.ComingProducts = append(resp.ComingProducts, &models.ComingProducts{
			Id:         id.String,
			ComingID:   comingID.String,
			CategoryID: category_id.String,
			Name:       name.String,
			Barcode:    barcode.String,
			Amount:     int(amount.Int64),
			Price:      price.Float64,
			TotalPrice: total_price.Float64,
			CreatedAt:  createdAt.String,
			UpdatedAt:  updatedAt.String,
		})
	}

	return resp, nil
}
func (r *comingProductsRepo) GetByComingID(ctx context.Context, req *models.ComingPrimaryKey) (*models.ComingProducts, error) {

	var (
		query       string
		id          sql.NullString
		comingID    sql.NullString
		category_id sql.NullString
		name        sql.NullString
		barcode     sql.NullString
		amount      sql.NullInt64
		price       sql.NullFloat64
		total_price sql.NullFloat64
		createdAt   sql.NullString
		updatedAt   sql.NullString
	)

	query = `
		select 
			id,
			coming_id,
			category_id,
			name,
			barcode,
			amount,
			price,
			total_price,
			created_at,
			updated_at
		FROM coming_products 
		WHERE coming_id = $1`

	fmt.Println(req.Id)
	err := r.db.QueryRow(ctx, query, req.Id).Scan(
		&id,
		&comingID,
		&category_id,
		&name,
		&barcode,
		&amount,
		&price,
		&total_price,
		&createdAt,
		&updatedAt,
	)

	if err != nil {
		return nil, err
	}
	return &models.ComingProducts{
		Id:         id.String,
		ComingID:   comingID.String,
		CategoryID: category_id.String,
		Name:       name.String,
		Barcode:    barcode.String,
		Amount:     int(amount.Int64),
		Price:      price.Float64,
		TotalPrice: total_price.Float64,
		CreatedAt:  createdAt.String,
		UpdatedAt:  updatedAt.String,
	}, nil
}

func (r *comingProductsRepo) Update(ctx context.Context, req *models.UpdateComingProducts) (int64, error) {

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
			coming_products
		SET
			coming_id = :coming_id
			category_id = :category_id,
			name = :name,
			barcode = :barcode,
			amount = :amount,
			price = :price,
			total_price =: total_price,
			updated_at = NOW()
		WHERE id = :id
	`

	params = map[string]interface{}{
		"id":          req.Id,
		"coming_id":   req.CategoryID,
		"category_id": req.CategoryID,
		"name":  req.Name,
		"barcode":     req.Barcode,
		"amount":      req.Amount,
		"price":       req.Price,
		"total_price": req.Price * float64(req.Amount),
	}

	query, args := helper.ReplaceQueryParams(query, params)

	result, err := trx.Exec(ctx, query, args...)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected(), nil
}

func (r *comingProductsRepo) Patch(ctx context.Context, req *models.PatchRequest) (int64, error) {

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
			coming_products
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

func (r *comingProductsRepo) Delete(ctx context.Context, req *models.ComingProductsPrimaryKey) error {

	_, err := r.db.Exec(ctx, "DELETE FROM coming_products WHERE id = $1", req.Id)
	if err != nil {
		return err
	}

	return nil
}

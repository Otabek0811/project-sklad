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

type remainderRepo struct {
	db *pgxpool.Pool
}

func NewRemainderRepo(db *pgxpool.Pool) *remainderRepo {
	return &remainderRepo{
		db: db,
	}
}

func (r *remainderRepo) Create(ctx context.Context, req *models.CreateRemainder) (string, error) {

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
		INSERT INTO remainder(id, filial_id, category_id, name, barcode, amount, price,total_price, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, NOW())
	`

	fmt.Printf("%+v", req)
	_, err = trx.Exec(ctx, query,
		id,
		req.FilialID,
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

func (r *remainderRepo) GetByID(ctx context.Context, req *models.RemainderPrimaryKey) (*models.Remainder, error) {

	var (
		query       string
		id          sql.NullString
		filial_id   sql.NullString
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
			filial_id,
			category_id,
			name,
			barcode,
			amount,
			price,
			total_price,
			created_at,
			updated_at
		FROM remainder 
		WHERE id = $1`

	fmt.Println(req.Id)
	err := r.db.QueryRow(ctx, query, req.Id).Scan(
		&id,
		&filial_id,
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
	return &models.Remainder{
		Id:         id.String,
		FilialID:   filial_id.String,
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

func (r *remainderRepo) GetList(ctx context.Context, req *models.RemainderGetListRequest) (*models.RemainderGetListResponse, error) {

	var (
		resp   = &models.RemainderGetListResponse{}
		query  string
		where  = " WHERE TRUE"
		offset = " OFFSET 0"
		limit  = " LIMIT 10"
	)

	query = `
		SELECT
			COUNT(*) OVER(),
			id,
			filial_id,
			category_id,
			name,
			barcode,
			amount,
			price,
			total_price,
			created_at,
			updated_at
		FROM remainder
	`

	if req.Offset > 0 {
		offset = fmt.Sprintf(" OFFSET %d", req.Offset)
	}

	if req.Limit > 0 {
		limit = fmt.Sprintf(" LIMIT %d", req.Limit)
	}

	if req.SearchByFilial != "" {
		where += ` AND filial_id ILIKE '%' || '` + req.SearchByFilial + `' || '%'`
	}
	if req.SearchByCategory != "" {
		where += ` AND category_id ILIKE '%' || '` + req.SearchByCategory + `' || '%'`
	}
	if req.SearchByBarcode != "" {
		where += ` AND barcode ILIKE '%' || '` + req.SearchByBarcode + `' || '%'`
	}

	query += where + offset + limit

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var (
			id          sql.NullString
			filial_id   sql.NullString
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
			&filial_id,
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

		resp.Remainders = append(resp.Remainders, &models.Remainder{
			Id:         id.String,
			FilialID:   filial_id.String,
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

func (r *remainderRepo) Update(ctx context.Context, req *models.UpdateRemainder) (int64, error) {

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
			remainder
		SET
			filial_id = :filial_id,
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
		"filial_id":   req.FilialID,
		"category_id": req.CategoryID,
		"name":        req.Name,
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

func (r *remainderRepo) Patch(ctx context.Context, req *models.PatchRequest) (int64, error) {

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
			remainder
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

func (r *remainderRepo) Delete(ctx context.Context, req *models.RemainderPrimaryKey) error {

	_, err := r.db.Exec(ctx, "DELETE FROM remainder WHERE id = $1", req.Id)
	if err != nil {
		return err
	}

	return nil
}

func (r *remainderRepo) AddProduct(ctx context.Context, req *models.RespProduct) (int64, error) {
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
		checkQuery string
		query      string
		id         sql.NullString
		count   int
	)

	checkQuery = `
		SELECT 
			id,
			amount
		FROM remainder
		where filial_id=$1 and barcode=$2
	`

	err = trx.QueryRow(ctx, checkQuery,
		req.FilialID,
		req.Barcode,
	).Scan(&id,&count)

	fmt.Println(err)
	err=nil


	if id.String != "" {

		query = `
		UPDATE 
			remainder
		SET
			filial_id = $1,
			category_id = $2,
			name = $3,
			barcode = $4,
			amount = $5,
			price = $6,
			total_price = $7,
			updated_at = NOW()
		WHERE id = $8 
	`



		fmt.Println(count)
		count+=req.Amount
		fmt.Println(count,req.Amount)
		result, err := trx.Exec(ctx, query,
			req.FilialID,
			req.CategoryID,
			req.Name,
			req.Barcode,
			count,
			req.Price,
			req.Price*float64(count),
			id,
		)
		if err != nil {
			return 0, err
		}

		return result.RowsAffected(), nil

	}else {
		fmt.Printf("%+v",req)
		var(
			id = uuid.New().String()
		)
		query = `
		INSERT INTO remainder(id, filial_id, category_id, name, barcode, amount, price,total_price, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, NOW())
	`
	_, err = trx.Exec(ctx, query,
		id,
		req.FilialID,
		helper.NewNullString(req.CategoryID),
		req.Name,
		req.Barcode,
		req.Amount,
		req.Price,
		req.Price*float64(req.Amount),
	)
	fmt.Println("1111111111111111111111111")

	fmt.Println(err)
	if err != nil {
		return 0, err
	}

	return 1, nil
	}

}

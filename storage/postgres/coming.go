package postgres

import (
	"app/api/models"
	"app/pkg/helper"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	uuid "github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
)

type comingRepo struct {
	db *pgxpool.Pool
}

func NewComingRepo(db *pgxpool.Pool) *comingRepo {
	return &comingRepo{
		db: db,
	}
}

func (r *comingRepo) Create(ctx context.Context, req *models.CreateComing) (string, error) {

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
		INSERT INTO coming(id, coming_id, filial_id, status, updated_at)
		VALUES ($1, $2, $3, $4, NOW())
	`

	fmt.Printf("%+v", req)


	if req.Status=="finished"{
		query = `
		INSERT INTO coming(id, coming_id, filial_id, date_time, status, updated_at)
		VALUES ($1, $2, $3, NOW(), $4, NOW())
	`
		_, err = trx.Exec(ctx, query,
			id,
			req.ComingID,
			req.FilialID,
			req.Status,
		)
	}else{
		_, err = trx.Exec(ctx, query,
			id,
			req.ComingID,
			req.FilialID,
			req.Status,
		)
	}
	

	if err != nil {
		return "", err
	}

	return id, nil
}

func (r *comingRepo) GetByID(ctx context.Context, req *models.ComingPrimaryKey) (*models.Coming, error) {

	var (
		query     string
		id        sql.NullString
		coming_id sql.NullString
		filial_id sql.NullString
		date_time sql.NullString
		status    sql.NullString
		createdAt sql.NullString
		updatedAt sql.NullString
	)

	query = `
		select 
			id,
			coming_id,
			filial_id,
			date_time,
			status,
			created_at,
			updated_at
		FROM coming 
		WHERE id = $1`

	fmt.Println(req.Id)
	err := r.db.QueryRow(ctx, query, req.Id).Scan(
		&id,
		&coming_id,
		&filial_id,
		&date_time,
		&status,
		&createdAt,
		&updatedAt,
	)

	if err != nil {
		return nil, err
	}
	return &models.Coming{
		Id:        id.String,
		ComingID:  coming_id.String,
		FilialID:  filial_id.String,
		DateTime:  date_time.String,
		Status:    status.String,
		CreatedAt: createdAt.String,
		UpdatedAt: updatedAt.String,
	}, nil
}

func (r *comingRepo) GetList(ctx context.Context, req *models.ComingGetListRequest) (*models.ComingGetListResponse, error) {

	var (
		resp   = &models.ComingGetListResponse{}
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
			filial_id,
			date_time,
			status,
			created_at,
			updated_at
		FROM coming
	`

	if req.Offset > 0 {
		offset = fmt.Sprintf(" OFFSET %d", req.Offset)
	}

	if req.Limit > 0 {
		limit = fmt.Sprintf(" LIMIT %d", req.Limit)
	}

	if req.Search_by_comingID != "" {
		where += ` AND coming_id ILIKE '%' || '` + req.Search_by_comingID + `' || '%'`
	}

	if req.Search_by_filial != "" {
		where += ` AND filial_id ILIKE '%' || '` + req.Search_by_filial + `' || '%'`
	}

	query += where + offset + limit

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var (
			id        sql.NullString
			coming_id sql.NullString
			filial_id sql.NullString
			date_time sql.NullString
			status    sql.NullString
			createdAt sql.NullString
			updatedAt sql.NullString
		)

		err := rows.Scan(
			&resp.Count,
			&id,
			&coming_id,
			&filial_id,
			&date_time,
			&status,
			&createdAt,
			&updatedAt,
		)

		if err != nil {
			return nil, err
		}

		resp.Comings = append(resp.Comings, &models.Coming{
			Id:        id.String,
			ComingID:  coming_id.String,
			FilialID:  filial_id.String,
			DateTime:  date_time.String,
			Status:    status.String,
			CreatedAt: createdAt.String,
			UpdatedAt: updatedAt.String,
		})
	}

	return resp, nil
}

func (r *comingRepo) Update(ctx context.Context, req *models.UpdateComing) (int64, error) {

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
			coming
		SET
			coming_id = :coming_id,
			filial_id = :filial_id,
			date_time = :date_time,
			status = :status,
			updated_at = NOW()
		WHERE id = :id
	`

	if req.Status == "finished" {
		params = map[string]interface{}{
			"id":        req.Id,
			"coming_id": req.ComingID,
			"filial_id": req.FilialID,
			"date_time": time.Now(),
			"status":    req.Status,
		}
	} else {

		params = map[string]interface{}{
			"id":        req.Id,
			"coming_id": req.ComingID,
			"filial_id": req.FilialID,
			"date_time": req.DateTime,
			"status":    req.Status,
		}

	}
	// params = map[string]interface{}{
	// 	"id":        req.Id,
	// 	"coming_id": req.ComingID,
	// 	"filial_id": req.FilialID,
	// 	"date_time": req.DateTime,
	// 	"status":    req.Status,
	// }

	query, args := helper.ReplaceQueryParams(query, params)

	result, err := trx.Exec(ctx, query, args...)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected(), nil
}

func (r *comingRepo) Patch(ctx context.Context, req *models.PatchRequest) (int64, error) {

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
			coming
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

func (r *comingRepo) Delete(ctx context.Context, req *models.ComingPrimaryKey) error {

	_, err := r.db.Exec(ctx, "DELETE FROM coming WHERE id = $1", req.Id)
	if err != nil {
		return err
	}

	return nil
}

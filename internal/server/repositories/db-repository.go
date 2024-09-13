package repositories

import (
	"context"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"screamer/internal/common/metric"
	"screamer/internal/server/db"
)

type DBRepository struct {
	pool *pgxpool.Pool
}

func (db *DBRepository) GetAll(_ context.Context) []metric.Metric {
	res := make([]metric.Metric, 0)

	return res
}

func (db *DBRepository) Add(ctx context.Context, m metric.Metric) (metric.Metric, error) {
	query := `INSERT INTO metrics (type, name, value) VALUES (@type, @name, @value)`
	args := pgx.NamedArgs{
		"type":  m.Ident.Type.String(),
		"name":  m.Ident.Name,
		"value": m.Value,
	}
	_, err := db.pool.Exec(ctx, query, args)
	return m, err
}

func (db *DBRepository) Get(_ context.Context, _ metric.Ident) (metric.Metric, error) {
	m, err := metric.NewMetric("", 1.1, "")
	return *m, err
}

func (db *DBRepository) Increase(_ context.Context, m metric.Metric) (metric.Metric, error) {
	return m, nil
}

func NewDBRepository(db *db.DB) *DBRepository {
	return &DBRepository{
		pool: db.GetPool(),
	}
}

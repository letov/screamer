package repositories

import (
	"context"
	"math"
	"screamer/internal/common"
	"screamer/internal/common/metric"
	"screamer/internal/common/retry"
	"screamer/internal/server/db"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

type DBRepository struct {
	pool *pgxpool.Pool
	log  *zap.SugaredLogger
}

func (db *DBRepository) BatchUpdate(ctx context.Context, ms []metric.Metric) error {
	tx, err := db.pool.Begin(ctx)
	if err != nil {
		return err
	}
	for _, m := range ms {
		if m.Ident.Type == metric.Counter {
			_, err = db.Increase(ctx, m)
		} else {
			_, err = db.Add(ctx, m)
		}
		if err != nil {
			_ = tx.Rollback(ctx)
			return err
		}
	}

	return tx.Commit(ctx)
}

func (db *DBRepository) GetAll(ctx context.Context) (ms []metric.Metric) {
	ms = make([]metric.Metric, 0)

	is, err := db.getUniqIdents(ctx)
	if err != nil || len(is) == 0 {
		return
	}

	for _, i := range is {
		m, err := db.Get(ctx, i)
		if err != nil {
			return
		}
		ms = append(ms, m)
	}

	return
}

func (db *DBRepository) Add(ctx context.Context, m metric.Metric) (metric.Metric, error) {
	query := `INSERT INTO metrics (type, name, value) VALUES (@type, @name, @value)`
	args := pgx.NamedArgs{
		"type":  m.Ident.Type.String(),
		"name":  m.Ident.Name,
		"value": m.Value,
	}

	job := db.execJob(query, args)
	_, err := retry.NewRetryJob(ctx, "db exec", job, []error{}, []int{1, 2, 5}, db.log)
	return m, err
}

func (db *DBRepository) Get(ctx context.Context, i metric.Ident) (m metric.Metric, err error) {
	query := `SELECT type, name, value FROM metrics WHERE type=@type AND name=@name ORDER BY id DESC LIMIT 1`
	args := pgx.NamedArgs{
		"type": i.Type.String(),
		"name": i.Name,
	}

	job := db.queryJob(query, args)
	rows, err := retry.NewRetryJob(ctx, "db query", job, []error{}, []int{1, 2, 5}, db.log)
	if err != nil {
		return
	}
	defer rows.Close()

	var ident metric.Ident
	if rows.Next() {
		err = rows.Scan(&ident.Type, &ident.Name, &m.Value)
		if err != nil {
			db.log.Warn("DB repo error: ", err)
			return
		}
		m.Ident = ident
		return
	}

	err = common.ErrMetricNotExists
	return
}

func (db *DBRepository) Increase(ctx context.Context, m metric.Metric) (metric.Metric, error) {
	var _, frac float64
	_, frac = math.Modf(m.Value)
	if frac != 0 {
		return m, common.ErrInvalidValue
	}

	currentM, err := db.Get(ctx, m.Ident)
	if err != nil && err == common.ErrMetricNotExists {
		addM := *metric.NewCounter(m.Ident.Name, m.Value)
		return db.Add(ctx, addM)
	}
	if err != nil {
		return currentM, err
	}
	m.Value += currentM.Value
	return db.Add(ctx, m)
}

func (db *DBRepository) getUniqIdents(ctx context.Context) (is []metric.Ident, err error) {
	query := `SELECT DISTINCT type, name FROM metrics`
	args := pgx.NamedArgs{}

	job := db.queryJob(query, args)
	rows, err := retry.NewRetryJob(ctx, "db query", job, []error{}, []int{1, 2, 5}, db.log)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		var i metric.Ident
		err = rows.Scan(&i.Type, &i.Name)
		if err != nil {
			db.log.Warn("DB repo error: ", err)
			return
		}
		is = append(is, i)
	}

	return
}

func (db *DBRepository) queryJob(query string, args pgx.NamedArgs) func(ctx context.Context) (pgx.Rows, error) {
	return func(ctx context.Context) (pgx.Rows, error) {
		return db.pool.Query(ctx, query, args)
	}
}

func (db *DBRepository) execJob(query string, args pgx.NamedArgs) func(ctx context.Context) (pgconn.CommandTag, error) {
	return func(ctx context.Context) (pgconn.CommandTag, error) {
		return db.pool.Exec(ctx, query, args)
	}
}

func NewDBRepository(db *db.DB, log *zap.SugaredLogger) *DBRepository {
	return &DBRepository{
		pool: db.GetPool(),
		log:  log,
	}
}

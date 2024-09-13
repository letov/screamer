package repositories

import (
	"context"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
	"math"
	"screamer/internal/common"
	"screamer/internal/common/metric"
	"screamer/internal/server/db"
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
		query := `INSERT INTO metrics (type, name, value) VALUES (@type, @name, @value)`
		args := pgx.NamedArgs{
			"type":  m.Ident.Type.String(),
			"name":  m.Ident.Name,
			"value": m.Value,
		}
		_, err := db.pool.Exec(ctx, query, args)
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
	_, err := db.pool.Exec(ctx, query, args)
	if err != nil {
		db.log.Warn("DB repo error: ", err)
	}
	return m, err
}

func (db *DBRepository) Get(ctx context.Context, i metric.Ident) (m metric.Metric, err error) {
	query := `SELECT type, name, value FROM metrics WHERE type=@type AND name=@name ORDER BY id DESC LIMIT 1`
	args := pgx.NamedArgs{
		"type": i.Type.String(),
		"name": i.Name,
	}

	rows, err := db.pool.Query(ctx, query, args)
	if err != nil {
		db.log.Warn("DB repo error: ", err)
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

	rows, err := db.pool.Query(ctx, query)
	if err != nil {
		db.log.Warn("DB repo error: ", err)
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

func NewDBRepository(db *db.DB, log *zap.SugaredLogger) *DBRepository {
	return &DBRepository{
		pool: db.GetPool(),
		log:  log,
	}
}

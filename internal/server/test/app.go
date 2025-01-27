package test

import (
	"context"
	"fmt"
	"screamer/internal/server/infrastructure/db"
	"screamer/internal/server/infrastructure/di"
	"testing"

	"go.uber.org/fx"
	"go.uber.org/fx/fxtest"
)

func initTest(t *testing.T, r interface{}) {
	t.Setenv("IS_TEST_ENV", "true")
	app := fxtest.New(t, di.InjectApp(), fx.Invoke(r))
	defer app.RequireStop()
	app.RequireStart()
}

func flushDB(ctx context.Context, db *db.DB) error {
	pool := db.GetPool()
	query := `SELECT table_name "table" FROM information_schema.tables WHERE table_schema='public' AND table_type='BASE TABLE' AND table_name != 'goose_db_version';`
	rows, err := pool.Query(ctx, query)
	if err != nil {
		return err
	}

	var queries []string
	for rows.Next() {
		var table string
		err = rows.Scan(&table)
		if err != nil {
			return err
		}
		queries = append(queries, fmt.Sprintf("TRUNCATE %v CASCADE;", table))
	}

	tx, _ := pool.Begin(ctx)
	for _, query := range queries {
		_, err = tx.Exec(ctx, query)
		if err != nil {
			return err
		}
	}
	return tx.Commit(ctx)
}

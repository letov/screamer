package db

import (
	"context"
	"embed"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/pressly/goose/v3"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"screamer/internal/common"
	"screamer/internal/server/config"
)

type DB struct {
	pool *pgxpool.Pool
	log  *zap.SugaredLogger
}

//go:embed migrations/*.sql
var embedMigrations embed.FS

func (db *DB) GetPool() *pgxpool.Pool {
	return db.pool
}

func (db *DB) SetPool(pool *pgxpool.Pool) {
	db.pool = pool
}

func (db *DB) Ping(ctx context.Context) (err error) {
	pool := db.GetPool()
	if pool == nil {
		return common.ErrNoDBConnection
	}

	return pool.Ping(ctx)
}

func (db *DB) makeMigrations() {
	goose.SetBaseFS(embedMigrations)

	if err := goose.SetDialect("postgres"); err != nil {
		panic(err)
	}

	sqlDB := stdlib.OpenDBFromPool(db.GetPool())
	if err := goose.Up(sqlDB, "migrations"); err != nil {
		panic(err)
	}
	if err := goose.Version(sqlDB, "migrations"); err != nil {
		db.log.Fatal(err)
	}
}

func NewDB(lc fx.Lifecycle, log *zap.SugaredLogger, c *config.Config) *DB {
	db := &DB{
		log: log,
	}

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			log.Info("Init DB pool")

			poolConfig, err := pgxpool.ParseConfig(c.DBAddress)
			if err != nil {
				log.Warn("Failed to parce config: ", err)
				return nil
			}

			pool, err := pgxpool.NewWithConfig(ctx, poolConfig)
			if err != nil {
				log.Warn("Failed to create db pool: ", err)
				return nil
			}

			err = pool.Ping(ctx)
			if err != nil {
				log.Warn("Failed to ping db: ", err)
				return nil
			}

			db.SetPool(pool)
			db.makeMigrations()

			return nil
		},
		OnStop: func(ctx context.Context) error {
			if db.pool == nil {
				return nil
			}

			log.Info("Close DB pool")
			db.pool.Close()
			return nil
		},
	})

	return db
}

package db

import (
	"context"
	"github.com/jackc/pgx"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"screamer/internal/server/config"
)

type DB struct {
	conn *pgx.Conn
}

func (c *DB) GetConn() *pgx.Conn {
	return c.conn
}

func (c *DB) SetConn(conn *pgx.Conn) {
	c.conn = conn
}

func NewDB(lc fx.Lifecycle, log *zap.SugaredLogger, c *config.Config) *DB {
	db := &DB{}

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			log.Info("Init DB connection")

			connConfig, err := pgx.ParseConnectionString(c.DBAddress)
			if err != nil {
				return err
			}

			conn, err := pgx.Connect(connConfig)
			if err != nil {
				log.Warn("Failed DB connection: ", err)
			}

			db.SetConn(conn)

			return nil
		},
		OnStop: func(ctx context.Context) error {
			if db.conn == nil {
				return nil
			}

			log.Info("Close DB connection")

			err := db.conn.Close()
			if err != nil {
				log.Warn("Failed close DB connection: ", err)
			}

			return nil
		},
	})

	return db
}

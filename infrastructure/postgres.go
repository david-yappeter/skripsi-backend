package infrastructure

import (
	"context"
	"database/sql"
	"log"
	"myapp/global"
	"myapp/internal/pgx/logger"
	"os"
	"time"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
)

type DBTX interface {
	GetContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error
	SelectContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
	NamedExecContext(ctx context.Context, query string, arg interface{}) (sql.Result, error)
}

func NewPostgreSqlDB(dbConfig global.PostgresConfig) *sqlx.DB {
	pgxConnConfig, err := pgx.ParseConfig("")
	if global.IsDebug() {
		pgxLoggerConfig := logger.Config{
			SlowThreshold: 200 * time.Millisecond,
			Colorful:      true,
			LogLevel:      pgx.LogLevelInfo,
		}
		pgxLogger := logger.New(log.New(os.Stdout, "\r\n", log.LstdFlags), pgxLoggerConfig)

		pgxConnConfig.Logger = pgxLogger
		pgxConnConfig.LogLevel = pgxLoggerConfig.LogLevel
	}

	if err != nil {
		panic(err)
	}

	pgConnConf := &pgxConnConfig.Config
	pgConnConf.Host = dbConfig.Host
	pgConnConf.Port = dbConfig.Port
	pgConnConf.Database = dbConfig.Database
	pgConnConf.User = dbConfig.Username
	pgConnConf.Password = dbConfig.Password
	pgConnConf.RuntimeParams["timezone"] = "UTC"

	pgxDB := stdlib.OpenDB(*pgxConnConfig)
	if err = pgxDB.Ping(); err != nil {
		pgxDB.Close()
		panic(err)
	}

	db := sqlx.NewDb(pgxDB, "pgx")

	return db
}

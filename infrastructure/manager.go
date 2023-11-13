package infrastructure

import (
	"fmt"
	"myapp/database/migration"
	"myapp/global"

	"github.com/golang-migrate/migrate/v4"
	migratePgx "github.com/golang-migrate/migrate/v4/database/pgx"
	"github.com/jmoiron/sqlx"
)

type InfrastructureManager interface {
	// database
	GetDB() *sqlx.DB
	MigrateDB(isRollingBack bool, steps int) error
	RefreshDB() error
	CloseDB() error

	// logger stack
	GetLoggerStack() LoggerStack
}

type infrastructureManager struct {
	sqlDB       *sqlx.DB
	loggerStack LoggerStack
}

func (i *infrastructureManager) createDB() error {
	dbConfig := global.GetPostgresConfig()
	dbConfig.Database = ""
	sqlDB := NewPostgreSqlDB(dbConfig)

	if _, err := sqlDB.Exec(fmt.Sprintf(`CREATE DATABASE "%s" WITH ENCODING='UTF8';`, global.GetPostgresConfig().Database)); err != nil {
		return err
	}

	if err := sqlDB.Close(); err != nil {
		return err
	}

	i.sqlDB = NewPostgreSqlDB(global.GetPostgresConfig())

	return nil
}

func (i infrastructureManager) MigrateDB(isRollingBack bool, steps int) error {
	dbDriver, err := migratePgx.WithInstance(i.GetDB().DB, &migratePgx.Config{})
	if err != nil {
		return err
	}

	migrator, err := migrate.NewWithInstance("", migration.SourceDriver(), "pgx", dbDriver)
	if err != nil {
		return err
	}

	if isRollingBack {
		_, _, err := migrator.Version()
		if err != nil {
			return err
		}

		if steps > 0 {
			err = migrator.Steps(-1 * int(steps))
		} else {
			err = migrator.Down()
		}

		if err != nil {
			return err
		}
	} else {
		var err error
		if steps > 0 {
			err = migrator.Steps(int(steps))
		} else {
			err = migrator.Up()
		}

		if err != nil && err != migrate.ErrNoChange {
			return err
		}
	}

	return nil
}

func (i infrastructureManager) dropDB() error {
	dbConfig := global.GetPostgresConfig()
	dbConfig.Database = ""
	sqlDB := NewPostgreSqlDB(dbConfig)

	if _, err := sqlDB.Exec(fmt.Sprintf(`DROP DATABASE IF EXISTS "%s" WITH (FORCE);`, global.GetPostgresConfig().Database)); err != nil {
		return err
	}

	if err := sqlDB.Close(); err != nil {
		return err
	}

	return nil
}

func (i *infrastructureManager) RefreshDB() error {
	if err := i.dropDB(); err != nil {
		return err
	}

	if err := i.createDB(); err != nil {
		return err
	}

	if err := i.MigrateDB(false, 0); err != nil {
		return err
	}

	return nil
}

func (i infrastructureManager) GetDB() *sqlx.DB {
	return i.sqlDB
}

func (i infrastructureManager) GetLoggerStack() LoggerStack {
	return i.loggerStack
}

func (i infrastructureManager) CloseDB() error {
	if sqlDB := i.GetDB(); sqlDB != nil {
		return sqlDB.Close()
	}

	return nil
}

func NewInfrastructureManager() InfrastructureManager {
	return &infrastructureManager{
		sqlDB:       NewPostgreSqlDB(global.GetPostgresConfig()),
		loggerStack: NewLoggerStack(),
	}
}

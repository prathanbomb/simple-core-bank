package db

import (
	"context"
	"fmt"
	"os"
	"time"

	pgxdecimal "github.com/jackc/pgx-shopspring-decimal"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/tracelog"
	log "github.com/oatsaysai/simple-core-bank/src/logger"
)

type DB interface {
	DBAccountInterface
	DBTransferInterface
	DBTransactionInterface

	Close() error
}

type PostgresqlDB struct {
	logger log.Logger
	DB     *pgxpool.Pool
	Config *Config
}

func New(config *Config, logger log.Logger) (pgdb *PostgresqlDB, err error) {
	pgdb = &PostgresqlDB{
		logger: logger.WithFields(log.Fields{
			"package": "db",
		}),
	}
	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable search_path=%s",
		config.DBHost,
		config.DBPort,
		config.DBUsername,
		config.DBPassword,
		config.DBName,
		config.DBSchemaName,
	)

	var connectConf, _ = pgxpool.ParseConfig(connStr)
	connectConf.MaxConns = config.MaxOpenConns
	connectConf.MaxConns = config.MaxOpenConns
	connectConf.HealthCheckPeriod = 15 * time.Second
	connectConf.ConnConfig.Tracer = &tracelog.TraceLog{
		Logger:   NewDatabaseLogger(&pgdb.logger),
		LogLevel: tracelog.LogLevelTrace,
	}

	// Set timezone to PGX runtime
	if s := os.Getenv("TZ"); s != "" {
		connectConf.ConnConfig.RuntimeParams["timezone"] = s
	}

	// Register Decimal Data Type to PGX Pool
	connectConf.AfterConnect = func(ctx context.Context, conn *pgx.Conn) error {
		pgxdecimal.Register(conn.TypeMap())
		return nil
	}

	pgdb.DB, err = pgxpool.NewWithConfig(context.Background(), connectConf)
	if err != nil {
		pgdb.logger.Errorf("Error connecting to postgres: %+v")
		return nil, err
	}

	pgdb.Config = config

	return pgdb, nil
}

func (pgdb *PostgresqlDB) Close() error {
	pgdb.DB.Close()
	return nil
}

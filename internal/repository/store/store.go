package store

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"soa-product-management/internal/config"

	gormPostgres "gorm.io/driver/postgres"

	"go.uber.org/fx"
	"gorm.io/gorm"
)

type Storage struct {
	client *gorm.DB
	db     *sql.DB
	dbName string
	conf   *config.Config

	ProductStore *ProductStore
}

func InitStorage(
	conf *config.Config,
	lc fx.Lifecycle,
) *Storage {
	s := NewStorage(conf)
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			err := s.db.Ping()
			if err != nil {
				log.Fatalf("failed to connect postgresql: %v", err)
			}

			return nil
		},
		OnStop: func(ctx context.Context) error {
			return s.db.Close()
		},
	})
	return s
}

func NewStorage(
	conf *config.Config,
) *Storage {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		conf.Postgresql.Host, conf.Postgresql.Port, conf.Postgresql.Username,
		conf.Postgresql.Password, conf.Postgresql.DbName, conf.Postgresql.SSLMode)

	db, err := gorm.Open(gormPostgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to init postgresql: %v", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("failed to init postgresql: %v", err)
	}
	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(conf.Postgresql.MaxIdleConns)
	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(conf.Postgresql.MaxOpenConns)
	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(conf.Postgresql.ConnMaxLifetime)

	s := &Storage{
		client:       db,
		db:           sqlDB,
		dbName:       conf.Postgresql.DbName,
		conf:         conf,
		ProductStore: NewProductStore(db),
	}
	return s
}

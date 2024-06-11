package configs

import (
	"context"
	"fmt"
	"time"

	"github.com/bagasunix/go-clean-architecture/pkg/config"
	"github.com/bagasunix/go-clean-architecture/pkg/db"
	"github.com/bagasunix/go-clean-architecture/pkg/errors"
	"github.com/golang-migrate/migrate/v4"
	migPostgres "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq" // PostgreSQL driver
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/plugin/dbresolver"
)

func InitDb(ctx context.Context, logger *zap.Logger, configs *config.Cfg) *gorm.DB {
	configBuilder := &db.DbPostgresConfig{
		Driver:          configs.Database.Driver,
		Host:            configs.Database.Host,
		Port:            configs.Database.Port,
		User:            configs.Database.User,
		Password:        configs.Database.Password,
		DatabaseName:    configs.Database.DbName,
		SSLMode:         configs.Database.SSLMode,
		MaxIdleConns:    configs.Database.MaxIdleConn,
		MaxOpenConns:    configs.Database.MaxOpenConn,
		ConnMaxLifetime: configs.Database.ConnMaxLifetime,
		Timezone:        configs.General.Timezone,
	}

	return NewDB(ctx, logger, configBuilder)
}
func NewDB(ctx context.Context, logger *zap.Logger, dbConfig *db.DbPostgresConfig) *gorm.DB {
	// Membuat koneksi ke database dengan DSN dari dbConfig
	db, err := gorm.Open(postgres.Open(dbConfig.GetDSN()), &gorm.Config{SkipDefaultTransaction: true, PrepareStmt: true})
	errors.HandlerWithOSExit(logger, err, "init", "database", "config", dbConfig.GetDSN())
	// Mengatur parameter koneksi
	errors.HandlerWithOSExit(logger, db.WithContext(ctx).Use(dbresolver.Register(dbresolver.Config{}).SetMaxIdleConns(dbConfig.MaxIdleConns).SetMaxOpenConns(dbConfig.MaxOpenConns).SetConnMaxLifetime(time.Hour)), "db_resolver")

	sqlDB, _ := db.DB()
	driver, err := migPostgres.WithInstance(sqlDB, &migPostgres.Config{})
	errors.HandlerWithOSExit(logger, err, "failed to initialize postgres driver")

	migrationsPath := "../migrations"
	// Buat instance migrasi
	m, err := migrate.NewWithDatabaseInstance("file://"+migrationsPath, "postgres", driver)
	errors.HandlerWithOSExit(logger, err, "failed to create migration instance")

	// Jalankan migrasi ke versi terbaru
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		errors.HandlerWithOSExit(logger, err, "Failed to run migrations.")
	}

	fmt.Println("Migration successful")

	// Verifikasi koneksi
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	if err := sqlDB.PingContext(ctx); err != nil {
		errors.HandlerWithOSExit(logger, err, "init", "database", "ping", "")
	}

	return db
}

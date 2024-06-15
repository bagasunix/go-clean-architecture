package db

import (
	"fmt"
	"time"
)

type DbPostgresConfig struct {
	Driver          string
	Host            string
	Port            string
	User            string
	Password        string
	DatabaseName    string
	SSLMode         string
	Timezone        string
	MaxIdleConns    int
	MaxOpenConns    int
	ConnMaxLifetime time.Duration
}

func (d *DbPostgresConfig) GetDSN() string {
	return fmt.Sprintf("%s://%s:%s@%s:%s/%s?sslmode=%s", d.Driver, d.User, d.Password, d.Host, d.Port, d.DatabaseName, d.SSLMode)
}

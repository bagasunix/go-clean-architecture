package config

import (
	"context"
	"embed"
	"os"
	"time"

	"go.uber.org/zap"
	"gopkg.in/yaml.v2"
)

type Cfg struct {
	Server   Server   `yaml:"server"`
	Database Database `yaml:"database"`
	General  General  `yaml:"general"`
}

type Server struct {
	Port     string `yaml:"port"`
	Timezone string `yaml:"time_zone"`
}

type Database struct {
	Driver          string        `yaml:"driver"`
	Host            string        `yaml:"host"`
	DbName          string        `yaml:"db_name"`
	User            string        `yaml:"user"`
	Port            string        `yaml:"port"`
	Password        string        `yaml:"password"`
	SSLMode         string        `yaml:"ssl_mode"`
	MaxOpenConn     int           `yaml:"max_connection"`
	MaxIdleConn     int           `yaml:"max_idle"`
	ConnMaxLifetime time.Duration `yaml:"max_life"`
}

type General struct {
	CurrentLanguage string `yaml:"current_language"`
	AppName         string `yaml:"app_name"`
	AppVersion      string `yaml:"app_version"`
	Prefork         bool   `ymai:"prefork"`
	Env             string
}

//go:embed *
var files embed.FS

func (c *Cfg) SetDefault() *Cfg {
	if os.Getenv("DATABASE_HOST") != "" {
		c.Database.Host = os.Getenv("DATABASE_HOST")
	}
	if os.Getenv("DATABASE_NAME") != "" {
		c.Database.DbName = os.Getenv("DATABASE_NAME")
	}
	if os.Getenv("DATABASE_USER") != "" {
		c.Database.User = os.Getenv("DATABASE_USER")
	}
	if os.Getenv("DATABASE_PASSWORD") != "" {
		c.Database.Password = os.Getenv("DATABASE_PASSWORD")
	}
	if os.Getenv("PORT") != "" {
		c.Database.Password = os.Getenv("PORT")
	}
	return c
}

func InitConfig(ctx context.Context, log *zap.Logger) *Cfg {
	var data []byte
	bytes, err := files.ReadFile("config.yaml")
	if err != nil {
		log.Fatal("error when load config: ", zap.Error(err))
	}
	data = bytes
	config := new(Cfg)
	err = yaml.Unmarshal(data, config)
	if err != nil {
		log.Fatal("error when unmarshal config: ", zap.Error(err))
	}

	return config.SetDefault()
}

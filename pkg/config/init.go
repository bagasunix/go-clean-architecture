package config

import (
	"context"
	"embed"
	"os"
	"time"

	"go.uber.org/zap"
	"gopkg.in/yaml.v2"
)

var Config *Cfg

type Cfg struct {
	Server   Server   `yaml:"server"`
	Database Database `yaml:"database"`
	General  General  `yaml:"general"`
}

type Server struct {
	Port     int    `yaml:"port"`
	Timezone string `yaml:"time_zone"`
	SSLMode  string `yaml:"sll_mode"`
}

type Database struct {
	Driver          string        `yaml:"driver"`
	Host            string        `yaml:"host"`
	DbName          string        `yaml:"db_name"`
	User            string        `yaml:"user"`
	Port            string        `yaml:"port"`
	Password        string        `yaml:"password"`
	MaxConn         int           `yaml:"max_connection"`
	MaxIdle         int           `yaml:"max_idle"`
	ConnMaxLifetime time.Duration `yaml:"max_life"`
}

type General struct {
	CurrentLanguage string `yaml:"current_language"`
	AppName         string `yaml:"app_name"`
	AppVersion      string `yaml:"app_version"`
	Env             string
}

//go:embed *
var files embed.FS

func (c *Cfg) SetDefault() {
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
}

func InitConfig(ctx context.Context, log *zap.Logger) {
	var data []byte

	bytes, err := files.ReadFile("config.yaml")
	if err != nil {
		log.Fatal("error when load config: ", zap.Error(err))
	}
	data = bytes
	err = yaml.Unmarshal(data, &Config)
	if err != nil {
		log.Fatal("error when unmarshal config: ", zap.Error(err))
	}
	Config.SetDefault()
}

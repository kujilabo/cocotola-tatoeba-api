package config

import (
	"io/ioutil"
	"os"
	"strings"

	"github.com/gin-contrib/cors"
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"

	lib "github.com/kujilabo/cocotola-tatoeba-api/pkg_lib/domain"
)

type AppConfig struct {
	Port int `yaml:"port" validate:"required"`
}

type SQLite3Config struct {
	File string `yaml:"file" validate:"required"`
}

type MySQLConfig struct {
	Username string `yaml:"username" validate:"required"`
	Password string `yaml:"password" validate:"required"`
	Host     string `yaml:"host" validate:"required"`
	Port     int    `yaml:"port" validate:"required"`
	Database string `yaml:"database" validate:"required"`
}

type DBConfig struct {
	DriverName string         `yaml:"driverName"`
	SQLite3    *SQLite3Config `yaml:"sqlite3"`
	MySQL      *MySQLConfig   `yaml:"mysql"`
}

type AuthConfig struct {
	Username string `yaml:"username" validate:"required"`
	Password string `yaml:"password" validate:"required"`
}

type CORSConfig struct {
	AllowOrigins []string `yaml:"allowOrigins"`
}

type ShutdownConfig struct {
	TimeSec1 int `yaml:"timeSec1" validate:"gte=1"`
	TimeSec2 int `yaml:"timeSec2" validate:"gte=1"`
}

type LogConfig struct {
	Level string `yaml:"level"`
}

type DebugConfig struct {
	GinMode bool `yaml:"ginMode"`
	Wait    bool `yaml:"wait"`
}

type SwaggerConfig struct {
	Host   string `yaml:"host"`
	Schema string `yaml:"schema"`
}

type Config struct {
	App      *AppConfig      `yaml:"app" validate:"required"`
	DB       *DBConfig       `yaml:"db" validate:"required"`
	Auth     *AuthConfig     `yaml:"auth" validate:"required"`
	CORS     *CORSConfig     `yaml:"cors" validate:"required"`
	Shutdown *ShutdownConfig `yaml:"shutdown" validate:"required"`
	Log      *LogConfig      `yaml:"log" validate:"required"`
	Debug    *DebugConfig    `yaml:"debug"`
	Swagger  *SwaggerConfig  `yaml:"swagger" validate:"required"`
}

func LoadConfig(env string) (*Config, error) {
	confContent, err := ioutil.ReadFile("./configs/" + env + ".yml")
	if err != nil {
		return nil, err
	}

	confContent = []byte(os.ExpandEnv(string(confContent)))
	conf := &Config{}
	if err := yaml.Unmarshal(confContent, conf); err != nil {
		return nil, err
	}

	if err := lib.Validator.Struct(conf); err != nil {
		return nil, err
	}

	return conf, nil
}

func InitLog(env string, cfg *LogConfig) error {
	formatter := &logrus.JSONFormatter{
		FieldMap: logrus.FieldMap{
			logrus.FieldKeyLevel: "severity",
			logrus.FieldKeyMsg:   "message",
		},
	}
	logrus.SetFormatter(formatter)

	switch strings.ToLower(cfg.Level) {
	case "trace":
		logrus.SetLevel(logrus.TraceLevel)
	case "debug":
		logrus.SetLevel(logrus.DebugLevel)
	case "info":
		logrus.SetLevel(logrus.InfoLevel)
	case "warn":
		logrus.SetLevel(logrus.WarnLevel)
	case "error":
		logrus.SetLevel(logrus.ErrorLevel)
	case "fatal":
		logrus.SetLevel(logrus.FatalLevel)
	default:
		logrus.Infof("Unsupported log level: %s", cfg.Level)
		logrus.SetLevel(logrus.WarnLevel)
	}

	logrus.SetOutput(os.Stdout)

	return nil
}

func InitCORS(cfg *CORSConfig) cors.Config {
	if len(cfg.AllowOrigins) == 1 && cfg.AllowOrigins[0] == "*" {
		return cors.Config{
			AllowAllOrigins: true,
			AllowMethods:    []string{"*"},
			AllowHeaders:    []string{"*"},
		}
	}

	return cors.Config{
		AllowOrigins: cfg.AllowOrigins,
		AllowMethods: []string{"*"},
		AllowHeaders: []string{"*"},
	}
}

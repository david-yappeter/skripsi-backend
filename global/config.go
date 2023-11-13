package global

import (
	"fmt"
	"os"
	"path"
	"runtime"
	"strings"
	"time"

	"gopkg.in/yaml.v2"
)

const (
	ApiVersion = "1.0.0"

	EnvironmentDevelopment = "development"
	EnvironmentProduction  = "production"
	EnvironmentTesting     = "testing"

	StaticStorageFs = "public"
)

type PostgresConfig struct {
	Host     string `yaml:"host"`
	Port     uint16 `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Database string `yaml:"database"`
}

type JwtConfig struct {
	PrivateKeyFilePath string `yaml:"private_key_file_path"`
	PublicKeyFilePath  string `yaml:"public_key_file_path"`
}

type TelegramBotConfig struct {
	Token  string `yaml:"token"`
	ChatId int64  `yaml:"chat_id"`
}

type YamlConfig struct {
	timeLocation       *time.Location
	BaseDir            string
	StorageDir         string
	AppName            string            `yaml:"app_name"`
	Environment        string            `yaml:"environment"`
	IsDebug            bool              `yaml:"debug"`
	LogChannel         []string          `yaml:"log_channel"`
	Timezone           string            `yaml:"timezone"`
	Port               uint              `yaml:"port"`
	Uri                string            `yaml:"uri"`
	Filesystem         string            `yaml:"filesystem"`
	PdfGeneratorUrl    string            `yaml:"pdf_generator_url"`
	CorsAllowedOrigins []string          `yaml:"cors_allowed_origins"`
	Postgres           PostgresConfig    `yaml:"postgres"`
	JwtConfig          JwtConfig         `yaml:"jwt"`
	TelegramBotConfig  TelegramBotConfig `yaml:"telegram"`
}

var config YamlConfig

func init() {
	if err := LoadConfig(); err != nil {
		panic(err)
	}
}

func LoadConfig() error {
	baseDir, err := os.Getwd()
	if err != nil {
		return err
	}

	if _, err := os.Stat(fmt.Sprintf("%s/conf.yml", baseDir)); err != nil {
		_, filename, _, _ := runtime.Caller(0)
		baseDir = path.Join(path.Dir(filename), "../")
	}

	config.BaseDir = strings.TrimRight(strings.ReplaceAll(baseDir, "\\\\", "/"), "/")
	config.StorageDir = fmt.Sprintf("%s/storage", config.BaseDir)

	yamlFilePath := fmt.Sprintf("%s/conf.yml", config.BaseDir)
	if _, err := os.Stat(yamlFilePath); err != nil {
		return err
	}

	yamlFile, err := os.ReadFile(yamlFilePath)
	if err != nil {
		return err
	}

	err = yaml.UnmarshalStrict(yamlFile, &config)
	if err != nil {
		return err
	}

	config.timeLocation, err = time.LoadLocation(config.Timezone)
	if err != nil {
		return err
	}

	return nil
}

func GetConfig() YamlConfig {
	return config
}

func GetTimeLocation() *time.Location {
	return config.timeLocation
}

func GetBaseDir() string {
	return config.BaseDir
}

func GetStorageDir() string {
	return config.StorageDir
}

func GetJwtPrivateKeyFilePath() string {
	return config.JwtConfig.PrivateKeyFilePath
}

func GetJwtPublicKeyFilePath() string {
	return config.JwtConfig.PublicKeyFilePath
}

func GetAppName() string {
	return config.AppName
}

func GetLogChannel() []string {
	return config.LogChannel
}

func GetTimezone() string {
	return config.Timezone
}

func GetFilesystem() string {
	return config.Filesystem
}

func GetPostgresConfig() PostgresConfig {
	return config.Postgres
}

func GetTelegramConfig() TelegramBotConfig {
	return config.TelegramBotConfig
}

func SetEnvironment(environment string) {
	switch environment {
	case EnvironmentDevelopment, EnvironmentProduction, EnvironmentTesting:
		config.Environment = environment
	}
}

func EnableDebug() {
	config.IsDebug = true
}

func DisableDebug() {
	config.IsDebug = false
}

func IsProduction() bool {
	return config.Environment == EnvironmentProduction
}

func IsDevelopment() bool {
	return config.Environment == EnvironmentDevelopment
}

func IsTesting() bool {
	return config.Environment == EnvironmentTesting
}

func IsDebug() bool {
	return config.IsDebug
}

func GetPdfGeneratorUrl() string {
	return config.PdfGeneratorUrl
}

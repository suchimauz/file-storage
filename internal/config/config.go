package config

import (
	"os"
	"time"

	"github.com/spf13/viper"
)

const (
	defaultBasicAuth = "please:let"

	defaultHTTPPort               = "8000"
	defaultHTTPRWTimeout          = 10 * time.Second
	defaultHTTPMaxHeaderMegabytes = 1

	defaultStoragePrefix = "storage"

	EnvLocal = "local"
	EnvProd  = "prod"
)

type (
	Config struct {
		Environment string
		HTTP        HTTPConfig
		Storage     FileStorageConfig
		Auth        AuthConfig
	}

	HTTPConfig struct {
		Host               string        `mapstructure:"host"`
		Port               string        `mapstructure:"port"`
		ReadTimeout        time.Duration `mapstructure:"readTimeout"`
		WriteTimeout       time.Duration `mapstructure:"writeTimeout"`
		MaxHeaderMegabytes int           `mapstructure:"maxHeaderBytes"`
	}

	FileStorageConfig struct {
		Driver    string `mapstructure:"driver"`
		Bucket    string `mapstructure:"bucket"`
		AccessKey string `mapstructure:"access_key"`
		SecretKey string `mapstructure:"secret_key"`
		Prefix    string `mapstructure:"prefix"`
	}

	AuthConfig struct {
		Basic string `mapstructure:"basic"`
	}
)

func Init(configPath string) (*Config, error) {
	populateDefaultValues()

	if err := parseYamlConfig(configPath, os.Getenv("APP_ENV")); err != nil {
		return nil, err
	}

	var config Config
	if err := unmarshalYamlConfig(&config); err != nil {
		return nil, err
	}

	setFromEnv(&config)

	return &config, nil
}

func parseYamlConfig(folder, env string) error {
	viper.AddConfigPath(folder)
	viper.SetConfigName("main")

	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	if env == EnvLocal {
		return nil
	}

	viper.SetConfigName(env)

	return viper.MergeInConfig()
}

func unmarshalYamlConfig(config *Config) error {
	if err := viper.UnmarshalKey("auth", &config.Auth); err != nil {
		return err
	}

	if err := viper.UnmarshalKey("storage", &config.Storage); err != nil {
		return err
	}

	if err := viper.UnmarshalKey("storage.driver", &config.Storage.Driver); err != nil {
		return err
	}

	return viper.UnmarshalKey("http", &config.HTTP)
}

func setFromEnv(config *Config) {
	config.Auth.Basic = os.Getenv("AUTH_BASIC")

	config.Storage.Driver = os.Getenv("STORAGE_DRIVER")
	config.Storage.Bucket = os.Getenv("STORAGE_BUCKET")
	config.Storage.SecretKey = os.Getenv("STORAGE_PREFIX")
	config.Storage.AccessKey = os.Getenv("STORAGE_ACCESS_KEY")
	config.Storage.SecretKey = os.Getenv("STORAGE_SECRET_KEY")

	config.HTTP.Host = os.Getenv("HTTP_HOST")
	config.HTTP.Port = os.Getenv("HTTP_PORT")

	config.Environment = os.Getenv("APP_ENV")
}

func populateDefaultValues() {
	viper.SetDefault("auth.basic", defaultBasicAuth)

	viper.SetDefault("http.port", defaultHTTPPort)
	viper.SetDefault("http.max_header_megabytes", defaultHTTPMaxHeaderMegabytes)
	viper.SetDefault("http.timeouts.read", defaultHTTPRWTimeout)
	viper.SetDefault("http.timeouts.write", defaultHTTPRWTimeout)
}

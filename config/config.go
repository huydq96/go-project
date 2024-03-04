package config

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

// Config struct for webapp config
type (
	Config struct {
		Server struct {
			Host    string `mapstructure:"host"`
			Port    string `mapstructure:"port"`
			Timeout struct {
				// Server is the general server timeout to use
				// for graceful shutdowns
				Server time.Duration `mapstructure:"server"`

				// Write is the amount of time to wait until an HTTP server
				// write operation is cancelled
				Write time.Duration `mapstructure:"write"`

				// Read is the amount of time to wait until an HTTP server
				// read operation is cancelled
				Read time.Duration `mapstructure:"read"`

				// IDLE is the amount of time to wait
				// until an IDLE HTTP session is closed
				Idle time.Duration `mapstructure:"idle"`
			} `mapstructure:"timeout"`
		} `mapstructure:"server"`
		Http    Http          `mapstructure:"http"`
		Adapter AdapterConfig `mapstructure:"adapter"`
	}

	Http struct {
		BasePath        string `mapstructure:"base_path"`
		DebugHeaders    string `mapstructure:"debug_headers"`
		HttpClientDebug bool   `mapstructure:"http_client_debug"`
		DebugErrorsResp bool   `mapstructure:"debug_errors_resp"`
	}

	AdapterConfig struct {
		Mongodb struct {
			Alert MongoConfig `mapstructure:"alert"`
		} `mapstructure:"mongodb"`
	}

	MongoConfig struct {
		URI        string `mapstructure:"uri"`
		Database   string `mapstructure:"database"`
		Collection string `mapstructure:"collection"`
	}
)

func InitConfig() (*Config, error) {
	cfgPath, err := ParseFlags()
	if err != nil {
		return nil, err
	}

	config := &Config{}
	viper.SetConfigType("yaml")
	viper.SetConfigFile(cfgPath)

	if err := viper.ReadInConfig(); err != nil {
		return nil, errors.Wrap(err, "viper.ReadInConfig")
	}

	if err := viper.Unmarshal(config); err != nil {
		return nil, errors.Wrap(err, "viper.Unmarshal")
	}

	return config, nil
}

// ParseFlags will create and parse the CLI flags
// and return the path to be used elsewhere
func ParseFlags() (string, error) {
	var configPath string
	flag.StringVar(&configPath, "config", "/config/config.yml", "path to config file")
	flag.Parse()

	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}
	configPath = fmt.Sprintf("%s%s", dir, configPath)

	return configPath, nil
}

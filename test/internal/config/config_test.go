package config

import (
	"os"
	"score-publisher-svc/internal/config"
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func TestNewConfig_Success(t *testing.T) {
	// Mock the config values instead of using a real file
	viper.Set("name", "test-service")
	viper.Set("host", "localhost")
	viper.Set("port", 8080)

	viper.Set("dbconfig.host", "db-host")
	viper.Set("dbconfig.port", 5432)
	viper.Set("dbconfig.username", "db-user")
	viper.Set("dbconfig.password", "db-pass")
	viper.Set("dbconfig.dbname", "testdb")

	viper.Set("kafka.broker", "kafka-broker")
	viper.Set("kafka.topic", "test-topic")

	viper.Set("logger.level", "info")
	viper.Set("logger.format", "json")

	// Load the configuration
	cfg, err := config.NewConfig()

	assert.NoError(t, err)
	assert.NotNil(t, cfg)

	assert.Equal(t, "test-service", cfg.Name)
	assert.Equal(t, "localhost", cfg.Host)
	assert.Equal(t, 8080, cfg.Port)

	assert.Equal(t, "db-host", cfg.DbConfig.Host)
	assert.Equal(t, 5432, cfg.DbConfig.Port)
	assert.Equal(t, "db-user", cfg.DbConfig.Username)
	assert.Equal(t, "db-pass", cfg.DbConfig.Password)
	assert.Equal(t, "testdb", cfg.DbConfig.DbName)

	assert.Equal(t, "kafka-broker", cfg.Kafka.Broker)
	assert.Equal(t, "test-topic", cfg.Kafka.Topic)

	assert.Equal(t, "info", cfg.Logger.Level)
	assert.Equal(t, "json", cfg.Logger.Format)
}

// TestNewConfig_Failure tests if an invalid file path causes failure.
func TestNewConfig_Failure(t *testing.T) {
	configPath := "./invalid/path/to/config.yaml"
	os.Setenv("CONFIG_PATH", configPath)

	_, err := config.NewConfig()

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "fatal error config file")
}

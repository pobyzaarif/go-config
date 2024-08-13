package config_test

import (
	"os"
	"testing"

	config "github.com/pobyzaarif/go-config"
	"github.com/stretchr/testify/assert"
)

type appConfig struct {
	Host string `env:"HOST" envDefault:"127.0.0.1"`
	Port string `env:"PORT" envDefault:"8081"`
}

type dbConfig struct {
	Host string `env:"HOST" envDefault:"192.168.1.1"`
	Port string `env:"PORT" envDefault:"3306"`
	User string `env:"USER" envDefault:"root"`
	Pass string `env:"PASS" envDefault:"toor"`
	Name string `env:"NAME" envDefault:"mydb"`
}

type myConfig struct {
	App appConfig `envPrefix:"APP_"`
	DB  dbConfig  `envPrefix:"DB_"`
}

func TestWithDefault(t *testing.T) {
	var cfg myConfig
	err := config.LoadConfig(&cfg)
	assert.NoError(t, err)
	assert.Equal(t, "127.0.0.1", cfg.App.Host)
	assert.Equal(t, "8081", cfg.App.Port)
	assert.Equal(t, "192.168.1.1", cfg.DB.Host)
	assert.Equal(t, "3306", cfg.DB.Port)
	assert.Equal(t, "root", cfg.DB.User)
	assert.Equal(t, "toor", cfg.DB.Pass)
	assert.Equal(t, "mydb", cfg.DB.Name)
}

func TestWithEnvFile(t *testing.T) {
	envContent := `APP_HOST=127.0.0.2
APP_PORT=8082
DB_HOST=192.168.1.2
DB_PORT=5432
DB_USER=admin
DB_PASS=secret
DB_NAME=mydb2`

	err := os.WriteFile(".env", []byte(envContent), 0644)
	assert.NoError(t, err)
	defer os.Remove(".env")

	var cfg myConfig
	err = config.LoadConfig(&cfg)
	assert.NoError(t, err)
	assert.Equal(t, "127.0.0.2", cfg.App.Host)
	assert.Equal(t, "8082", cfg.App.Port)
	assert.Equal(t, "192.168.1.2", cfg.DB.Host)
	assert.Equal(t, "5432", cfg.DB.Port)
	assert.Equal(t, "admin", cfg.DB.User)
	assert.Equal(t, "secret", cfg.DB.Pass)
	assert.Equal(t, "mydb2", cfg.DB.Name)
}

func TestWithLocalEnv(t *testing.T) {
	os.Setenv("APP_HOST", "127.0.0.3")
	os.Setenv("APP_PORT", "8083")
	os.Setenv("DB_HOST", "192.168.1.3")
	os.Setenv("DB_PORT", "5433")
	os.Setenv("DB_USER", "user")
	os.Setenv("DB_PASS", "pass")
	os.Setenv("DB_NAME", "mydb3")

	defer func() {
		os.Unsetenv("APP_HOST")
		os.Unsetenv("APP_PORT")
		os.Unsetenv("DB_HOST")
		os.Unsetenv("DB_PORT")
		os.Unsetenv("DB_USER")
		os.Unsetenv("DB_PASS")
		os.Unsetenv("DB_NAME")
	}()

	var cfg myConfig
	err := config.LoadConfig(&cfg)
	assert.NoError(t, err)
	assert.Equal(t, "127.0.0.3", cfg.App.Host)
	assert.Equal(t, "8083", cfg.App.Port)
	assert.Equal(t, "192.168.1.3", cfg.DB.Host)
	assert.Equal(t, "5433", cfg.DB.Port)
	assert.Equal(t, "user", cfg.DB.User)
	assert.Equal(t, "pass", cfg.DB.Pass)
	assert.Equal(t, "mydb3", cfg.DB.Name)
}

func TestInvalidConfig(t *testing.T) {
	var cfg myConfig
	err := config.LoadConfig(cfg)
	assert.Error(t, err)
}

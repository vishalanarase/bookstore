package config

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// GlobalConfig application global configuration
type GlobalConfig struct {
	Port             int
	DatabaseName     string
	DatabaseHostname string
	DatabaseUsername string
	DatabasePassword string
	DatabaseNameTest string
}

// Config reads config from env to config struct
func Config(configPath string) GlobalConfig {
	var config GlobalConfig
	v := viper.New()

	v.AddConfigPath(configPath)
	v.SetConfigName(".env")
	v.SetConfigType("env")

	v.AutomaticEnv()

	if _, err := os.Stat(configPath + ".env"); err == nil {
		err := v.ReadInConfig()
		if err != nil {
			panic(fmt.Errorf("fatal error reading config file: %w", err))
		}
	}

	config.Port = v.GetInt("port")
	config.DatabaseUsername = strings.TrimSuffix(v.GetString("database_username"), "\n")
	config.DatabasePassword = strings.TrimSuffix(v.GetString("database_password"), "\n")
	config.DatabaseHostname = strings.TrimSuffix(v.GetString("database_hostname"), "\n")
	config.DatabaseName = strings.TrimSuffix(v.GetString("database_name"), "\n")
	config.DatabaseNameTest = strings.TrimSuffix(v.GetString("database_name_test"), "\n")

	return config
}

// DatabaseConnection open connection to db
func DatabaseConnection(config GlobalConfig) (*gorm.DB, error) {
	// Open mysql connection
	var err error

	if os.Getenv("API_ENV") == "test" {
		config.DatabaseName = config.DatabaseNameTest
	}

	dns := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.DatabaseUsername, config.DatabasePassword, config.DatabaseHostname, config.DatabaseName)
	db, err := gorm.Open(mysql.Open(dns), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		return nil, err
	}
	// defer db.Close()

	return db, nil
}

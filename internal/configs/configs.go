package configs

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	log "github.com/sirupsen/logrus"
)

// GlobalConfig application global configuration
type GlobalConfig struct {
	Port             int
	DatabaseName     string
	DatabaseHostname string
	DatabaseUsername string
	DatabasePassword string
	DatabasePort     int
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
			log.Fatal("Failed to read config: %w", err)
		}
	}

	config.Port = v.GetInt("port")
	config.DatabaseUsername = strings.TrimSuffix(v.GetString("database_username"), "\n")
	config.DatabasePassword = strings.TrimSuffix(v.GetString("database_password"), "\n")
	config.DatabaseHostname = strings.TrimSuffix(v.GetString("database_hostname"), "\n")
	config.DatabaseName = strings.TrimSuffix(v.GetString("database_name"), "\n")
	config.DatabasePort = v.GetInt("database_port")
	config.DatabaseNameTest = strings.TrimSuffix(v.GetString("database_name_test"), "\n")

	return config
}

// DatabaseConnection open connection to db
func DatabaseConnection(config GlobalConfig) (*gorm.DB, error) {
	// Open mysql connection
	var err error

	log.Info("Connecting to database")

	if os.Getenv("API_ENV") == "test" {
		config.DatabaseUsername = os.Getenv("DATABASE_USERNAME")
		config.DatabasePassword = os.Getenv("DATABASE_PASSWOR")
		config.DatabaseName = os.Getenv("DATABASE_NAME")
		config.DatabasePort, _ = strconv.Atoi(os.Getenv("DATABASE_PORT"))
	}

	dns := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.DatabaseUsername, config.DatabasePassword, config.DatabaseHostname, config.DatabasePort, config.DatabaseName)
	db, err := gorm.Open(mysql.Open(dns), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		log.WithError(err).Error("Failed to open database connection")
		return nil, err
	}
	// defer db.Close()

	return db, nil
}

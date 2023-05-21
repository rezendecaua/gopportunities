package config

import (
	"os"

	"github.com/rezendecaua/gopportunities/schemas"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitializeSQLite() (*gorm.DB, error) {
	logger := GetLogger("sqlite")
	dbPath := "./db/main.db"

	_, err := os.Stat(dbPath)
	if os.IsNotExist(err) {
		logger.Infof("sqlite database not found, creating one")
		err = os.MkdirAll("./db", os.ModePerm)
		if err != nil {
			logger.Errorf("sqlite database creation folder error: %v", err)
			return nil, err
		}

		file, err := os.Create(dbPath)
		if err != nil {
			logger.Errorf("sqlite database creation file error: %v", err)
			return nil, err
		}

		file.Close()
	}

	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		logger.Errorf("sqlite opening error: %v", err)
		return nil, err
	}

	err = db.AutoMigrate(&schemas.Opening{})
	if err != nil {
		logger.Errorf("sqlite migration error: %v", err)
		return nil, err
	}

	logger.Infof("sqlite connection established")
	return db, nil
}

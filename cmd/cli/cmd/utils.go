package cmd

import (
	"errors"
	"os"

	"github.com/hculpan/vinylbase/pkg/db"
	"github.com/joho/godotenv"
)

func startup() error {
	err := godotenv.Load()
	if err != nil {
		return errors.New("error loading .env file")
	}

	return nil
}

func startupWithDb() error {
	if err := startup(); err != nil {
		return err
	}

	dbName := os.Getenv("DB_NAME")
	if dbName == "" {
		return errors.New("DB_NAME undefined")
	}

	tursoToken := os.Getenv("TURSO_TOKEN")
	if tursoToken == "" {
		return errors.New("TURSO_TOKEN undefined")
	}

	if err := db.InitDb(dbName, tursoToken); err != nil {
		return err
	}

	return nil
}

/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/hculpan/vinylbase/pkg/db"
	"github.com/hculpan/vinylbase/pkg/entities"
	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
)

// initdbCmd represents the initdb command
var initdbCmd = &cobra.Command{
	Use:   "initdb",
	Short: "Initializes the DB, erasing any contents",
	Long: `This command initializes the database, erasing
any existing contents. It will rebuild the schema, and 
populate it with any required data.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		// Load environment variables
		err := godotenv.Load()
		if err != nil {
			return errors.New("error loading .env file")
		}

		dbName := os.Getenv("DB_NAME")
		if dbName == "" {
			return errors.New("DB_NAME undefined")
		}

		tursoToken := os.Getenv("TURSO_TOKEN")
		if tursoToken == "" {
			return errors.New("TURSO_TOKEN undefined")
		}

		if strings.HasPrefix(dbName, "file://") {
			err := os.Remove(dbName[7:])
			if err != nil && !errors.Is(err, os.ErrNotExist) {
				return fmt.Errorf("unable to remove database file: %w", err)
			}
		}

		if err := db.InitDb(dbName, tursoToken); err != nil {
			return err
		}
		defer db.CloseDb()

		if err := entities.CreateUserTable(); err != nil {
			return err
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(initdbCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initdbCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// initdbCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

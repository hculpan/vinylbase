/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/hculpan/vinylbase/pkg/db"
	"github.com/spf13/cobra"
)

// listusersCmd represents the listusers command
var listusersCmd = &cobra.Command{
	Use:   "listusers",
	Short: "List the users",
	Long:  `List the users`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) > 0 {
			return errors.New("no arguments permitted for listusers")
		}

		if err := startupWithDb(); err != nil {
			return err
		}
		defer db.CloseDb()

		if err := getUsers(); err != nil {
			return err
		}

		return nil
	},
}

func getUsers() error {
	_, err := db.Query("SELECT rowid, username, realname, password FROM users", db.QueryFunc(func(index int, rows *sql.Rows) error {
		var id int
		var username string
		var realname string
		var password string
		err := rows.Scan(&id, &username, &realname, &password)
		if err != nil {
			return err
		}
		fmt.Printf("%3d: %-30s %-50q %-20s\n", id, username, realname, password)

		return nil
	}))
	if err != nil {
		return err
	}

	return nil
}

func init() {
	rootCmd.AddCommand(listusersCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listusersCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listusersCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

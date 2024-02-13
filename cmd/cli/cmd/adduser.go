/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"errors"
	"fmt"

	"github.com/hculpan/vinylbase/pkg/db"
	"github.com/hculpan/vinylbase/pkg/entities"
	"github.com/spf13/cobra"
)

// adduserCmd represents the adduser command
var adduserCmd = &cobra.Command{
	Use:   "adduser",
	Short: "Create a new user",
	Long:  `Create a new user, setting the username, realname, and password.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) != 3 {
			return errors.New("must specify username, real name, and password")
		}

		if err := startupWithDb(); err != nil {
			return err
		}
		defer db.CloseDb()

		user, err := entities.FetchUser(args[0])
		if err != nil {
			return err
		} else if user != nil {
			return fmt.Errorf("user %q already registered", args[0])
		}

		user = &entities.User{
			Username: args[0],
			Realname: args[1],
		}
		user.SetPassword(args[2])

		if err := user.SaveUser(); err != nil {
			return err
		}

		fmt.Printf("User %q added\n", user.Username)

		return nil
	},
}

func init() {
	rootCmd.AddCommand(adduserCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// adduserCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// adduserCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

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

// removeuserCmd represents the removeuser command
var removeuserCmd = &cobra.Command{
	Use:   "removeuser",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return errors.New("must specify username")
		}

		if err := startupWithDb(); err != nil {
			return err
		}
		defer db.CloseDb()

		user, err := entities.FetchUser(args[0])
		if err != nil {
			return err
		} else if user == nil {
			return fmt.Errorf("user %q not found", args[0])
		}

		if err := user.DeleteUser(); err != nil {
			return err
		}

		fmt.Printf("User %q removed\n", user.Username)

		return nil
	},
}

func init() {
	rootCmd.AddCommand(removeuserCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// removeuserCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// removeuserCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

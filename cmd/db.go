package cmd

import (
	"time"

	"github.com/spf13/cobra"
	"github.com/vsantos1/Goke/config"
	"github.com/vsantos1/Goke/database"
)

var testCmd = &cobra.Command{
	Use:   "test",
	Short: "Test your database connection from config.yaml",
	Run: func(cmd *cobra.Command, args []string) {
		time.Sleep(1000)
		c, err := config.ReadConfigFile("config.yaml")
		if err != nil {
			panic("Error while reading file config.yaml ")
		}

		// Test base configuration to connect on database if not exists
		println(database.TestDatabaseconnection(c))
		defer database.CloseDatabase(c)

	},
}

func init() {
	RootCmd.AddCommand(testCmd)
}
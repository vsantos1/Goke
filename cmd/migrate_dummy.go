package cmd

import (
	"log"
	"time"

	"github.com/spf13/cobra"
	"github.com/vsantos1/Goke/config"
	"github.com/vsantos1/Goke/handlers"
	"github.com/vsantos1/Goke/utils"
)

var schema string
var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Migrate your sql model from dummy to database",
	Run: func(cmd *cobra.Command, args []string) {
		time.Sleep(1000)

		c, err := config.ReadConfigFile("goke-config.yaml")
		if err != nil {
			panic("Error while reading file goke-config..yaml ")
		}

		typ := handlers.HandleWithMigrateModel(c, handlers.ReadDummyBytes(), schema)
		if typ.Problem != nil {
			log.Fatalf("Error: %v", typ.Problem)
		}
		
		handlers.CreateHistoryDb(schema, handlers.ReadDummyBytes())

		println(typ.Message)
	},
}

func init() {

	migrateCmd.Flags().StringVarP(&schema, "name", "n", utils.RandomMigrationName(), "Create history for migrations")
	RootCmd.AddCommand(migrateCmd)
}

package cmd

import (
	"log"

	"os"

	"github.com/drud/ddev/pkg/util"
	"github.com/spf13/cobra"
)

var dbSource string

// ImportDBCmd represents the `ddev import-db` command.
var ImportDBCmd = &cobra.Command{
	Use:   "import-db",
	Short: "Import the database of an existing site to the local dev environment.",
	Long:  "Import the database of an existing site to the local development environment. The database can be provided as a SQL dump in a .sql, .sql.gz, or .tar.gz format. For the .tar.gz format, a SQL dump in .sql format must be present at the root of the archive.",
	PreRun: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			err := cmd.Usage()
			util.CheckErr(err)
			os.Exit(0)
		}

		client, err := util.GetDockerClient()
		if err != nil {
			log.Fatal(err)
		}

		err = util.EnsureNetwork(client, netName)
		if err != nil {
			log.Fatal(err)
		}

	},
	Run: func(cmd *cobra.Command, args []string) {
		app, err := getActiveApp()
		if err != nil {
			log.Fatalf("Could not find an active ddev configuration, have you run 'ddev config'?: %v", err)
		}

		err = app.ImportDB(dbSource)
		if err != nil {
			util.Failed("Failed to import database for %s: %s", app.GetName(), err)
		}
		util.Success("Successfully imported database for %s", app.GetName())
	},
}

func init() {
	ImportDBCmd.Flags().StringVarP(&dbSource, "src", "", "", "Provide the path to a sql dump in .sql or .tar.gz format")
	RootCmd.AddCommand(ImportDBCmd)
}
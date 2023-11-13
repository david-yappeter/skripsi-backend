//go:build devtools

package cmd

import (
	"fmt"
	"io/ioutil"
	"myapp/util"

	"github.com/spf13/cobra"
)

const migrationFilePath string = "database/migration"

// Argument to format is the version number.
const migrationContent = `package migration

func init() {
	sourceDriver.append(
		%s,
		` + "`" + `
		` + "`" + `,
		` + "`" + `
		` + "`" + `,
	)
}
`

func init() {
	rootCmd.AddCommand(newMigrateGenCommand())
}

func newMigrateGenCommand() *cobra.Command {
	var filename string

	cmd := &cobra.Command{
		Use:   "migrate-gen",
		Short: "Generate migration file",
		Long:  "Generate migration file",
		Run: func(_ *cobra.Command, _ []string) {
			var (
				version           = util.CurrentDateTime().Format("200601021504")
				migrationFilePath = migrationFilePath + "/" + fmt.Sprintf("%s_%s.go", version, filename)
			)

			// Write migration file.
			if err := ioutil.WriteFile(migrationFilePath, []byte(fmt.Sprintf(migrationContent, version)), 0644); err != nil {
				panic(err)
			}
		},
	}

	cmd.Flags().StringVarP(&filename, "filename", "f", "", "Specify the file name without timestamp ex : create_user_table")
	cmd.MarkFlagRequired("filename")

	return cmd
}

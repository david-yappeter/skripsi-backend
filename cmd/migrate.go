//go:build tools

package cmd

import (
	"myapp/global"
	"myapp/manager"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(newMigrateCommand())
}

func newMigrateCommand() *cobra.Command {
	var (
		isRollingBack bool
		steps         int
	)

	cmd := &cobra.Command{
		Use:   "migrate",
		Short: "Migrate the database table",
		Long:  "Migrate the database table",
		Run: func(_ *cobra.Command, _ []string) {
			global.DisableDebug()

			container := manager.NewContainer(manager.LoadDefault)
			if err := container.MigrateDB(isRollingBack, steps); err != nil {
				panic(err)
			}
		},
	}

	cmd.Flags().BoolVarP(&isRollingBack, "rollback", "", false, "Indicate whether migration is rollback or not")
	cmd.Flags().IntVarP(&steps, "steps", "s", 0, "Specify steps if want to migrate n number of migrations")

	return cmd
}

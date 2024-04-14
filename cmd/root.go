package cmd

import (
	"fmt"
	"log"
	"myapp/delivery/api"
	"myapp/global"
	"myapp/manager"
	"os"

	"github.com/spf13/cobra"
)

var flagVersion bool

var rootCmd = &cobra.Command{
	Use:   "myapp",
	Short: "Start the http server",
	Run: func(_ *cobra.Command, _ []string) {
		if flagVersion {
			fmt.Println(global.ApiVersion)
			return
		}

		container := manager.NewContainer(manager.FullConfig)
		defer func() {
			if err := container.Close(); err != nil {
				panic(err)
			}
		}()

		router := api.NewRouter(container)

		addr := fmt.Sprintf(":%d", global.GetConfig().Port)

		if err := router.Run(addr); err != nil {
			panic(err)
		}
	},
}

func Execute() {
	rootCmd.Flags().BoolVarP(&flagVersion, "version", "v", false, "Version")

	if err := rootCmd.Execute(); err != nil {
		log.Println(err)
		fmt.Fprintf(os.Stderr, "Whoops. There was an error while executing your CLI '%s'", err)
		os.Exit(1)
	}

}

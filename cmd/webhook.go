//go:build tools

package cmd

import (
	"fmt"
	"myapp/delivery/api"
	"myapp/global"
	"myapp/manager"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(newWebhookCommand())
}

func newWebhookCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "webhook",
		Short: "Start the webhook http server",
		Run: func(_ *cobra.Command, _ []string) {
			container := manager.NewContainer(manager.DefaultConfig)
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

	return cmd
}

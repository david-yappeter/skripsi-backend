//go:build tools

package cmd

import (
	"context"
	"fmt"
	"log"
	"myapp/global"
	"myapp/manager"
	"myapp/model"
	"myapp/util"

	"github.com/go-co-op/gocron"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(newScheduleCommand())
}

func newScheduleCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "schedule",
		Short: "Run the scheduled commands",
		Run: func(_ *cobra.Command, _ []string) {
			global.DisableDebug()

			container := manager.NewContainer(manager.DefaultConfig)
			defer func() {
				if err := container.Close(); err != nil {
					panic(err)
				}
			}()

			useCaseManager := container.UseCaseManager()

			scheduler := gocron.NewScheduler(global.GetTimeLocation())

			// nextMinute := time.Now().Truncate(time.Minute).Add(1 * time.Minute)
			// sleepDuration := time.Until(nextMinute)
			// time.Sleep(sleepDuration)

			scheduler.SingletonModeAll()

			executeAction := func(action string, fn func(ctx context.Context)) {
				ctx := newAuditLogCtx(fmt.Sprintf("SCHEDULE %s", action))

				actionId := model.MustGetRequestIdCtx(ctx)

				logMessage := fmt.Sprintf("%s [%s]", action, actionId)

				log.Printf("STARTED %s\n", logMessage)
				defer log.Printf("FINISHED %s\n", logMessage)

				defer util.PanicHandler()
				fn(ctx)
			}

			scheduler.Every(1).Day().Do(
				func() {
					executeAction(
						"auto-update-tiktok-access-token",
						func(ctx context.Context) {
							useCaseManager.TiktokConfigUseCase().AutoUpdate(ctx)
						},
					)
				},
			)

			scheduler.StartBlocking()
		},
	}

	return cmd
}

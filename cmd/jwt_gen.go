//go:build devtools

package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"myapp/global"
	"myapp/manager"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(newJwtGen())
}

func newJwtGen() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "jwt-gen",
		Short: "Generate JWT (for development purpose)",
		Long:  "This command is used to generate JWT",
		Run: func(_ *cobra.Command, args []string) {
			if len(args) == 0 {
				panic("First Argument (username) is required")
			}
			username := args[0]

			global.DisableDebug()

			container := manager.NewContainer(manager.DefaultConfig)
			defer func() {
				if err := container.Close(); err != nil {
					panic(err)
				}
			}()

			ctx := context.Background()

			userRepository := container.RepositoryManager().UserRepository()
			authUseCase := container.UseCaseManager().AuthUseCase()

			user, err := userRepository.GetByUsernameAndIsActive(ctx, username, true)
			if err != nil {
				panic(err)
			}

			accessToken, err := authUseCase.GenerateJWT(ctx, user.Id)
			if err != nil {
				panic(err)
			}

			marshaled, _ := json.MarshalIndent(accessToken, "", " ")

			fmt.Println(string(marshaled))
		},
	}

	return cmd
}

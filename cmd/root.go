package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

const APP_NAME = "bjc"
const APP_USAGE = "A Count Keeper for BlackJack. Can be used to keep track of deck counts, and calculate the true count."

type RootCmd struct {
	Command *cobra.Command
	Execute func()
}

func NewRootCmd() *RootCmd {
	rootCmd := &cobra.Command{
		Use:   APP_NAME,
		Short: APP_USAGE,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(APP_USAGE)
		},
	}
	return &RootCmd{
		Command: rootCmd,
		Execute: func() {
			if err := rootCmd.Execute(); err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		},
	}
}

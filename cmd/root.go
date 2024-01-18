package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

const APP_NAME = "bjc"
const APP_USAGE = "A Count Keeper for BlackJack. Can be used to keep track of deck counts, and calculate the true count."

var rootCmd = &cobra.Command{
	Use:   APP_NAME,
	Short: APP_USAGE,
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

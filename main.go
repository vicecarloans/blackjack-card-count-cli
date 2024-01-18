package main

import (
	"fmt"
	"vicecarloans/blackjack-card-count-cli/cmd"
)

func main() {
	fmt.Println("BlackJack Counter")

	rootCmd := cmd.NewRootCmd()
	rootCmd.Command.AddCommand(cmd.NewVersionCmd().Command)
	rootCmd.Execute()
}

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

type VersionCmd struct {
	Command *cobra.Command
}

func NewVersionCmd() *VersionCmd {
	versionCmd := &cobra.Command{
		Use:   "version",
		Short: "Print the version number of BlackJack Counter",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("BlackJack Counter v0.1 -- HEAD")
		},
	}

	return &VersionCmd{
		Command: versionCmd,
	}
}

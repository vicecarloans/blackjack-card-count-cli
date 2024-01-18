/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"strconv"
	"vicecarloans/blackjack-card-count-cli/services"

	"github.com/go-playground/validator/v10"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

var validate = validator.New()

func promptDeckCount() int {
	prompt := promptui.Prompt{
		Label: "Estimate the number of decks in the shoe",
		Validate: func(input string) error {
			return validate.Var(input, "required,number")
		},
	}

	result, err := prompt.Run()

	if err != nil {
		panic(err)
	}

	i, err := strconv.Atoi(result)

	if err != nil {
		panic(err)
	}

	return i
}

func promptPenetrationRate() int {
	prompt := promptui.Prompt{
		Label: "Estimate the number of decks in the shoe",
		Validate: func(input string) error {
			return validate.Var(input, "required,number")
		},
	}

	result, err := prompt.Run()

	if err != nil {
		panic(err)
	}

	i, err := strconv.Atoi(result)

	if err != nil {
		panic(err)
	}

	return i
}

type ConfigureCmd struct {
	Command        *cobra.Command
	counterService *services.CounterService
}

func NewConfigureCmd(counterService *services.CounterService) *ConfigureCmd {
	configureCmd := &cobra.Command{
		Use:   "configure",
		Short: "Configure your blackjack game settings",

		Run: func(cmd *cobra.Command, args []string) {
			deckCount := promptDeckCount()

			counterService.Setup(deckCount, 0.75)
		},
	}

	return &ConfigureCmd{
		Command:        configureCmd,
		counterService: counterService,
	}
}

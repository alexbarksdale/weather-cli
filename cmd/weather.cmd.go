package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(getWeatherCommand)
}

var getWeatherCommand = &cobra.Command{
	Use:   "get",
	Short: "Get the weather of a given zipcode.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Executed get", args)

	},
}

package cmd

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"weather-cli/utils"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(getWeatherCommand)
}

var getWeatherCommand = &cobra.Command{
	Use:   "get",
	Short: "Get the weather of a given zipcode.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Print("What city would you like the weather for? ")

		// Create a scanner and read input
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		input := scanner.Text()

		// Get the weather report
		weatherRes, err := utils.GetWeather(input)
		if err != nil {
			log.Fatalf("Error getting the weather %v", err)
		}

		fmt.Printf("The temperature in %v is %v°F, but feels like %v°F. \n", input, weatherRes.Main.Temp, weatherRes.Main.FeelsLike)

		fmt.Print("How does the weather make you feel? ")
		scanner.Scan()
		input = scanner.Text()

		fmt.Printf("Logged: When the weather is %v°F it makes you feel %v.", weatherRes.Main.Temp, strings.ToLower(input))
	},
}

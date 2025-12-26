package cmd

import (
	"fmt"
	"math/rand"

	"github.com/spf13/cobra"
)

var inspireCmd = &cobra.Command{
	Use:   "inspire",
	Short: "Print An inspiring quote",
	Run: func(cmd *cobra.Command, args []string) {
		quotes := []string{
			"Stay hungry, stay foolish. — Steve Jobs",
			"Simplicity is the soul of efficiency. — Austin Freeman",
			"Make it work, make it right, make it fast. — Kent Beck",
			"The best way to predict the future is to invent it. — Alan Kay",
			"Per aspera ad astra.",
		}

		fmt.Println(quotes[rand.Intn(len(quotes))])
	},
}

func init() {
	rootCmd.AddCommand(inspireCmd)
}

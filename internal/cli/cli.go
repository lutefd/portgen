package cli

import (
	"fmt"

	"github.com/lutefd/portgen/internal/app"
	"github.com/lutefd/portgen/internal/ui"
	"github.com/spf13/cobra"
)

var (
	minPort         int
	maxPort         int
	copyToClipboard bool
	shortMode       bool
	testMode        bool
)

func Execute() error {
	var rootCmd = &cobra.Command{
		Use:   "portgen",
		Short: "Generate a random unused port",
		Long:  ui.GetLongDescription(),
		Run: func(cmd *cobra.Command, args []string) {
			port := app.GeneratePort(minPort, maxPort)

			if testMode {
				fmt.Println(port)
				return
			}

			if copyToClipboard {
				if err := app.CopyToClipboard(port); err != nil {
					fmt.Printf("Failed to copy to clipboard: %v\n", err)
				} else if !shortMode {
					fmt.Println("Port copied to clipboard!")
				}
			}

			if shortMode {
				fmt.Println(port)
			} else {
				app.RunInteractiveMode(minPort, maxPort, copyToClipboard)
			}
		},
	}

	rootCmd.Flags().IntVarP(&minPort, "min", "m", 10000, "Minimum port number (inclusive)")
	rootCmd.Flags().IntVarP(&maxPort, "max", "M", 65535, "Maximum port number (inclusive)")
	rootCmd.Flags().BoolVarP(&copyToClipboard, "copy", "c", false, "Copy the generated port to clipboard")
	rootCmd.Flags().BoolVarP(&shortMode, "short", "s", false, "Print only the generated port number")
	rootCmd.Flags().BoolVar(&testMode, "test", false, "Run in test mode")

	rootCmd.SetUsageTemplate(ui.GetUsageTemplate())

	return rootCmd.Execute()
}

/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package universe

import (
	"axis-cli/data/AxisUniverse"

	"github.com/spf13/cobra"
)

// showCmd represents the show command
var ShowCmd = &cobra.Command{
	Use:   "show",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		entry, err := AxisUniverse.Find(args[0])
		if err != nil {
			return err
		}
		PrintEntries([]AxisUniverse.AxisUniverse{entry})
		return err
	},
}

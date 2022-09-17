/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package universe

import (
	"axis-cli/data/AxisUniverse"
	"fmt"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var ListCmd = &cobra.Command{
	Use:   "list",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Args: cobra.ExactArgs(0),
	RunE: func(cmd *cobra.Command, args []string) error {
		entries, err := AxisUniverse.List()
		if err != nil {
			return err
		}
		PrintEntries(entries)
		return err
	},
}

func PrintEntries(entries []AxisUniverse.AxisUniverse) {
	fmt.Println("Universe\t\tDescription")
	fmt.Println("--------\t\t-----------")
	for _, universe := range entries {
		fmt.Println(universe.Name, "\t\t", universe.Description)
	}
}

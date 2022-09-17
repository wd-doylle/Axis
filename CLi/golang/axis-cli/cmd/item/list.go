/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package item

import (
	"axis-cli/data"
	"axis-cli/data/AxisItem"

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
		entries, err := AxisItem.ListToMap()
		if err != nil {
			return err
		}
		data.PrintMapEntries(entries, []string{"id", "title", "description", "universe_id", "priority", "created_at", "time_due", "status"})
		return err
	},
}

func init() {

}

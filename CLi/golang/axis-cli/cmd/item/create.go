/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package item

import (
	"axis-cli/data/AxisItem"
	"time"

	"github.com/spf13/cobra"
)

var (
	description  string
	priority     uint8
	durationDue  time.Duration
	universeName string
)

// createCmd represents the create command
var CreateCmd = &cobra.Command{
	Use:   "create",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		timeDue := time.Now().Add(durationDue)
		err := AxisItem.Create(args[0], description, priority, timeDue, universeName)
		return err
	},
}

func init() {
	CreateCmd.Flags().StringVarP(&description, "description", "d", "", "An optional description of the item")
	CreateCmd.Flags().Uint8VarP(&priority, "priority", "p", 0, "Priority of the item")
	CreateCmd.Flags().DurationVar(&durationDue, "due", 0, "Due time of the item")
	CreateCmd.Flags().StringVarP(&universeName, "universe", "u", "", "Universe that the item belongs to")
	CreateCmd.MarkFlagRequired("universe")
}

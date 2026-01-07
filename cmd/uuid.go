package cmd

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/spf13/cobra"
)

var uuidCmd = &cobra.Command{
	Use:   "uuid",
	Short: "Generate a random UUID",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("%s %s\n", labelStyle.Render("Uuid:"), valueStyle.Render(uuid.NewString()))
	},
}

func init() {
	rootCmd.AddCommand(uuidCmd)
}

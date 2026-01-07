package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"golang.org/x/crypto/bcrypt"
)

var bcryptCmd = &cobra.Command{
	Use:   "bcrypt <password>",
	Short: "Generate a bcrypt hash for a password",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		password := args[0]
		hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			fmt.Println(errStyle.Render("Error generating bcrypt hash:"), err)
			return
		}
		fmt.Printf("%s %s\n", labelStyle.Render("Bcrypt:"), valueStyle.Render(string(hash)))
	},
}

func init() {
	rootCmd.AddCommand(bcryptCmd)
}

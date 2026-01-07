package cmd

import (
	"fmt"
	"os"

	"github.com/charmbracelet/lipgloss"
	"github.com/spf13/cobra"
)

var (
	titleStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("#38bdf8")).Bold(true)
	usageStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("#a3e635"))
	commandStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#fbbf24")).Bold(true)
	descStyle    = lipgloss.NewStyle().Foreground(lipgloss.Color("#a5b4fc"))
)

var rootCmd = &cobra.Command{
	Use:   "grist",
	Short: "Grist stands for grain that is made into flour which is a CLI tool to generate and transform data",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(titleStyle.Render("Grist stands for \"grain that is made into flour\" which is a CLI tool to generate and transform data\n"))
		fmt.Println(usageStyle.Render("Usage:"))
		fmt.Println("  grist [command]")
		fmt.Println()
		fmt.Println(usageStyle.Render("Available Commands:"))
		fmt.Printf("  %s\t%s\n", commandStyle.Render("uuid  "), descStyle.Render("Generate a random UUID"))
		fmt.Printf("  %s\t%s\n", commandStyle.Render("time  "), descStyle.Render("Convert between time formats (rfc, pg, s, ms, us, ns, pb)"))
		fmt.Printf("  %s\t%s\n", commandStyle.Render("bcrypt"), descStyle.Render("Generate a bcrypt hash for a password"))
		fmt.Printf("  %s\t%s\n", commandStyle.Render("color "), descStyle.Render("Convert between color formats (hex, rgb, hsl)"))
		fmt.Println()
		fmt.Println(`Use "grist [command] --help" for more information about a command.`)
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		_, err := fmt.Fprintln(os.Stderr, err)
		if err != nil {
			os.Exit(1)
		}
	}
}

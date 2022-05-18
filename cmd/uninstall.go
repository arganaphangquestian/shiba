package cmd

import (
	"github.com/arganaphangquestian/shiba/pkg"
	"github.com/spf13/cobra"
)

var uninstallCmd = &cobra.Command{
	Use:   "uninstall",
	Short: "Uninstall git hook",
	Long:  "Uninstall git hook",
	Run: func(cmd *cobra.Command, args []string) {
		pkg.Uninstall()
	},
}

func init() {
	rootCmd.AddCommand(uninstallCmd)
}

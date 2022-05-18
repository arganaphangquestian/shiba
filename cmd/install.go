package cmd

import (
	"github.com/arganaphangquestian/shiba/pkg"
	"github.com/spf13/cobra"
)

var installCmd = &cobra.Command{
	Use:   "install",
	Short: "Install git hook",
	Long:  "Install git hook",
	Run: func(cmd *cobra.Command, args []string) {
		pkg.Install()
	},
}

func init() {
	rootCmd.AddCommand(installCmd)
}

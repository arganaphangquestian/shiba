package cmd

import (
	"fmt"
	"os"

	"github.com/arganaphangquestian/shiba/pkg"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var setCmd = &cobra.Command{
	Use:   "set",
	Short: "Set git hook",
	Long:  "Set git hook",
	Run: func(cmd *cobra.Command, args []string) {
		if file == "" || command == "" {
			color.Set(color.FgRed)
			defer color.Unset()
			fmt.Println("failed to run `set` command")
			os.Exit(1)
		}
		err := pkg.Set(file, command)
		if err != nil {
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(setCmd)
	setCmd.Flags().StringVarP(&file, "file", "f", "", "File path for git hook")
	setCmd.Flags().StringVarP(&command, "command", "c", "", "Command for git hook")
}

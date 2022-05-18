package cmd

import (
	"fmt"
	"os"

	"github.com/arganaphangquestian/shiba/pkg"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var file, command string

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add git hook",
	Long:  "Add git hook",
	Run: func(cmd *cobra.Command, args []string) {
		if file == "" || command == "" {
			color.Set(color.FgRed)
			defer color.Unset()
			fmt.Println("failed to run `add` command")
			os.Exit(1)
		}
		err := pkg.Add(file, command)
		if err != nil {
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
	addCmd.Flags().StringVarP(&file, "file", "f", "", "File path for git hook")
	addCmd.Flags().StringVarP(&command, "command", "c", "", "Command for git hook")
}

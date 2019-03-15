package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "commit-activity-log",
	Short: "cal: Commit Activity Log",
	Long:  "Git Activity Log and Statistics",
}

// Execute run program main loop
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"golang.org/x/crypto/ssh/terminal"
)

func init() {
	rootCmd.AddCommand(listCmd)
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List",
	Long:  "List your activity in repository",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Print("Password: ")
		pw, err := terminal.ReadPassword(int(os.Stdin.Fd()))
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println(string(pw))
	},
}

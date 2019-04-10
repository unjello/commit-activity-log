package command

import (
	"fmt"
	"os"

	"golang.org/x/crypto/ssh/terminal"
)

func List() {
	fmt.Print("Password: ")
	pw, err := terminal.ReadPassword(int(os.Stdin.Fd()))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(string(pw))

	//c := bitbucket.NewBasicAuth("sg0224609", string(pw))
}

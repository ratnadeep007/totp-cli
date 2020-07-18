package controllers

import "fmt"

// HandleHelpCommnad ..
func HandleHelpCommnad() {
	help := `TOPT CLI - Simple CLI to manage your TOTPs
For usage: https://github.com/ratnadeep007/totp-cli/tree/v1.1#readme
Version: 1.1.1
`
	fmt.Println(help)
}

// HandleVersionCommand ...
func HandleVersionCommand() {
	fmt.Println("TOTP CLI - version 1.1.1")
}

package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/user"

	"theskylab.in/totp-cli/controllers"
)

func main() {
	help := flag.Bool("help", false, "Help Message")
	version := flag.Bool("version", false, "Show version")

	flag.Parse()

	if *help {
		controllers.HandleHelpCommnad()
		return
	}

	if *version {
		controllers.HandleVersionCommand()
		return
	}

	usr, err := user.Current()
	if err != nil {
		log.Fatal(err.Error())
	}
	usrHomeDir := usr.HomeDir
	envText, err := ioutil.ReadFile(usrHomeDir + "/.totp")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(0)
	}
	if len(os.Args) > 1 {
		providerName := os.Args[1]
		if providerName == "list" {
			controllers.ListProviders(envText)
		}
		if providerName != "" {
			controllers.GetOneToken(string(envText), providerName)
			return
		}
	}
	controllers.ListAllTokens(envText)
}

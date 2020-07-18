package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/user"
	"strings"
	"time"

	"github.com/pquerna/otp/totp"
)

func main() {
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
		if providerName != "" {
			envs := strings.Split(string(envText), "\n")
			for _, env := range envs {
				envSplit := strings.Split(env, "=")
				secret := envSplit[1]
				provider := envSplit[0]
				if strings.ToLower(provider) == strings.ToLower(providerName) {
					token := getToken(string(secret))
					printToken(string(envSplit[0]), token)
				}
			}
			return
		}
	}
	envs := strings.Split(string(envText), "\n")
	for _, env := range envs {
		envSplit := strings.Split(env, "=")
		secret := envSplit[1]
		token := getToken(string(secret))
		printToken(string(envSplit[0]), token)
	}
}

func getToken(secret string) string {
	token, _ := totp.GenerateCode(secret, time.Now())
	return token
}

func printToken(provider string, token string) {
	fmt.Printf("%s: %s\n", provider, token)
}

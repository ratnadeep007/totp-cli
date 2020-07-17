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
	envs := strings.Split(string(envText), "\n")
	for _, env := range envs {
		envSplit := strings.Split(env, "=")
		secret := envSplit[1]
		token, _ := totp.GenerateCode(string(secret), time.Now())
		fmt.Printf("%s: %s\n", string(envSplit[0]), token)
	}
}

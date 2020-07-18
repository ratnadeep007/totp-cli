package controllers

import (
	"fmt"
	"strings"
)

// ListProviders ...
func ListProviders(envText []byte) {
	envs := strings.Split(string(envText), "\n")
	fmt.Println("List of providers added:")
	for _, env := range envs {
		envSplit := strings.Split(env, "=")
		provider := envSplit[0]
		fmt.Printf("%s\n", provider)
	}
}

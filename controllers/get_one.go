package controllers

import (
	"strings"

	"theskylab.in/totp-cli/utils"
)

// GetOneToken ...
func GetOneToken(envText string, providerName string) {
	envs := strings.Split(envText, "\n")
	for _, env := range envs {
		envSplit := strings.Split(env, "=")
		secret := envSplit[1]
		provider := envSplit[0]
		if strings.ToLower(provider) == strings.ToLower(providerName) {
			token := utils.GetToken(secret)
			utils.PrintToken(provider, token)
		}
	}
}

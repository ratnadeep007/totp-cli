package controllers

import (
	"strings"

	"theskylab.in/totp-cli/utils"
)

// ListAllTokens ...
func ListAllTokens(envText []byte) {
	envs := strings.Split(string(envText), "\n")
	for _, env := range envs {
		envSplit := strings.Split(env, "=")
		secret := envSplit[1]
		provider := envSplit[0]
		token := utils.GetToken(string(secret))
		utils.PrintToken(provider, token)
	}
}

package config

import (
	"fmt"

	"github.com/markbates/goth"
	"github.com/markbates/goth/providers/github"
)

func setScheme() string {
	if Getenv("ENV") == "production" {
		return "https"
	}
	return "http"
}

func SetupGoth() {
	callbackUrl := fmt.Sprintf("%s://%s:%s/auth/github/callback", setScheme(), Getenv("HOST"), Getenv("PORT"))
	goth.UseProviders(github.New(Getenv("GITHUB_KEY"), Getenv("GITHUB_SECRET"), callbackUrl))
}

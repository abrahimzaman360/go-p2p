package config

import (
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var googleOAuthConfig = &oauth2.Config{
	RedirectURL:  "http://localhost:3000/auth/google/callback",
	ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
	ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
	Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email"},
	Endpoint:     google.Endpoint,
}

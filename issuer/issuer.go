package main

import (
	"context"
	"encoding/base64"
	"flag"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/bradleyfalzon/ghinstallation"
	"github.com/google/go-github/v34/github"
)

// Create a token from app id, private key, an repository
func main() {
	appID, pem, repository := getParameters()

	tr := http.DefaultTransport
	itr, err := ghinstallation.NewAppsTransport(tr, appID, []byte(pem))
	if err != nil {
		fatal("Unable to create apps transport: %v", err)
	}

	client := github.NewClient(&http.Client{Transport: itr})
	ctx := context.Background()

	// Get installation from repository
	installation, _, err := client.Apps.FindRepositoryInstallation(ctx, repository[0], repository[1])
	if err != nil {
		fatal("Unable to get installation from repository: %v", err)
	}

	// Get token from installation
	token, _, err := client.Apps.CreateInstallationToken(ctx, installation.GetID(), &github.InstallationTokenOptions{})
	if err != nil {
		fatal("Unable to create app token: %v", err)
	}

	// Output token so it can be captured
	fmt.Println(token.GetToken())
}

// Throw an error and die
func fatal(format string, substitutions ...interface{}) {
	fmt.Println(fmt.Sprintf(format, substitutions...))
	os.Exit(1)
}

// Get parameters from command line
func getParameters() (int64, string, []string) {
	var appID int64
	var pem, repo string

	flag.Int64Var(&appID, "appid", -1, "GitHub application ID")
	flag.StringVar(&pem, "pem", "", "GitHub application private key")
	flag.StringVar(&repo, "repository", "", "GitHub repository in owner/repository format")
	flag.Parse()

	// Split repository into owner/repository
	repository := strings.SplitN(repo, "/", 2)

	// Ensure all parameters are passed
	if appID < 1 {
		fatal("Unable to generate token: %s (%d)", "Application ID parameter is missing or invalid", appID)
	}

	if strings.TrimSpace(pem) == "" {
		fatal("Unable to generate token: %s", "Private key parameter is missing")
	}

	if len(repository) != 2 {
		fatal("Unable to generate token: %s (%s)", "Repository parameter is missing or invalid", repo)
	}

	// If pem is encoded, decode
	if isBase64(pem) {
		decoded, err := base64.StdEncoding.DecodeString(pem)
		if err == nil {
			pem = string(decoded)
		}
	}

	return appID, pem, repository
}

// Determine if a string is base64 encoded or not
func isBase64(s string) bool {
	_, err := base64.StdEncoding.DecodeString(s)

	return err == nil
}

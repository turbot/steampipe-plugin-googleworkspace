/*
  Script to generate a token to authenticate an user using OAuth 2.0

  This script uses Client secret credential to protect the user's data by only granting tokens to authorized requestors

  NOTE:
    Authorization information is stored on the file system, so subsequent executions don't prompt for authorization.

  Steps to follow:
    1. Generate the client secret, refer https://developers.google.com/workspace/guides/create-credentials
    2. Goto your terminal and run the script, i.e. `go run generateClientToken.go`
    3. The first time you run the sample, it prompts you to authorize access:
      * Browse to the provided URL in your web browser.
      * If you're not already signed in to your Google account, you're prompted to sign in.
        If you're signed in to multiple Google accounts, you are asked to select one account to use for authorization.
      * Click the Accept button.
      * Copy the code you're given, paste it into the command-line prompt, and press Enter.
    4. After successful execution, a token file will be generated inside steampipe directory (i.e. `~/.steampipe/token.json`)
*/

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/user"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/calendar/v3"
	"google.golang.org/api/drive/v3"
	"google.golang.org/api/gmail/v1"
)

// Retrieve a token, saves the token
func generateToken(config *oauth2.Config) {
	// The file token.json stores the user's access and refresh tokens, and is
	// created automatically when the authorization flow completes for the first
	// time.

	// Get the home dir path
	myself, err := user.Current()
	if err != nil {
		log.Fatalf("Unable to get current path: %v", err)
	}
	homedir := myself.HomeDir
	tokFile := homedir + "/.steampipe/token.json"

	tok := getTokenFromWeb(config)
	saveToken(tokFile, tok)
}

// Request a token from the web, then returns the retrieved token.
func getTokenFromWeb(config *oauth2.Config) *oauth2.Token {
	authURL := config.AuthCodeURL("state-token", oauth2.ApprovalForce)
	fmt.Printf("Go to the following link in your browser then type the "+
		"authorization code: \n%v\n", authURL)

	var authCode string
	if _, err := fmt.Scan(&authCode); err != nil {
		log.Fatalf("Unable to read authorization code: %v", err)
	}

	tok, err := config.Exchange(context.TODO(), authCode)
	if err != nil {
		log.Fatalf("Unable to retrieve token from web: %v", err)
	}
	return tok
}

// Saves a token to a file path.
func saveToken(path string, token *oauth2.Token) {
	fmt.Printf("Saving credential file to: %s\n", path)
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		log.Fatalf("Unable to cache oauth token: %v", err)
	}
	defer f.Close()
	json.NewEncoder(f).Encode(token)
}

func main() {
	// Get the client secret path from user
	var clientSecretPath string
	fmt.Println("Enter the client secret file path: ")
	fmt.Scanln(&clientSecretPath) // Read the input

	// Return, if no client secret provided
	if clientSecretPath == "" {
		log.Fatalf("Client secret path must be configured")
	}

	// Read the client secret
	b, err := ioutil.ReadFile(clientSecretPath)
	if err != nil {
		log.Fatalf("Unable to read client secret file: %v", err)
	}

	// If modifying these scopes, delete your previously saved token.json.
	config, err := google.ConfigFromJSON(
		b,
		drive.DriveReadonlyScope,
		calendar.CalendarReadonlyScope,
		gmail.GmailReadonlyScope,
	)
	if err != nil {
		log.Fatalf("Unable to parse client secret file to config: %v", err)
	}

	// Generate token, and saves it
	generateToken(config)
}

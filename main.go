package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	clientid := os.Getenv("CLIENT_ID")
	clientSecret := os.Getenv("CLIENT_SEC")

	fmt.Println("clientid: " + clientid)
	fmt.Println("clientSecret: " + clientSecret)
	// Replace with your actual client ID and secret
	config := &oauth2.Config{
		ClientID:     clientid,
		ClientSecret: clientSecret,
		RedirectURL:  "http://localhost:8080/callback",
		Scopes:       []string{"email", "profile"},
		Endpoint:     google.Endpoint,
	}

	// Create an OAuth2 flow
	authURL := config.AuthCodeURL("state", oauth2.AccessTypeOffline)
	fmt.Println("Visit this URL to authorize:", authURL)

	// Handle the callback
	http.HandleFunc("/callback", func(w http.ResponseWriter, r *http.Request) {
		// Get the authorization code from the query parameters
		code := r.URL.Query().Get("code")

		// Exchange the authorization code for an access token
		_, err := config.Exchange(context.Background(), code)
		if err != nil {
			// Handle error
			fmt.Println("Error exchanging code for token:", err)
			return
		}

		// Use the access token to make API calls
		// ...
	})

	// Start the HTTP server
	http.ListenAndServe(":8080", nil)
}

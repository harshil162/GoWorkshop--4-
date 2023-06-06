package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/sheets/v4"
)

var items []Item

func CredentialsFromFile(ctx context.Context, credentialsPath string, scopes ...string) (*oauth2.Config, error) {
	// Read the service account credentials file
	credentialsFile, err := os.Open(credentialsPath)
	if err != nil {
		return nil, fmt.Errorf("failed to open credentials file: %v", err)
	}
	defer credentialsFile.Close()

	// Parse the JSON credentials file
	credentialsJSON, err := ioutil.ReadAll(credentialsFile)
	if err != nil {
		return nil, fmt.Errorf("failed to read credentials file: %v", err)
	}

	// Create a Config from the JSON credentials
	config, err := google.ConfigFromJSON(credentialsJSON, scopes...)
	if err != nil {
		return nil, fmt.Errorf("failed to parse credentials: %v", err)
	}

	return config, nil
}
func main() {
	fileServer := http.FileServer(http.Dir("www/"))
	http.HandleFunc("/", fileServer.ServeHTTP)

	http.HandleFunc("/Items", getItemsHandler)
	// Replace with the path to your service account credentials JSON file
	credentialsPath := "/path/to/service_account_credentials.json"

	// Read the service account credentials file
	credentials, err := ioutil.ReadFile(credentialsPath)
	if err != nil {
		log.Fatalf("Unable to read client secret file: %v", err)
	}

	// Create a Config from the JSON credentials
	config, err := google.JWTConfigFromJSON(credentials, sheets.SpreadsheetsScope)
	if err != nil {
		log.Fatalf("Unable to parse client secret file to config: %v", err)
	}

	// Create a client for the Google Sheets API
	client := config.Client(context.Background())

	// Replace with the ID of your Google Sheet
	spreadsheetID := "1Ffs21UxsHPnvwM4l-wHdPxphKcm6usWWZMOyoGx8WjA/edit#gid=0"

	// Replace with the name or range of the sheet you want to retrieve data from
	sheetName := "MusicDownloaderSheet"

	// Make the API call to retrieve data from the sheet
	resp, err := client.Get(fmt.Sprintf("https://docs.google.com/spreadsheets/d/1Ffs21UxsHPnvwM4l-wHdPxphKcm6usWWZMOyoGx8WjA/edit#gid=0", spreadsheetID, sheetName))
	if err != nil {
		log.Fatalf("Unable to retrieve data from sheet: %v", err)
	}
	defer resp.Body.Close()

	// Read the response body
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Unable to read response body: %v", err)
	}

	// Print the response data
	fmt.Println(string(data))

	// Process the response and print the data
	/*if len(resp.Values) == 0 {
		fmt.Println("No data found.")
	} else {
		fmt.Println("Data:")
		for _, row := range resp.Values {
			for _, cell := range row {
				fmt.Printf("%s ", cell)
			}
			fmt.Println()
		}
	}*/
	startUpdateTimer()

	log.Println(http.ListenAndServe(":"+strconv.Itoa(3000), nil))
}

func jsonResponse(w http.ResponseWriter, x interface{}) {
	bytes, err := json.Marshal(x)
	if err != nil {
		log.Println(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(bytes)
}

func getItemsHandler(w http.ResponseWriter, _ *http.Request) {
	jsonResponse(w, items)
}

func startUpdateTimer() {
	ticker := time.NewTicker(2 * time.Second)
	go func() {
		for {
			_, ok := <-ticker.C
			if !ok {
				// channel is now closed
				break
			}
			items = getItems()
		}
	}()
}

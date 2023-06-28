package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	//"strconv"
	//"time"

	//"github.com/google/uuid"
	//"github.com/gorilla/mux"
	//"golang.org/x/oauth2"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"

	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
)

var items []Item

// Struct for JSON response
type Response struct {
	//*mux.Router
	Success bool   `json:"success"`
	Message string `json:"message"`
	Items   []Item `json:"data"`
}

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

/*
	func NewServer() *Response {
		i := &Item{
			Router: mux.NewRouter(),
			songs:  []Item{},
		}
		return i
	}
*/
func searchSong(songName string) (bool, error) {
	//Create a context and read the credentials file
	ctx := context.Background()
	creds, err := google.FindDefaultCredentials(ctx, sheets.SpreadsheetsScope)
	if err != nil {
		return false, err
	}
	//Create a new Sheets client using the credentials
	client, err := sheets.NewService(ctx, option.WithCredentials(creds))
	if err != nil {
		return false, err
	}
	//Specify the spreadsheet ID and range
	spreadsheetID := "1Ffs21UxsHPnvwM4l-wHdPxphKcm6usWWZMOyoGx8WjA/edit#gid=0"
	readRange := "MusicDownloaderSheet!A1:E6"
	//Make the API call to read the data from the sheet
	resp, err := client.Spreadsheets.Values.Get(spreadsheetID, readRange).Do()
	if err != nil {
		return false, err
	}
	for _, row := range resp.Values {
		if len(row) > 0 && row[0] == songName {
			return true, nil
		}
	}
	return false, nil
}
func searchHandler(w http.ResponseWriter, r *http.Request) {
	// Get the search query from the request
	query := r.FormValue("search")
	// Perform the search
	found, err := searchSong(query)
	if err != nil {
		log.Printf("Search error: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	// Display the results on the website
	if found {
		fmt.Fprint(w, "Song Found!")
	} else {
		fmt.Fprint(w, "No Results")
	}
}
func main() {
	// Set up routes
	http.HandleFunc("/search", searchHandler)

	// Start the server
	log.Println("Server started on port 3000")
	//log.Fatal(http.ListenAndServe(":3000", nil))

	/*fileServer := http.FileServer(http.Dir("www/"))
	http.HandleFunc("/", fileServer.ServeHTTP)
	http.HandleFunc("/Items", getItemsHandler)
	log.Println(http.ListenAndServe(":3000", nil))
	// Replace with the path to your service account credentials JSON file (to-do)
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
	//resp, err := client.Get(fmt.Sprintf("https://docs.google.com/spreadsheets/d/1Ffs21UxsHPnvwM4l-wHdPxphKcm6usWWZMOyoGx8WjA/edit#gid=0", spreadsheetID, sheetName))
	resp, err := client.Get(fmt.Sprintf("https://sheets.googleapis.com/v4/spreadsheets/%s/values/%s", spreadsheetID, sheetName))
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
	} else if len(resp.Values) > 0 {
		fmt.Println("Data:")
		for _, row := range resp.Values {
			for _, cell := range row {
				fmt.Printf("%s ", cell)
			}
			fmt.Println()
		}
	}*/
	//startUpdateTimer()

	//log.Println(http.ListenAndServe(":"+strconv.Itoa(localhost:3000), nil))
	//Start the HTTP server on localhost:8080
	//log.Println("Server is running on port 8080")
	log.Println(http.ListenAndServe(":3000", nil))
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
	jsonResponse(w, &items)
}

/*func startUpdateTimer() {
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
}*/

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
	"time"

	//"github.com/google/uuid"
	"github.com/gorilla/mux"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
)

var items []Item

// Struct for JSON response
type Response struct {
	*mux.Router
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

/*func NewServer() *Response {
	i := &Item{
		Router: mux.NewRouter(),
		songs:  []Item{},
	}
	return i
}*/

func main() {

	fileServer := http.FileServer(http.Dir("www/"))
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
	startUpdateTimer()

	//log.Println(http.ListenAndServe(":"+strconv.Itoa(localhost:3000), nil))
	//Start the HTTP server on localhost:8080
	//log.Println("Server is running on port 8080")
	//log.Println(http.ListenAndServe(":8080", nil))
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
func getSong(client *http.Client, spreadsheetID, sheetName string) ([][]interface{}, error) {
	// Create a new Sheets service client
	srv, err := sheets.NewService(context.Background(), option.WithHTTPClient(client))
	if err != nil {
		return nil, fmt.Errorf("unable to create Sheets service client: %v", err)
	}

	// Define the range to read from (assuming all columns)
	readRange := fmt.Sprintf("%s!A:Z", sheetName)

	// Make the request to retrieve the data
	resp, err := srv.Spreadsheets.Values.Get(spreadsheetID, readRange).Do()
	if err != nil {
		return nil, fmt.Errorf("unable to retrieve data from sheet: %v", err)
	}

	return resp.Values, nil
}

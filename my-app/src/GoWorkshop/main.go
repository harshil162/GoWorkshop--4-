package main

import (
	//"context"
	//"encoding/json"
	//"encoding/csv"
	"fmt"
	"log"

	"github.com/go-gota/gota/dataframe"
	"github.com/go-gota/gota/series"

	//"net/http"
	"os"
	//"strconv"
	//"time"
	//"github.com/google/uuid"
	//"github.com/gorilla/mux"
	//"golang.org/x/oauth2"
	//"golang.org/x/oauth2"
	//"golang.org/x/oauth2/google"
	//"google.golang.org/api/option"
	//"google.golang.org/api/sheets/v4"
)

//var items []Item

type Item struct {
	Name      string `json:"name"`
	Artist    string `json:"price"`
	Year      uint   `json:"year"`
	Genre     string `json:"genre"`
	Available bool   `json:"available"`
}

/*func readCSVFile(filePath string) [][]string {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Unable to read input file "+filePath, err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse file as CSV for "+filePath, err)
	}

	return records
}*/

func main() {
	//records := readCSVFile("./src/pages/MusicSheet.csv")
	file, err := os.OpenFile(".././MusicSheet.csv", os.O_RDONLY, os.ModePerm)
	if err != nil {
		log.Panic(err)
	}
	df := dataframe.ReadCSV(file)
	filter := df.Filter(dataframe.F{Colname: "price", Comparator: series.Greater, Comparando: 5000})
	fmt.Println(filter)
}

// Struct for JSON response
/*type Response struct {
	//*mux.Router
	Success bool   `json:"success"`
	Message string `json:"message"`
	Items   []Item `json:"data"`
}*/

/*func CredentialsFromFile(ctx context.Context, credentialsPath string, scopes ...string) (*oauth2.Config, error) {
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
}*/

/*
	func NewServer() *Response {
		i := &Item{
			Router: mux.NewRouter(),
			songs:  []Item{},
		}
		return i
	}
*/
/*func searchSong(songName string) (bool, error) {
	// Open Google Sheet using ID
	var ss = SpreadsheetApp.openById("1dBBPUN-4cXYnNYgshk3m64iotqWo7A3rp1o5rmRR9_CtGMffis5pD0Vh");
	var sheet = ss.getSheetByName("MusicSheet");
	// Read all data rows from Google Sheet
	const values = sheet.getRange(2, 1, sheet.getLastRow() - 1, sheet.getLastColumn()).getValues();
	// Converts data rows in json format
	const result = JSON.stringify(values.map(([a,b]) => ({ProductId: a,ProductName:b,})));
	// Returns Result
	return ContentService.createTextOutput(result).setMimeType(ContentService.MimeType.JSON);
}*/
/*func searchHandler(w http.ResponseWriter, r *http.Request) {
	// Get the search query from the request
	query := r.FormValue("search")
	// Perform the search
	found, err := doGet(query)
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

func readCSVFile(filePath string) [][]string {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Unable to read input file "+filePath, err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse file as CSV for "+filePath, err)
	}

	return records
}
func main() {
	records := readCSVFile("../MusicSheet.csv")
	fmt.Println(records)
	log.Println(http.ListenAndServe(":3000", nil))
}

/*func jsonResponse(w http.ResponseWriter, x interface{}) {
	bytes, err := json.Marshal(x)
	if err != nil {
		log.Println(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(bytes)
}*/

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

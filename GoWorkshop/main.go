package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"
)

var items []Item

func main() {
	fileServer := http.FileServer(http.Dir("www/"))
	http.HandleFunc("/", fileServer.ServeHTTP)

	http.HandleFunc("/Items", getItemsHandler)

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

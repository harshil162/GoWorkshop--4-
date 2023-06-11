package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"
)

type Item struct {
	Name      string `json:"name"`
	Artist    string `json:"price"`
	Year      int    `json:"year"`
	Genre     string `json:"genre"`
	Available bool   `json:"available"`
}

type Sheet struct {
	Range          string `json:"range"`
	MajorDimension string `json:"majorDimension"`
	Items          []Item `json:"values"`
}

// unmarshal one entry of json: values into an instance of an item
func (item *Item) UnmarshalJSON(p []byte) error {
	var tmp []interface{}
	if err := json.Unmarshal(p, &tmp); err != nil {
		return err
	}
	item.Name = tmp[0].(string)
	year, err := strconv.Atoi(tmp[1].(string))
	if err != nil {
		return err
	}
	item.Year = year
	if tmp[2].(string) == "Yes" {
		item.Available = true
	} else {
		item.Available = false
	}

	return nil
}

func Get(target interface{}) error {
	req, err := http.NewRequest("GET", "https://sheets.googleapis.com/v4/spreadsheets/1ZnwlhY5_gbbnuyCZjDS-w_Fw99MOhI_-uZTa7v9SQqk/values/Sheet1?valueRenderOption=FORMATTED_VALUE&key="+getKey(), bytes.NewBufferString(""))
	if err != nil {
		return err
	}

	client := http.Client{Timeout: time.Duration(10 * time.Second)}

	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()
	return json.NewDecoder(resp.Body).Decode(target)
}

func getItems() []Item {
	var sheet Sheet
	err := Get(&sheet)
	if err != nil {
		log.Println(err)
	}

	return sheet.Items
}

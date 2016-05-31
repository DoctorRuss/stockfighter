package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type VenueStatus struct {
	OK    bool `json:"ok"`
	Venue string `json:"venue"`
}

type APIStatus struct {
	OK    bool `json:"ok"`
	Error string `json:"error"`
}


func main() {
	// first call to Stockfighter API
	endpoint := "https://api.stockfighter.io/ob/api/heartbeat"
	bodytext := stockfighterAPI(endpoint)
	fmt.Printf("Stockfighter heartbeat: %s\n", bodytext)

	var apiStates APIStatus
	err := json.Unmarshal(bodytext, &apiStates)
	if err != nil {
		panic(err)
	}
	fmt.Printf("API heartbeat: %+v\n", apiStates)
	
	venueHeartbeat := "https://api.stockfighter.io/ob/api/venues/TESTEX/heartbeat"
	bodytext = stockfighterAPI(venueHeartbeat)
	fmt.Printf("TESTEX heartbeat: %s\n", bodytext)

	var venueStates VenueStatus
	err = json.Unmarshal(bodytext, &venueStates)
	if err != nil {
		panic(err)
	}

	fmt.Printf("TESTEX heartbeat: %+v\n", venueStates)
}

func stockfighterAPI(endpoint string) []byte {
	res, err := http.Get(endpoint)

	if err != nil {
		panic(err)
	}
	bodytext, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		panic(err)
	}
	return bodytext
}

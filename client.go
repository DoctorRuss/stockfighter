package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type VenueStatus struct {
	ok    bool
	venue string
}

func main() {
	endpoint := "https://api.stockfighter.io/ob/api/heartbeat"
	bodytext := stockfighterAPI(endpoint)
	fmt.Printf("Stockfighter heartbeat: %s\n", bodytext)

	venueHeartbeat := "https://api.stockfighter.io/ob/api/venues/TESTEX/heartbeat"
	bodytext = stockfighterAPI(venueHeartbeat)
	fmt.Printf("TESTEX heartbeat: %s\n", bodytext)

	var venueStates VenueStatus
	err := json.Unmarshal(bodytext, &venueStates)
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

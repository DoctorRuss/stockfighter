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

type StockDetails struct {
	Name    string `json:"name"`
	Symbol string `json:"symbol"`
}

type VenueStocks struct {
	OK    bool `json:"ok"`
	Stocks []StockDetails `json:"symbols"`
}

type Order struct {
	Price    int `json:"price"`
	Quantity int `json:"qty"`
	IsBuy    bool `json:"isBuy"`
}

type OrderBook struct {
	OK    bool `json:"ok"`
	Venue string `json:"venue"`
	Symbol string `json:"symbol"`
	Timestamp string `json:"ts"`

	Bids []Order `json:"bids"`
	Asks []Order `json:"asks"`
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
	
	venueStock := "https://api.stockfighter.io/ob/api/venues/TESTEX/stocks"
	bodytext = stockfighterAPI(venueStock)
	fmt.Printf("TESTEX stocks: %s\n", bodytext)

	var venueStocks VenueStocks
	err = json.Unmarshal(bodytext, &venueStocks)
	if err != nil {
		panic(err)
	}

	fmt.Printf("TESTEX stocks: %+v\n", venueStocks)
	
	stock := venueStocks.Stocks[0].Symbol
	stockOrder := "https://api.stockfighter.io/ob/api/venues/TESTEX/stocks/" + stock
	bodytext = stockfighterAPI(stockOrder)
	fmt.Printf("TESTEX order book: %s\n", bodytext)

	var orderbook OrderBook
	err = json.Unmarshal(bodytext, &orderbook)
	if err != nil {
		panic(err)
	}

	fmt.Printf("TESTEX order book: %+v\n", orderbook)
	
	
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

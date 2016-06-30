package main

import (
	"fmt"
	"time"
	"github.com/doctorruss/stockfighter/stockfighter"
)

func main() {
	// first call to Stockfighter API
	stockfighter.CheckStockfighterHeartbeat()
	
	// create level by name
	levelname := "first_steps"
	
	g := stockfighter.GM{LevelName:levelname}
	g.GetLevels()
	fmt.Printf("GM %+v\n", g)
	levelDetails := g.CreateLevel()
	fmt.Printf("GM %+v\n", g)
	
	if levelDetails.OK == false {
		panic(levelDetails.Error)
	}
	defer g.StopLevel()
	// create order broker with level details
	venue := levelDetails.Venues[0]
	tradingAccount := levelDetails.Account
	stock := levelDetails.Tickers[0]
	ob := stockfighter.OrderBroker{venue, tradingAccount, stock}
	
	// check venue heartbeat, because why not
	ob.CheckVenueHeartbeat()
	
	// get stock from venue
	// venueStocks := ob.GetStock()
    // fmt.Printf("%s stocks: %+v\n", venue, venueStocks)
	// // assign stock variable
	// stock := venueStocks.Stocks[0].Symbol
	
	// get order book
	for i := 0; i < 5; i++ {
	
		orderbook := ob.GetOrderBook()
		fmt.Printf("%s order book: %+v\n", venue, orderbook)
		
		quote := ob.GetQuote()
		
		fmt.Printf("%s %d quote: %+v\n", stock, quote.Spread(), quote)
	    time.Sleep(2000 * time.Millisecond)
	}
	quote := ob.GetQuote()
		
	ob.Buy(100, quote.Ask)
	
	fmt.Printf("instance details: %+v\n",	g.GetInstanceDetails())
}







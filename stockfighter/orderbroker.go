package stockfighter

import (
	"encoding/json"
	"fmt"
	
    "bytes"
)

func CheckStockfighterHeartbeat() {
	// first call to Stockfighter API
	var buffer bytes.Buffer
    buffer.WriteString("https://api.stockfighter.io")
    buffer.WriteString("/ob/api/heartbeat")
    
	
	bodytext := StockfighterAPI(buffer.String())
	//debug fmt.Printf("Stockfighter heartbeat: %s\n", bodytext)

	var apiStates APIStatus
	err := json.Unmarshal(bodytext, &apiStates)
	if err != nil {
		panic(err)
	}
	fmt.Printf("API heartbeat: %+v\n", apiStates)
}

type OrderBroker struct {
	Venue string
	TradingAccount string
	Stock string
}

func (ob *OrderBroker) CheckVenueHeartbeat() {
	venueHeartbeat := "https://api.stockfighter.io/ob/api/venues/" + ob.Venue + "/heartbeat"
	bodytext := StockfighterAPI(venueHeartbeat)
	// fmt.Printf("%s heartbeat: %s\n",ob.Venue, bodytext)
	var venueStates VenueStatus
	err := json.Unmarshal(bodytext, &venueStates)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s heartbeat:  %+v\n", ob.Venue, venueStates)
}

func (ob *OrderBroker) Buy(quantity int, price int) {
    var buffer bytes.Buffer
    buffer.WriteString("https://api.stockfighter.io")
    buffer.WriteString("/ob/api/venues/")
    buffer.WriteString( ob.Venue )
    buffer.WriteString("/stocks/")
    buffer.WriteString( ob.Stock)
    buffer.WriteString("/orders")
	fmt.Printf("endpoint: %s\n", buffer.String())
    
	placeOrder := PlaceOrderType{}
	placeOrder.Account = ob.TradingAccount
	placeOrder.Venue = ob.Venue
	placeOrder.Stock = ob.Stock
	placeOrder.Direction = "buy"
	placeOrder.Price = price
	placeOrder.Quantity = quantity
	placeOrder.OrderType = "limit"
	jsonStr, err := json.Marshal(placeOrder)
	if err != nil {
		panic(err)
	}
	
	body := StockfighterPost(buffer.String(), bytes.NewBuffer(jsonStr))
	fmt.Printf("order: %s\n", body)
}

func (ob *OrderBroker) GetOrderBook() OrderBook {
	var buffer bytes.Buffer
    buffer.WriteString("https://api.stockfighter.io")
    buffer.WriteString("/ob/api/venues/")
    buffer.WriteString( ob.Venue )
    buffer.WriteString("/stocks/")
    buffer.WriteString( ob.Stock)
	
	
	bodytext := StockfighterAPI(buffer.String())
	fmt.Printf("%s order book: %s\n", ob.Venue, bodytext)

	var orderbook OrderBook
	err := json.Unmarshal(bodytext, &orderbook)
	if err != nil {
		panic(err)
	}

    return orderbook
}



func (ob *OrderBroker) GetQuote() QuoteType {
	
	var buffer bytes.Buffer
    buffer.WriteString("https://api.stockfighter.io")
    buffer.WriteString("/ob/api/venues/")
    buffer.WriteString( ob.Venue )
    buffer.WriteString("/stocks/")
    buffer.WriteString( ob.Stock)
    buffer.WriteString("/quote")
	
	bodytext := StockfighterAPI(buffer.String())
	// debug fmt.Printf("%s quote: %s\n", ob.Venue, bodytext)

	var quote QuoteType
	err := json.Unmarshal(bodytext, &quote)
	if err != nil {
		panic(err)
	}

    return quote
}

// get stock from venue

func (ob *OrderBroker) GetStock() VenueStocks {
	var buffer bytes.Buffer
    buffer.WriteString("https://api.stockfighter.io")
    buffer.WriteString("/ob/api/venues/")
    buffer.WriteString( ob.Venue )
    buffer.WriteString("/stocks")
	bodytext := StockfighterAPI(buffer.String())
	//debug fmt.Printf("%s stocks: %s\n", venue, bodytext)
	var venueStocks VenueStocks
	err := json.Unmarshal(bodytext, &venueStocks)
	if err != nil {
		panic(err)
	}
	return venueStocks
}
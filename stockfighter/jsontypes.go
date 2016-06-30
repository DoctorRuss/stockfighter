package stockfighter 

type QuoteType struct {
	Symbol    string `json:"symbol"`
	Venue    string `json:"venue"`
	Bid    int `json:"bid"`
	Ask int `json:"ask"`
	BidSize    int `json:"bidSize"`
	AskSize int `json:"askSize"`
	BidDepth    int `json:"bidDepth"`
	AskDepth int `json:"askDepth"`
	Last    int `json:"last"`
	LastSize int `json:"lastSize"`
	LastTrade    string `json:"lastTrade"`
	QuoteTime    string `json:"quoteTime"`
}

type Tickertape struct {
	OK    bool `json:"ok"`
	Quote QuoteType  `json:"quote"`
}


type PlaceOrderType struct {
	Account    string `json:"account"`
	Venue    string `json:"venue"`
	Stock    string `json:"stock"`
	Price    int `json:"price"`
	Quantity int `json:"qty"`
	Direction    string `json:"direction"`
	OrderType    string `json:"orderType"`
}

type VenueStatus struct {
	OK    bool `json:"ok"`
	Venue string `json:"venue"`
}

type APIStatus struct {
	OK    bool `json:"ok"`
	Error string `json:"error"`
}

type InstanceType struct {
	OK    bool `json:"ok"`
	Done    bool `json:"done"`
	Error string `json:"error"`
	InstanceID int `json:"id"`
	State string `json:"state"`
}

type LevelCreate struct {
	OK    bool `json:"ok"`
	InstanceID int `json:"instanceID"`
	Account string `json:"account"`
	//Instructions string `json:"instructions"`
	Tickers []string `json:"tickers"`
	Venues []string `json:"venues"`
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

func (g *QuoteType) Spread() int {
    return g.Bid - g.Ask
}
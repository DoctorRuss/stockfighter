package main

import (
    "log"
    "os"
    "os/signal"
    "github.com/doctorruss/stockfighter/stockfighter"
)

func main() {
    // control logging format
    log.SetFlags(0)

    // redirect OS Interrupt signal to interrupt channel
    // to allow close down
    interrupt := make(chan os.Signal, 1)
    signal.Notify(interrupt, os.Interrupt)
    
    // create level by name
    levelname := "first_steps"
    g := stockfighter.GM{LevelName:levelname}
    levelDetails := g.CreateLevel()
    if levelDetails.OK == false {
        panic(levelDetails.Error)
    }
    defer g.StopLevel()
    
    // create websocket endpoint with level details
    venue := levelDetails.Venues[0]
    tradingAccount := levelDetails.Account

    //tickertape := "/ob/api/ws/" + tradingAccount + "/venues/" + venue+ "/tickertape"
    
    stock := levelDetails.Tickers[0]
    quotes := "/ob/api/ws/" + tradingAccount + "/venues/" + venue+ "/tickertape/stocks/" + stock

    // create channel to comunicate between goroutines
    // this channel will be closed when the ReceiveWebsocketMsg routines exits
    done := make(chan struct{})
    
    // create websocket connection and defer the close down
    connection := stockfighter.CreateWebsocket(quotes)
    defer stockfighter.CloseWebSocket(*connection, done)
    
    // fire goroutine to receive messages
    go stockfighter.ReceiveWebsocketMsg(*connection, done)
    
    // wait until Ctrl+C
     <-interrupt
}
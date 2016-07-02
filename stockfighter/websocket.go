package stockfighter


import (
	 "log"
	 "github.com/gorilla/websocket"
	"net/http"
	 "net/url"
	 "os"
	 "time"
	// "encoding/json"
)

type Websocket struct {
	connection websocket.Conn
	// this channel will be closed when the ReceiveWebsocketMsg routines exits
    doneChan chan struct{}
}
var hostAddr = "api.stockfighter.io"
    
func NewWebsocket(path string) Websocket {
	websock := Websocket{*CreateWebsocket(path), make(chan struct{})}
	return websock
}
func CreateWebsocket(path string) *websocket.Conn {
    u := url.URL{Scheme: "wss", Host: hostAddr, Path: path}
	log.Printf("connecting to %s", u.String())

    header := http.Header{}
    header.Set("X-Starfighter-Authorization", os.Getenv("Value"))
    c, resp, err := websocket.DefaultDialer.Dial(u.String(), header)
	if err != nil {
	    log.Fatal(resp)
		log.Fatal("dial:", err)
	}
	return c
}

func (ws *Websocket) ReceiveWebsocketMsg(msgs chan []byte)  {
	defer ws.connection.Close()
	defer close(ws.doneChan)
	for {
		_, message, err := ws.connection.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			return
		}
		msgs <- message
//		log.Printf("recv: %s", message)
// 		var ticker Tickertape
// 		err := c.ReadJSON(&ticker)
// 		if err != nil {
// 			log.Println("read:", err)
// 			return
// 		}
	
// 		log.Printf("Trade ticker: %+v %+v\n", ticker.Quote.LastTrade, ticker.Quote.Last)
	}
}

func (ws *Websocket) CloseWebSocket() {
	log.Println("interrupt")
    // To cleanly close a connection, a client should send a close
    // frame and wait for the server to close the connection.
    err := ws.connection.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
    if err != nil {
        log.Println("write close:", err)
        return
    }
    select {
    case <-ws.doneChan:
        log.Println("Received done")
    case <-time.After(time.Second):
        log.Println("timeout")
    }
    ws.connection.Close()
}
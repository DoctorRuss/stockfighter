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

var hostAddr = "api.stockfighter.io"
    
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


func ReceiveWebsocketMsg(c websocket.Conn, done chan struct{})  {
	defer c.Close()
	defer close(done)
	for {
		_, message, err := c.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			return
		}
		log.Printf("recv: %s", message)
// 		var ticker Tickertape
// 		err := c.ReadJSON(&ticker)
// 		if err != nil {
// 			log.Println("read:", err)
// 			return
// 		}
	
// 		log.Printf("Trade ticker: %+v %+v\n", ticker.Quote.LastTrade, ticker.Quote.Last)
	}
}

func CloseWebSocket(connection websocket.Conn, done chan struct{}) {
	log.Println("interrupt")
    // To cleanly close a connection, a client should send a close
    // frame and wait for the server to close the connection.
    err := connection.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
    if err != nil {
        log.Println("write close:", err)
        return
    }
    select {
    case <-done:
        log.Println("Received done")
    case <-time.After(time.Second):
        log.Println("timeout")
    }
    connection.Close()
}
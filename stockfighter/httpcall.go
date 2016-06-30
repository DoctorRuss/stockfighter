package stockfighter
    
import (
	"io"
	"io/ioutil"
	"os"
	"net/http"
)

func StockfighterAPI(endpoint string) []byte {
	req, err := http.NewRequest("GET", endpoint, nil)
    if err != nil {
        panic(err)
    }
	return commonAPI(req)
}

func StockfighterPost(endpoint string, body io.Reader) []byte {
	req, err := http.NewRequest("POST", endpoint, body)
    if err != nil {
        panic(err)
    }
	return commonAPI(req)
}

func commonAPI(req *http.Request) []byte {
    req.Header.Set("X-Starfighter-Authorization", os.Getenv("Value"))
    req.Header.Set("Content-Type", "application/json")

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        panic(err)
    }
    defer resp.Body.Close()
    body, _ := ioutil.ReadAll(resp.Body)
	
	return body
}
    
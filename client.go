package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	res, err := http.Get("https://api.stockfighter.io/ob/api/heartbeat")

	if err != nil {
		log.Fatal(err)
	}
	bodytext, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", bodytext)
}    
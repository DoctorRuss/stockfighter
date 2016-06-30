package stockfighter

import (
	"encoding/json"
	"fmt"
	"strconv"
    "bytes"
)

// base part of all Game Master URLs
const baseURL = "https://api.stockfighter.io/gm/"

type GM struct {
	LevelName string
	InstanceID int
}

func (g *GM) GetLevels()  {
    var buffer bytes.Buffer
    buffer.WriteString(baseURL)
    buffer.WriteString("levels")
	body := StockfighterAPI(buffer.String())
	
    // debug only 
    fmt.Printf("%s\n", body)
}

func (g *GM) CreateLevel() LevelCreate {
    var buffer bytes.Buffer
    buffer.WriteString(baseURL)
    buffer.WriteString("levels/")
    buffer.WriteString( g.LevelName )
	body := StockfighterPost(buffer.String(), nil)
	
    var creation LevelCreate
	err := json.Unmarshal(body, &creation)
	if err != nil {
		panic(err)
	}
    // debug only fmt.Printf("%s\n", body)
	
	fmt.Printf("%+v\n", creation)
	g.InstanceID = creation.InstanceID
	return creation
}

func (g *GM) GetInstanceDetails() InstanceType {
    var buffer bytes.Buffer
    buffer.WriteString(baseURL)
    buffer.WriteString("instances/")
    buffer.WriteString( strconv.Itoa(g.InstanceID) )
	body := StockfighterAPI(buffer.String())
	
    var data InstanceType
	err := json.Unmarshal(body, &data)
	if err != nil {
		panic(err)
	}
    fmt.Printf("%s\n", body)
	
	fmt.Printf("%+v\n", data)
	return data
}

func (g *GM)  StopLevel()  {
    var buffer bytes.Buffer
    buffer.WriteString(baseURL)
    buffer.WriteString("instances/")
    buffer.WriteString( strconv.Itoa(g.InstanceID) )
    buffer.WriteString("/stop")
	body := StockfighterPost(buffer.String(), nil)
    fmt.Printf("%s\n",body)
}

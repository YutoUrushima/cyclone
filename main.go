package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

type RequestBody struct {
	Channel string
	Text    string
}

func main() {
	godotenv.Load()
	requestBody := &RequestBody{
		Channel: os.Getenv("CHANNEL_ID"),
		Text:    "Hello world",
	}
	jsonString, err := json.Marshal(requestBody)
	if err != nil {
		panic("Error")
	}
	req, err := http.NewRequest("POST", os.Getenv("SLACK_URL"), bytes.NewBuffer(jsonString))
	if err != nil {
		panic(("Error"))
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+os.Getenv("BEARER_TOKEN"))
	client := new(http.Client)
	res, err := client.Do(req)
	if err != nil {
		panic("Error")
	}
	defer res.Body.Close()

	resArray, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic("Error")
	}

	fmt.Printf("%#v", string(resArray))
}

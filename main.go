package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

type RequestBody struct {
	channel string
	text    string
}

func main() {
	godotenv.Load()
	// requestBody := &RequestBody{
	// 	channel: os.Getenv("CHANNEL_ID"),
	// 	text:    "Hello world",
	// }
	// jsonString, err := json.Marshal(requestBody)
	// if err != nil {
	// 	panic("Error")
	// }
	req, err := http.NewRequest("POST", os.Getenv("SLACK_URL"), nil)
	if err != nil {
		panic(("Error"))
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Authorization", "Bearer "+os.Getenv("BEARER_TOKEN"))

	params := req.URL.Query()
	params.Add("channel", os.Getenv("CHANNEL_ID"))
	params.Add("text", "Hello,world2")
	req.URL.RawQuery = params.Encode()

	fmt.Printf("request -> %v\n", req)
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

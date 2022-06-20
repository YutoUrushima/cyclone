package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	req, err := http.NewRequest("POST", os.Getenv("SLACK_URL"), nil)
	fmt.Printf("req: %v, err: %v", req, err)
}

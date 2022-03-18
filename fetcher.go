package main

import (
    "fmt"
    "log"
    "os"

    "github.com/joho/godotenv"
)

func init() {

    err := godotenv.Load(".env")

    if err != nil {
        log.Fatal("Error loading .env file")
    }
}

func main() {
	var url = os.Getenv("EIA_URL") + "?api_key=" + os.Getenv("EIA_KEY") + "&series_id=" + os.Getenv("EIA_SERIES")

	fmt.Printf(url)
}
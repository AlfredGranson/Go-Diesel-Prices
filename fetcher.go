package main

import (
    "fmt"
    "log"
    "os"
    "net/http"
    "io/ioutil"

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

	resp, err := http.Get(url)

	if err != nil {
      log.Fatalln(err)
   }

	//We Read the response body on the line below.
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
	  log.Fatalln(err)
	}

	//Convert the body to type string
	sb := string(body)

	fmt.Printf(sb)
}
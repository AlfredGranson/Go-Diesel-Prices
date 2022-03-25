package main

import (
    "fmt"
    "log"
    "os"
    "net/http"
    "io/ioutil"
    "bufio"
    "strings"
    "regexp"
    "time"
    "strconv"
    "encoding/json"

    "github.com/joho/godotenv"
)

func init() {

    err: = godotenv.Load(".env")

    if err != nil {
        log.Fatal("Error loading .env file")
    }
}

func GetInput() string {
    fmt.Println("Please enter requested YYYYMM or blank for current:")

    reader: = bufio.NewReader(os.Stdin)

    // ReadString will block until the delimiter is entered
    input, err: = reader.ReadString('\n')
    if err != nil {
        fmt.Println("An error occured while reading input. Please try again", err)
        return input
    }

    // remove the delimeter from the string
    input = strings.TrimSuffix(input, "\n")

    return input
}

func CheckInput(input string) bool {
    if input == "" {
        return true
    }

    match, _: = regexp.MatchString("^\\d{4}(0[1-9]|1[0-2])(0[1-9]|)$", input)

    if !match {
        fmt.Println("Not a valid entry")
    }

    return match
}

func main() {
    fmt.Print("Connecting to API...")
    var url = os.Getenv("EIA_URL") + "?api_key=" + os.Getenv("EIA_KEY") + "&series_id=" + os.Getenv("EIA_SERIES")

    resp, err: = http.Get(url)

    if err != nil {
        log.Fatalln(err)
    }

    //We Read the response body on the line below.
    body, err: = ioutil.ReadAll(resp.Body)
    if err != nil {
        log.Fatalln(err)
    }

    //Convert the body to type string
    sb: = string(body)

    //Json decode
    var json_result map[string] interface {}
    json.Unmarshal([] byte(sb), & json_result)

    //Extract the data from the JSON
    doe_data: = [] interface {}(json_result["series"].([] interface {}))[0].(map[string] interface {})["data"].([] interface {})

    fmt.Println("Ready")

    var input = GetInput()

    //Current year/month to fall back on
    t: = time.Now()
    var date_to_search = strconv.Itoa(t.Year()) + fmt.Sprintf("%02d", int(t.Month()))

    //Check if user input is valid. If not, do not proceed
    if input != "" {
        for !CheckInput(input) {
            input = GetInput()
        }
        if input != "" {
            date_to_search = input
        }
    }

    fmt.Println("Searching for prices in", date_to_search, "...")

    //Look for matching dates
    for _, data_point: = range doe_data {
        data_point_array: = interface {}(data_point).([] interface {})
        if strings.HasPrefix(data_point_array[0].(string), date_to_search) {
            fmt.Println(data_point_array[0], ":", data_point_array[1])
        }
    }

}
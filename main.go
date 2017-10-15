package main

import (
    "fmt"
    "io/ioutil"
    "net/http"
    "encoding/json"
)

type BuoyResponse struct {
    MetaData `json:"metadata"`
}

type MetaData struct {
    Id int64 `json:"id,string"`
    Name string `json:"name"`
    Lat float64 `json:"lat,string"`
    Lon float64 `json:"lon,string"`
}

type test_struct struct {
    Test string
}

const baseBuoyUrl = "https://tidesandcurrents.noaa.gov/api/datagetter?begin_date=20130101%2010:00&end_date=20130101%2010:24&station=8454000&product=water_level&datum=mllw&units=metric&time_zone=gmt&application=web_services&format=json"

func main() {
    http.HandleFunc("/", getIndex)
    http.ListenAndServe(":8080", nil)
} 

func getIndex(w http.ResponseWriter, r *http.Request) {

    result, err := requestUrl(baseBuoyUrl)
   
    if err != nil {
        fmt.Println("error:", err) 
    } 

    parsed, err := parseResponse(result)

    if err != nil {
        fmt.Println("error:", err) 
    } 

    //Do something with the data
    fmt.Println(parsed.MetaData.Id)

    w.Header().Set("Content-Type", "application/json")
    w.Write(result)

}

func requestUrl(url string) ([]byte, error){
    res, err := http.Get(url)

    if err != nil {
        return nil, err
    }

    json_body, err := ioutil.ReadAll(res.Body)

    if err != nil {
        return nil, err
    }

    return []byte(json_body), err
}

func parseResponse(response []byte) (*BuoyResponse, error) {
    var tmp = new(BuoyResponse)
    err := json.Unmarshal(response, &tmp)
    if err != nil {
        return nil, err
    }
    return tmp, err
}

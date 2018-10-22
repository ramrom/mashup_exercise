package main

import (
  "fmt"
  "io/ioutil"
  "os/user"
  "encoding/json"
  "net/http"
)

var GCreds GoogleCredentials

type GoogleCredentials struct {
  ApiKey     string `json:"api_key"`
}

type LongLatResult struct {
  Location LongLat
  Accuracy float64
}
type LongLat struct {
  Lat float64
  Lng float64
}

func LoadGoogleCredentials() {
  usr, _ := user.Current()
  cred_file, err := ioutil.ReadFile(usr.HomeDir + "/.creds/google_api.json")
  if err != nil {
    panic(fmt.Sprintf("failed to read nasa creds file!:  %v",err))
  }
  err = json.Unmarshal(cred_file, &GCreds)
  if err != nil {
    panic(fmt.Sprintf("failed to json parse nasa creds file!:  %v",err))
  }
}

// TODO: google has a API where it returns the long/lat of any IP right?
func GetLocalLongLat() (*LongLatResult, error) {
  url := fmt.Sprintf("https://www.googleapis.com/geolocation/v1/geolocate?key=%v",GCreds.ApiKey)
  res, err := http.Post(url, "text", nil)
  if err != nil {
    return nil, fmt.Errorf("failed longlat google api call!: %v",err)
  }
  defer res.Body.Close()
  body, err := ioutil.ReadAll(res.Body)
  //fmt.Println(string(body))
  var result LongLatResult
  err = json.Unmarshal(body, &result)
  if err != nil {
    return nil, fmt.Errorf("failed unmarshalling google longlat response: %v",err)
  } else {
    return &result, nil
  }
}

// Obtain the common human readable location like country/state/city/street
func GetHumanLocation() {
  //ADDRESS_URL := "https://maps.googleapis.com/maps/api/geocode/json?latlng=LONG,LAT&key=APIKEY"
}

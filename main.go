package main

import (
  "fmt"
  "net/http"
  "os"
  "strconv"
  "encoding/json"
)

var TwilioPhone string

func init() {
  // TODO: be able to specify creds file location in env var
  LoadGoogleCredentials()
  LoadYelpCredentials()
  LoadTwilioCredentials()
  TwilioPhone = os.Getenv("PHONE")
  if TwilioPhone == "" {
    panic("must supply the twilio phone number for sms notifications")
  }
}

func main() {
  //fmt.Println(GetTopFiveRestaurantsInArea("bogus"))
  StartServer()
}

func StartServer() {
  // TODO: def bad, should use Accept headers, not routes to determine content-types of responses
  http.Handle("/json", http.HandlerFunc(JSONHandler))
  http.Handle("/web", http.HandlerFunc(HTMLHandler))
  http.Handle("/sms", http.HandlerFunc(SMSHandler))

  port := os.Getenv("PORT")
  if port == "" {
    port = "8080"
  }
  fmt.Printf("starting server on %v\n", port)
  http.ListenAndServe(fmt.Sprintf(":%v",port),nil)
}

// TODO: use a geoloc API that returns long/lat of IP Address
func GetTopFiveRestaurantsInArea(IPAddress string) []string {
  fmt.Println("Querying google for geoloc...")
  loc, err := GetLocalLongLat()
  if err != nil {
    panic(fmt.Sprintf("failed getting google geloc:  %v",err))
  }
  fmt.Println(loc)

  fmt.Println("Making yelp query...")
  //yres, err := QueryYelpBusinesses("-87.64","41.90")
  longstring := strconv.FormatFloat(loc.Location.Lng,'f',8,64)
  latstring := strconv.FormatFloat(loc.Location.Lat,'f',8,64)
  yelpresponse, err := QueryYelpBusinesses(longstring,latstring)
  if err != nil {
    panic(fmt.Sprintf("failed querying yelp:  %v",err))
  }
  //fmt.Println(yelpresponse)
  var toprestaurants []string
  for _, biz := range yelpresponse.Businesses {
    toprestaurants = append(toprestaurants, biz.Name)
  }
  return toprestaurants
}

func JSONHandler(w http.ResponseWriter,r *http.Request) {
  w.Header().Set("Content-Type","application/json")
  rr := GetTopFiveRestaurantsInArea("doesntmatternow")
  rrjson, _ := json.Marshal(rr)
  w.Write(rrjson)
}

func HTMLHandler(w http.ResponseWriter,r *http.Request) {
  rr := GetTopFiveRestaurantsInArea("doesntmatternow")
  w.Header().Set("Content-Type","text/html")
  list := "<ul>"
  for _, rest := range rr {
    list = list + "<li>" + rest + "</li>"
  }
  list = list + "</ul>"
  html := fmt.Sprintf(`<html><body><h1>Top Five Restaurants</h1>%v</body></html>`,list)
  w.Write([]byte(html))
}

func SMSHandler(w http.ResponseWriter,r *http.Request) {
  rr := GetTopFiveRestaurantsInArea("doesntmatternow")
  msg := "Top 5 restaurants in your area: "
  for _, rest := range rr {
    msg = msg + ", " + rest
  }
  fmt.Println("sending sms to ph #:", TwilioPhone)
  SendTwilioText(TwilioPhone, msg)
  w.Write([]byte("sent SMS"))
}

func LogIP(r *http.Request) {
  fmt.Println("Remote address: ",r.RemoteAddr)
}

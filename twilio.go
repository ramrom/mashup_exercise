package main

// see this blog on directly writing the http client code instead of this 3rd party gotwilio package:
// https://www.twilio.com/blog/2017/09/send-text-messages-golang.html

import (
    "github.com/sfreiberg/gotwilio"
    "fmt"
    "io/ioutil"
    "os/user"
    "encoding/json"
)

var TwilCreds TwilioCredentials

type TwilioCredentials struct {
  AuthToken  string `json:"auth_token"`
  SID        string `json:"sid"`
  PhoneNum   string `json:"phone_num"`
}

func LoadTwilioCredentials() {
  usr, _ := user.Current()
  cred_file, err := ioutil.ReadFile(usr.HomeDir + "/.creds/twilio_api.json")
  if err != nil {
    panic(fmt.Sprintf("failed to read twilio creds file!:  %v",err))
  }
  err = json.Unmarshal(cred_file, &TwilCreds)
  if err != nil {
    panic(fmt.Sprintf("failed to json parse twilio creds file!:  %v",err))
  }
}

func SendTwilioText(to string, msg string) {
  twilio := gotwilio.NewTwilioClient(TwilCreds.SID, TwilCreds.AuthToken)
  res, ex, err := twilio.SendSMS(TwilCreds.PhoneNum, to, msg, "", "")
  fmt.Println("sms response body: %v", res)
  fmt.Println("exception: %v", ex)
  fmt.Println("error: %v", err)
}

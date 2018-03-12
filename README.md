# Twilio SMS client

A lightweight Go library for sending SMS with Twilio.

## Usage

``` go
baseURL := "https://api.twilio.com/2010-04-01"
accountSid := "AC00xxxxxxxx"
authToken := "2a7bxxxxxxxxx"
twilio := NewTwilioClient(baseURL, accountSid, authToken)

msgData := url.Values{}
msgData.Set("To", "+44XXXXXXXXXX")
msgData.Set("From", "+44XXXXXXXXXX")
msgData.Set("Body", "Hello World")

twilio.SendMessage(msgData)
```

## Credit

* [So It Goes: Sending SMS with Golang](https://www.twilio.com/blog/2017/09/send-text-messages-golang.html)
* [sfreiberg/gotwilio](https://github.com/sfreiberg/gotwilio)
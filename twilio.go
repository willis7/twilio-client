package twilio

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"
	"strings"
)

// Twilio stores basic information important for connecting to the
// twilio.com REST api such as AccountSid and AuthToken.
type Twilio struct {
	AccountSid string
	AuthToken  string
	BaseURL    string
	HTTPClient *http.Client
}

// NewTwilioClient is a Twilio constructor.
func NewTwilioClient(baseURL, accountSid, authToken string) *Twilio {
	return NewTwilioClientCustomHTTP(baseURL, accountSid, authToken, nil)
}

// NewTwilioClientCustomHTTP is a Twilio constructor optionally using a custom http.Client
func NewTwilioClientCustomHTTP(baseURL, accountSid, authToken string, HTTPClient *http.Client) *Twilio {

	if HTTPClient == nil {
		HTTPClient = http.DefaultClient
	}

	return &Twilio{accountSid, authToken, baseURL, HTTPClient}
}

func (twilio *Twilio) post(formValues url.Values, twilioURL string) (*http.Response, error) {
	req, err := http.NewRequest("POST", twilioURL, strings.NewReader(formValues.Encode()))
	if err != nil {
		return nil, err
	}
	req.SetBasicAuth(twilio.AccountSid, twilio.AuthToken)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	client := twilio.HTTPClient
	if client == nil {
		client = http.DefaultClient
	}

	return client.Do(req)
}

// SendMessage sends a SMS with Twilio
func (twilio *Twilio) SendMessage(formValues url.Values) {
	messageURL := twilio.BaseURL + "/Accounts/" + twilio.AccountSid + "/Messages.json"

	resp, err := twilio.post(formValues, messageURL)
	if err != nil {
		log.Printf("message send failed: %v", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		var data map[string]interface{}
		decoder := json.NewDecoder(resp.Body)
		err := decoder.Decode(&data)
		if err == nil {
			log.Println(data["sid"])
		} else {
			log.Printf("message decode failed: %v", data["message"])
		}
	} else {
		log.Println(resp.Status)
	}
}

package iftttWebhook

import (
	"bytes"
	"encoding/json"
	"net/http"
	"strconv"
)

//IFTTT struct holding all necessary values to connect to IFTTT
type IFTTT struct {
	apiKey string
}

// New generates a new struct necessary to connect to IFTTT webhooks and send metrics
func New(apiKey string) IFTTT {
	i := IFTTT{apiKey: apiKey}

	return i
}

// Emit sends events to IFTTT, v can be up to 3 values which should be included
func (i IFTTT) Emit(eventName string, v ...string) error {
	url := "https://maker.ifttt.com/trigger/" + eventName + "/with/key/" + i.apiKey
	values := map[string]string{}

	for x, value := range v {
		values["value"+strconv.Itoa(x+1)] = value

		// only include up to 3 values
		if x == 2 {
			break
		}
	}

	body, err := json.Marshal(values)

	if err != nil {
		return err
	}

	_, err = http.Post(url, "application/json", bytes.NewReader(body))

	if err != nil {
		return err
	}

	return nil
}

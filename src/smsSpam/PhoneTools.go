package smsSpam

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
)

//this get a random service
func getServices() Phones {

	f, _ := ioutil.ReadFile("jsons/services.json")
	var phones Phones
	json.NewDecoder(bytes.NewReader(f)).Decode(&phones)
	return phones
}

// this is going to rebuild the phone number
func Normalize(phone string) string {

	if phone[0] == '+' {
		phone = phone[1:]

	}
	return phone
}

// this is for %phone5% numbers
func transformNumber(phone string, i int) string {

	if i == 5 {
		return "+" + string(phone[0]) + " (" + phone[1:4] + ") " + phone[4:7] + " " + phone[7:9] + " " + phone[9:11]
	}
	return phone
}

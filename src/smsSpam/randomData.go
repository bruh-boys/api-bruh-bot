package smsSpam

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"strconv"
)

func randomName() string {
	return names[rand.Intn(len(names))]
}
func randomSuffix(rangeNumber int) string {
	numbers := ""
	for i := 0; i <= rangeNumber; i++ {
		numbers += strconv.Itoa(rand.Intn(8) + 1)

	}
	return numbers
}

func randomEmail() string {
	return randomName() + "@" + email[rand.Intn(len(email))]
}

func randomPass() string {
	return randomName() + randomSuffix(10)
}
func randomUserAgent() string {
	var userAgent []string
	f, _ := ioutil.ReadFile("jsons/user-agent.json")
	json.NewDecoder(bytes.NewReader(f)).Decode(&userAgent)
	return string(userAgent[rand.Intn(len(userAgent))])
}

func randomService() map[string]string {
	p := getServices()
	fmt.Println(p)
	return p[rand.Intn(len(p))]
}

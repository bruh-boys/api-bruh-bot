package controller

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os/exec"
	"time"
)

func SmsSpam(w http.ResponseWriter, r *http.Request) {
	var inputPhone phone
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Error")
	}
	json.Unmarshal([]byte(reqBody), &inputPhone)

	input := inputPhone.Phone
	validatePhone := string(v.Find([]byte(input)))

	if input != validatePhone {
		json.NewEncoder(w).Encode("Error you didn't pass with a +, or is not a number, or did you pass it with spaces, example: +22123456789")
	} else {
		json.NewEncoder(w).Encode("sms spam sent")
		fmt.Println("request body:", validatePhone)
		start := time.Now()
		var stdout, stderr bytes.Buffer
		cmd := exec.Command("sh", "-c", "cd routes/quack ; python3 quack --tool SMS --target "+validatePhone+" --threads 60 --timeout 90")
		cmd.Stdout = &stdout
		cmd.Stderr = &stderr
		peo := cmd.Run()
		if peo != nil {
			fmt.Println(err)
		}
		// print the stdout and stderr
		fmt.Println(stdout.String() + stderr.String())
		duration := time.Since(start)
		json.NewEncoder(w).Encode(duration.Seconds())

	}

}

package routes

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os/exec"
	"time"
)

func SpamEmail(w http.ResponseWriter, r *http.Request) {
	var emailInput email
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Error")
	}
	json.Unmarshal([]byte(reqBody), &emailInput)

	input2 := emailInput.Email
	validateEmail := string(e.Find([]byte(input2)))

	if input2 != validateEmail {
		json.NewEncoder(w).Encode("Error you don't pass a valid email")
	} else {
		json.NewEncoder(w).Encode("email spam sent")
		fmt.Println("request body:", validateEmail)
		start := time.Now()
		var stdout, stderr bytes.Buffer
		cmd := exec.Command("sh", "-c", "python3 routes/impulse.py --method EMAIL --time 60 --target "+validateEmail)
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

package smsSpam

import "regexp"

var (
	PhoneRegex = regexp.MustCompile(`\+?[\d]{10,14}`)
)

type Phones []map[string]string

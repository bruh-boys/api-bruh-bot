package controller

import "regexp"

var (
	v = regexp.MustCompile(`\+?[\d]{10,14}`)
)

type phone struct {
	Phone string
}

type phoneExample []phone

var phones = phoneExample{
	{
		Phone: "+243434335",
	},
}

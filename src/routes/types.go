package routes

import "regexp"

var (
	v = regexp.MustCompile(`\+?[\d]{10,14}`)
	e = regexp.MustCompile(`^[a-zA-Z0-9_.+-]+@[a-zA-Z0-9-]+\.[a-zA-Z0-9-.]+$`)
)

type phone struct {
	Phone string
}

type email struct {
	Email string
}

type phoneExample []phone
type emailExample []email

var emails = emailExample{
	{
		Email: "monda@monda.com",
	},
}

var phones = phoneExample{
	{
		Phone: "+243434335",
	},
}

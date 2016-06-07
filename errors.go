package iqfeed

import "fmt"

// ErrorMsg contains error messages reported to the client including symbol not found messages
type ErrorMsg struct {
	Message string // The error message
	Code    int    // The http status representation of the error.
}

// UnMarshall sends the data into the usable struct for consumption by the application.
func (e *ErrorMsg) UnMarshall(d []byte, code int) {
	fmt.Printf("Unmarshall: %s", string(d))
}

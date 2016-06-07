package iqfeed

// ErrorMsg contains error messages reported to the client including symbol not found messages
type ErrorMsg struct {
	Symbol  string // Symbol is set on 404 messages to indicate the missing symbol
	Message string // The error message
	Code    int    // The http status representation of the error.
}

// UnMarshall sends the data into the usable struct for consumption by the application.
func (e *ErrorMsg) UnMarshall(notFound bool, d []byte, code int) {
	if notFound {
		e.Symbol = string(d)
		e.Code = 404
		e.Message = "Symbol not found"
		return
	}
	e.Code = 500
	e.Message = string(d)
}

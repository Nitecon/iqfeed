package iqfeed

import (
	"fmt"
	"time"
)

// TimeMsg represents a current timestamp from the network.
type TimeMsg struct {
	TimeStamp time.Time
}

// UnMarshall sends the data into the usable struct for consumption by the application.
func (t *TimeMsg) UnMarshall(d []byte) {
	fmt.Printf("Unmarshall: %s", string(d))
}

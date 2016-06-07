package iqfeed

import (
	"fmt"
	"time"
)

// NewsMsg represents a news message that is associated to one or more symbols.
type NewsMsg struct {
	DistributorCode string    // Distributor type code
	StoryID         int       // Numerical Story ID
	SymbolList      []string  // List of symbols associated with news story.
	DateTime        time.Time // Format is in YYYYMMDD HHMMSS
	Headline        string    // The text headline
}

// UnMarshall sends the data into the usable struct for consumption by the application.
func (n *NewsMsg) UnMarshall(d []byte) {
	fmt.Printf("Unmarshall: %s", string(d))
}

package iqfeed

import (
	"strings"
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
func (n *NewsMsg) UnMarshall(d []byte, loc *time.Location) {
	items := strings.Split(string(d), ",")
	n.DistributorCode = items[0]
	n.StoryID = GetIntFromStr(items[1])
	n.SymbolList = strings.Split(items[2], ":")
	t, _ := time.ParseInLocation("20060102 150405", items[3], loc)
	n.DateTime = t
	n.Headline = items[4]

}

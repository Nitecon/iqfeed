package iqfeed

import "time"

// SystemMessage is the main system message that will be returned and set by the client.
type SystemMessage struct {
	Customer CustomerData
	Stats    SystemStats
}

// CustomerData is a subset of SystemMessage which is returned when requesting customer data.
type CustomerData struct {
	ServiceType      string // Will either be real_time / delayed
	IP               string // this is the IP address of the quote server
	Port             int    // This will be the port of the quote server in use
	Token            string // This is a daily auth for the feed
	Version          string // will be the most current version of the iqfeed client
	Deprecated1      int    // not used item
	VerboseExchanges string // verbose exchanges in text that customer will get in real time.
	Deprecated2      string // unused
	MaxSymbols       int    // max number of symbols a user can watch at a time
	Flags            string // are any special flags that may be in their account (ie: NO_EOD / BETA etc)
	Deprecated3      string // unused
	Deprecated4      string // unused
}

// SystemStats is a subset of SystemMessage which is returned when requesting stats.
type SystemStats struct {
	ServerIP               string    // This is the IP address of the Quote server in use
	ServerPort             int       // This is the Port of the Quote server in use
	MaxSymbols             int       // The maximum # of symbols that can be watched at any given time
	NumberOfSymbols        int       // The # of symbols that are currently being watched
	ClientsConnected       int       // The # of clients that are currently connected
	SecondsSinceLastUpdate int       // The # of seconds since the last update from the Quote server
	Reconnections          int       // The # of times that IQFeed has reconnected
	AttemptedReconnections int       // The # of times that IQFeed has attempted to reconnect, but failed
	StartTime              time.Time // The time that the connection (or reconnection) to IQFeed was made in the format [short month][space][Day][space][hour][colon][minute][AM/PM]
	MarketTime             time.Time // the current time of the market in the format [short month][space][Day][space][hour][colon][minute][AM/PM]
	Status                 string    // Represents whether IQFeed is connected or not. Values are "Connected" OR "Not Connected"
	IQFeedVersion          string    // Represents the version of IQFeed that is running
	LoginID                string    // The LoginID that is currently logged in
	TotalKBsRecv           float32   // Found in the “Internet Bandwidth” section of the IQFeed Connection Manager. Formula: total bytes received / 1024
	KBsPerSecRecv          float32   // Found in the “Internet Bandwidth” section of the IQFeed Connection Manager. Formula: bytes received in the past second / 1024
	AvgKBsPerSecRecv       float32   // Found in the “Internet Bandwidth” section of the IQFeed Connection Manager. Formula: total KB's received / total seconds
	TotalKBsSent           float32   // Found in the “Local Bandwidth” section of the IQFeed Connection Manager. Formula: total bytes sent / 1024
	KBsPerSecSent          float32   // Found in the “Local Bandwidth” section of the IQFeed Connection Manager. Formula: bytes sent in the past second / 1024
	AvgKBsPerSecSent       float32   // Found in the “Local Bandwidth” section of the IQFeed Connection Manager. Formula: total KB's sent / total seconds
}

// UnMarshall sends the data into the usable struct for consumption by the application.
func (f *SystemMessage) UnMarshall(d []byte, loc *time.Location) {
	//fmt.Printf("System Message: %s", string(d))
	return
}

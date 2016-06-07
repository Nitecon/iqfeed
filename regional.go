package iqfeed

import (
	"fmt"
	"time"
)

// RegionalMsg A regional update message. See complete message definition in Regional Messages. (http://www.iqfeed.net/dev/api/docs/RegionalMessageFormat.cfm).
type RegionalMsg struct {
	Symbol           string  // the  symbol that is being tracked
	Exchange         string  // Deprecated - Use field 12 Below. - See Market Center Codes
	RegBid           float64 //
	RegBidSize       int
	RegBidTime       time.Time // Currently Time of last Trade.
	RegAsk           float64
	RegAskSize       int
	RegAskTime       time.Time // Currently Time of last Trade.
	FractionDispCode int       // Display formatting code see Price Format Codes (http://www.iqfeed.net/dev/api/docs/PriceFormatCodes.cfm).
	DecPrecision     int       // Last Precision used.
	MarketCenter     int       // The regional exchange that the updae occurred at. See the Listed Markets Codes for a list of possible values.(http://www.iqfeed.net/dev/api/docs/ListedMarkets.cfm).
}

// UnMarshall sends the data into the usable struct for consumption by the application.
func (r *RegionalMsg) UnMarshall(d []byte) {
	fmt.Printf("Unmarshall: %s", string(d))
}

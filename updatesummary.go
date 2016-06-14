package iqfeed

import "time"

// UpdSummaryMsg is the main struct for both update and summary messages.
type UpdSummaryMsg struct {
	SevenDayYield          float64   // A price field, the value from a Money Market fund over the last seven days.
	Ask                    float64   // The lowest price a market maker or broker is willing to accept for a security.
	AskChange              float64   // Change in Ask since last offer.
	AskMktCenter           int       // The Market Center that sent the ask information. See Listed Market Codes for possible values.
	AskSize                int       // The share size available for the ask price in a given security.
	AskTime                time.Time // The time for the last ask.
	AvailRegions           string    // Dash delimited string of available regions.
	AvgMaturity            float64   // The average number of days until maturity of a Money Market Fund’s assets.
	Bid                    float64   // The highest price a market maker or broker is willing to pay for a security.
	BidChange              float64   // Change in Bid since last offer.
	BidMktCenter           int       // The Market Center that sent the bid information. See Listed Market Codes for possible values.
	BidSize                int       // The share size available for the bid price in a given security
	BidTime                time.Time // The time of the last bid.
	Change                 float64   // Today's change (Last - Close)
	ChangeFrmOpen          float64   // Change in last since open
	Close                  float64   // The closing price of the day. For commodities this will be the last TRADE of the session
	CloseRng1              float64   // For commodities only. Range value for closing trades that aren’t reported individually.
	CloseRng2              float64   // For commodities only. Range value for closing trades that aren’t reported individually.
	DaysToExpir            string    // Number of days to contract expiration
	DecPrecision           string    // Last Precision used
	Delay                  int       // The number of minutes a quote is delayed when not authorized for real-time data
	ExchangeID             string    // This is the exchange ID. Convert to decimal and use the Listed Markets lookup to decode this value.
	ExtendedTrd            float64   // Price of the most recent extended trade (last qualified trades + Form T trades).
	ExtendedTrdDate        time.Time // Date of the extended trade. (MM/DD/CCYY)
	ExtendedTrdMktCntr     int       // Market Center of the most recent extended trade (last qualified trades + Form T trades).
	ExtendedTrdSize        int       // Size of the most recent extended trade (last qualified trades + Form T trades).
	ExtendedTrdTime        time.Time // Time (including microseconds) of the most recent extended trade (last qualified trades + Form T trades).
	ExtendedTrdChange      float64   // Extended Trade minus Yesterday's close.
	ExtendedTrdDiff        float64   // Extended Trade minus Last
	FinancialStatusInd     string    // Denotes if an issuer has failed to submit its regulatory filings on a timely basis, has failed to meet the exchange's continuing listing standards, and/or has filed for bankruptcy. See Financial Status Indicator Codes.
	FractionDispCode       string    // Display formatting code see Price Format Codes.
	High                   float64   // Today's highest trade price
	Last                   float64   // Last trade price from the regular trading session
	LastDate               time.Time // Date of the last qualified trade. (MM/DD/CCYY).
	LastMktCntr            int       // Market Center of most recent last qualified trade.
	LastSize               int       // Size of the most recent last qualified trade.
	LastTime               time.Time // Time (including microseconds) of most recent last qualified trade (HH:MM:SS.fff)
	LastTrdDate            time.Time // Date of last trade
	Low                    float64   // Today's lowest trade price
	MktCapitilization      float64   // Real-time calculated market cap (Last * Common Shares Outstanding).
	MktOpen                int       // 1 = market open, 0 = market closed NOTE: This field is valid for Futures and Future Options only.
	MsgContents            string    // Possible single character values include: C - Last Qualified Trade. |E - Extended Trade = Form T trade.|O - Other Trade = Any trade not accounted for by C or E.|b - A bid update occurred.|a - An ask update occurred.|o - An Open occurred.|h - A High occurred.|l - A Low occurred.|c - A Close occurred.|s - A Settlement occurred.|v - A volume update occurred.|NOTE: you can get multiple codes in a single message but you will only get one trade identifier per message. NOTE: It is also possible to receive no codes in a message if the fields that updated were not trade or quote related.
	MostRecentTrade        float64   // Price of the most recent trade (including all non-last-qualified trades).
	MostRecntTradeCond     string    // Conditions that identify the type of trade that occurred.
	MostRecntTradeDate     time.Time // Date of the most recent trade (MM/DD/CCYY)
	MostRecentTradeMktCntr int       // Market Center of the most recent trade (including all non-last-qualified trades).
	MostRecentTradeSize    int       // Size of the most recent trade (including all non-last-qualified trades).
	MostRecentTradeTime    time.Time // Time (including microseconds) of the most recent trade (including all non-last-qualified trades).
	NetAssetValue          float64   // Mutual Funds only. The market value of a mutual fund share equal to the net asset of a fund divided by the total number of shares outstanding. NOTE: this field is the same as the Bid field for Mutual Funds.
	NumTradesToday         int       // The number of trades for the current day.
	Open                   float64   // The opening price of the day. For commodities this will be the first TRADE of the session.
	OpenInterest           int       // IEOptions, Futures, FutureOptions, and SSFutures only.
	OpenRange1             float64   // For commodities only. Range value for opening trades that aren’t reported individually.
	OpenRange2             float64   // For commodities only. Range value for opening trades that aren’t reported individually.
	PcntChange             float64   // (Change / Close)
	PcntOffAvgVol          float64   // Current Total Volume divided by Average Volume (from fundamental message).
	PrevDayVol             int       // Previous Day's Volume
	PERatio                float64   // Real-time calculated PE (Today's Last / Earnings Per Share)
	Range                  float64   // Trading range for the current day (high - low).
	RestrictedCode         string    // Short Sale Restricted flag - "N" for Not restricted or "R" for Restricted.
	Settle                 float64   // Futures or FutureOptions only.
	SettleDate             time.Time // The date that the Settle is valid for.
	Spread                 float64   // The difference between Bid and Ask prices
	Symbol                 string    // The Symbol ID to match with watch request
	Tick                   int       // "173"=Up, "175"=Down, "183"=No Change. Only valid for Last qualified trades.
	TickID                 int       // Identifier for tick
	TotalVol               int       // Today's cumulative volume in number of shares.
	Type                   string    // Valid values are Q or P. The character Q indicates an Update message, and the character P indicates a Summary Message.
	Volatility             float64   // Real-time calculated volatility (Today's High - Today's Low) / Last
	VWAP                   float64   // Volume Weighted Average Price.

}

// UnMarshall sends the data into the usable struct for consumption by the application.
func (u *UpdSummaryMsg) UnMarshall(items []string, fields map[int]string, tz string) {
	//DynFields: map[4:Most Recent Trade Market Center 7:Bid Size 11:High 1:Most Recent Trade 8:Ask 9:Ask Size 12:Low 10:Open 15:Most Recent Trade Conditions 13:Close 14:Message Contents 0:Symbol 2:Most Recent Trade Size 3:Most Recent Trade TimeMS 5:Total Volume 6:Bid]
	//Unmarshall: AAPL,95.0200,100,09:35:57.022,26,1325032,95.0200,100,95.0400,400,95.0000,95.3800,94.8600,94.4800,ba,01,

	//fmt.Printf("Unmarshall: %#v\n", items)
	for k, v := range items {
		switch fields[k] {
		case "7 Day Yield":
			u.SevenDayYield = GetFloatFromStr(v)
		case "Ask":
			u.Ask = GetFloatFromStr(v)
		case "Ask Change":
			u.AskChange = GetFloatFromStr(v)
		case "Ask Market Center":
			u.AskMktCenter = GetIntFromStr(v)
		case "Ask Size":
			u.AskSize = GetIntFromStr(v)
		case "Ask Time":
			u.AskTime = GetTimeInHMS(v, tz)
		case "Available Regions":
			u.AvailRegions = v
		case "Average Maturity":
			u.AvgMaturity = GetFloatFromStr(v)
		case "Bid":
			u.Bid = GetFloatFromStr(v)
		case "Bid Change":
			u.BidChange = GetFloatFromStr(v)
		case "Bid Market Center":
			u.BidMktCenter = GetIntFromStr(v)
		case "Bid Size":
			u.BidSize = GetIntFromStr(v)
		case "Bid Time":
			u.BidTime = GetTimeInHMS(v, tz)
		case "Change":
			u.Change = GetFloatFromStr(v)
		case "Change From Open":
			u.ChangeFrmOpen = GetFloatFromStr(v)
		case "Close":
			u.Close = GetFloatFromStr(v)
		case "Close Range 1":
			u.CloseRng1 = GetFloatFromStr(v)
		case "Close Range 2":
			u.CloseRng2 = GetFloatFromStr(v)
		case "Days to Expiration":
			u.DaysToExpir = v
		case "Decimal Precision":
			u.DecPrecision = v
		case "Delay":
			u.Delay = GetIntFromStr(v)
		case "Exchange ID":
			u.ExchangeID = v
		case "Extended Trade":
			u.ExtendedTrd = GetFloatFromStr(v)
		case "Extended Trade Date":
			u.ExtendedTrdDate = GetDateMMDDCCYY(v, tz)
		case "Extended Trade Market Center":
			u.ExtendedTrdMktCntr = GetIntFromStr(v)
		case "Extended Trade Size":
			u.ExtendedTrdSize = GetIntFromStr(v)
		case "Extended Trade Time":
			u.ExtendedTrdTime = GetTimeInHMSmicro(v, tz)
		case "Extended Trading Change":
			u.ExtendedTrdChange = GetFloatFromStr(v)
		case "Extended Trading Difference":
			u.ExtendedTrdDiff = GetFloatFromStr(v)
		case "Financial Status Indicator":
			u.FinancialStatusInd = v
		case "Fraction Display Code":
			u.FractionDispCode = v
		case "High":
			u.High = GetFloatFromStr(v)
		case "Last":
			u.Last = GetFloatFromStr(v)
		case "Last Date":
			u.LastDate = GetDateMMDDCCYY(v, tz)
		case "Last Market Center":
			u.LastMktCntr = GetIntFromStr(v)
		case "Last Size":
			u.LastSize = GetIntFromStr(v)
		case "Last Time":
			u.LastTime = GetTimeInHMSmicro(v, tz)
		case "Last Trade Date":
			u.LastTrdDate = GetDateMMDDCCYY(v, tz)
		case "Low":
			u.Low = GetFloatFromStr(v)
		case "Market Capitalization":
			u.MktCapitilization = GetFloatFromStr(v)
		case "Market Open":
			u.MktOpen = GetIntFromStr(v)
		case "Message Contents":
			u.MsgContents = v
		case "Most Recent Trade":
			u.MostRecentTrade = GetFloatFromStr(v)
		case "Most Recent Trade Conditions":
			u.MostRecntTradeCond = v
		case "Most Recent Trade Date":
			u.MostRecntTradeDate = GetDateMMDDCCYY(v, tz)
		case "Most Recent Trade Market Center":
			u.MostRecentTradeMktCntr = GetIntFromStr(v)
		case "Most Recent Trade Size":
			u.MostRecentTradeSize = GetIntFromStr(v)
		case "Most Recent Trade Time":
			u.MostRecentTradeTime = GetTimeInHMSmicro(v, tz)
		case "Net Asset Value":
			u.NetAssetValue = GetFloatFromStr(v)
		case "Number of Trades Today":
			u.NumTradesToday = GetIntFromStr(v)
		case "Open":
			u.Open = GetFloatFromStr(v)
		case "Open Interest":
			u.OpenInterest = GetIntFromStr(v)
		case "Open Range 1":
			u.OpenRange1 = GetFloatFromStr(v)
		case "Open Range 2":
			u.OpenRange2 = GetFloatFromStr(v)
		case "Percent Change":
			u.PcntChange = GetFloatFromStr(v)
		case "Percent Off Average Volume":
			u.PcntOffAvgVol = GetFloatFromStr(v)
		case "Previous Day Volume":
			u.PrevDayVol = GetIntFromStr(v)
		case "Price-Earnings Ratio":
			u.PERatio = GetFloatFromStr(v)
		case "Range":
			u.Range = GetFloatFromStr(v)
		case "Restricted Code":
			u.RestrictedCode = v
		case "Settle":
			u.Settle = GetFloatFromStr(v)
		case "Settlement Date":
			u.SettleDate = GetDateMMDDCCYY(v, tz)
		case "Spread":
			u.Spread = GetFloatFromStr(v)
		case "Symbol":
			u.Symbol = v
		case "Tick":
			u.Tick = GetIntFromStr(v)
		case "TickID":
			u.TickID = GetIntFromStr(v)
		case "Total Volume":
			u.TotalVol = GetIntFromStr(v)
		case "Type":
			u.Type = v
		case "Volatility":
			u.Volatility = GetFloatFromStr(v)
		case "VWAP":
			u.VWAP = GetFloatFromStr(v)
		}
	}
}

package iqfeed

import (
	"fmt"
	"time"
)

// FundamentalMsg cannot be customized and is used to provide detail of a particular matched symbol.
type FundamentalMsg struct {
	Symbol             string    // The Symbol ID to match with watch request
	ExchaangeID        string    // Deprecated Use Listed Market (field 45 below) instead
	PE                 float64   // Price/earnings ratio
	AvgVolume          int       // Average daily volume (4 week avg)
	Fifty2WkHigh       float64   // Highest price of the last 52 weeks for futures, this is the contract High.
	Fifty2WkLow        float64   // Lowest price of the last 52 weeks. For futures, this is the contract Low.
	CalYearHigh        float64   // High price for the current calendar year.
	CalyearLow         float64   // Low price for the current calendar year.
	DivYield           float64   // The annual dividends per share paid by the company divided by the current market price per share of stock sent as a percentage.
	DivAmt             float64   // The current quarter actual dividend
	DivRate            float64   // The annualized amount at which a dividend is expected to be paid by a company.
	PayDate            time.Time // Date on which a company made its last dividend payment (MM/DD/YYYY).
	ExDivDate          time.Time // The actual date in which a stock goes ex-dividend, typically about 3 weeks before the dividend is paid to shareholders of record. Also the amount of the dividend is reflected in a reduction of the share price on this date. (MM/DD/YYYY).
	Reserved1          string    // Reserved field.
	Reserved2          string    // Reserved field.
	Reserved3          string    // Reserved field.
	ShortInterest      int       // The total number of shares of a security that have been sold short by customers and securities firms that have not been repurchased to settle outstanding short positions in the market.
	Reserved4          string    // Reserved field.
	CurrentYrEPS       float64   // The portion of a company's profit allocated to each outstanding share of common stock.
	NextYrEPS          float64   // The total amount of earnings per share a company is estimated to accumulate over the next four quarters of the current fiscal year.
	FiveYrGrowthPct    float64   // Earnings Per Share growth rate over a five year period.
	FiscalYrEnd        float64   // The two digit month that the fiscal year ends for a company.
	Reserved5          string    // Reserved field.
	CompanyName        string    // Company name or contract description
	RootOptionSymbol   []string  // A list of root option symbols, there may be more than one.
	PctHeldByInst      float64   // A percentage of outstanding shares held by banks and institutions.
	Beta               float64   // A coefficient measuring a stock’s relative volatility. It is the covariance of a stock in relation to the rest of the stock market. 30 day historical volatility.
	Leaps              string    // Long term equity anticipation securities.
	CurrentAssets      float64   // The amount of total current assets held by a company as of a specific date in Millions (lastADate).
	CurrentLiabilities float64   // The amount of total current liabilities held by a company as of a specific date in Millions (lastADate).
	BalSheetDate       time.Time // Last date that a company issued their quarterly report. (MM/DD/YYYY).
	LongTermDebt       float64   // The amount of long term debt held by a company as of a specific date in Millions(lastADate).
	ComShrOutstanding  float64   // The amount of common shares outstanding.
	Reserved6          string    // Reserved field.
	SplitFactor1       string    // A float a space, then MM/DD/YYYY
	SplitFactor2       string    // A float a space, then MM/DD/YYYY
	Reserved7          string    // Reserved field.
	Reserved8          string    // Reserved field.
	FormatCode         string    // Display format code, See: Price Format Codes http://www.iqfeed.net/dev/api/docs/PriceFormatCodes.cfm.
	Precision          int       // Number of decimal digits.
	SIC                int       // Federally designed numbering system identifying companies by industry. This 4 digit number corresponds to a specific industry.
	HistVolatility     float64   // 30-trading day volatility that it is calculated using Black-Scholes (https://en.wikipedia.org/wiki/Black%E2%80%93Scholes_model).
	SecurityType       string    // The security type code, See: Security Types (http://www.iqfeed.net/dev/api/docs/SecurityTypes.cfm).
	ListedMarket       string    // The listing market ID, See: Listed Markets
	Fifty2WkHighDate   time.Time // The date of the lowest price of the last 52 weeks. For futures, this is the contract Low Date. (MM/DD/YYYY)
	Fifty2WkLowDate    time.Time // The date of the lowest price of the last 52 weeks. For futures, this is the contract Low Date. (MM/DD/YYYY)
	CalYearHighDate    time.Time // Date at which the High price for the current calendar year occurred. (MM/DD/YYYY)
	CalYearLowDate     time.Time // Date at which the Low price for the current calendar year occurred. (MM/DD/YYYY)
	YrEndClose         float64   // Price of Year End Close. (Equities Only)
	MaturityDate       float64   // Date of maturity for a bond.
	CouponRate         float64   // Interest rate for a bond.
	ExpirationDate     time.Time // IEOptions, Futures, FutureOptions, and SSFutures only
	StrikePrice        float64   // IEOptions only
	NAICS              int       // North American Industry Classification System (http://www.census.gov/eos/www/naics/)
	ExchangeRoot       string    // The root symbol that you can find this symbol listed under at the exchange.
}

// UnMarshall sends the data into the usable struct for consumption by the application.
func (f *FundamentalMsg) UnMarshall(d []byte) {
	fmt.Printf("Unmarshall: %s", string(d))
}

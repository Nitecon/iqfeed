package iqfeed

import (
	"strconv"
	"time"
)

// GetFloatFromStr returns the float representation fo the string.
func GetFloatFromStr(d string) float64 {
	val, _ := strconv.ParseFloat(string(d), 0)
	return val
}

// GetIntFromStr returns a guaranteed int from the string even if it may be 0 on an invalid parse.
func GetIntFromStr(d string) int {
	val, _ := strconv.Atoi(d)
	return val
}

// GetTimeInHMS parses the time field in iqfeed and returns a time object.
func GetTimeInHMS(d string, loc *time.Location) time.Time {
	// We care not about errors here as we require a time field, even if it's in the past or we need to invalidate the entire struct, which may be worse overall.
	// This is unfortunately important as the timefields may not always be set so we at least still return a time object which can be checked for invalid date.
	t, _ := time.ParseInLocation("15:04:05", d, loc)

	return t
}

// GetTimeInHMSmicro parses the time field in iqfeed and returns a time object, this parse includes microsecond accuracy parsing.
func GetTimeInHMSmicro(d string, loc *time.Location) time.Time {
	// We care not about errors here as we require a time field, even if it's in the past or we need to invalidate the entire struct, which may be worse overall.
	// This is unfortunately important as the timefields may not always be set so we at least still return a time object which can be checked for invalid date.
	t, _ := time.ParseInLocation("15:04:05.000", d, loc)

	return t
}

// GetDateMMDDCCYY returns a time object after parsing the MM/DD/CCYY layout in iqfeed.
func GetDateMMDDCCYY(d string, loc *time.Location) time.Time {
	t, _ := time.ParseInLocation("01/02/06", d, loc)

	return t
}

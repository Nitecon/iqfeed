package iqfeed

import (
	"fmt"
	"os"
	"strings"
	"time"
)

// Write performs a write on the channel data which will be picked up by the writer concurrently and written to iqfeed.
func (c *IQC) Write(data string) {
	fmt.Printf("Writing: %s", data)
	c.Conn.Write([]byte(data))
}

// WriteBackup does as the name suggests and write the []byte data directly to a file for re-use later.
func (c *IQC) writeBackup(d []byte) {
	if !c.CreateBackup {
		return
	}
	if _, err := os.Stat(c.BackupFile); os.IsNotExist(err) {
		os.Create(c.BackupFile)
	}
	f, err := os.OpenFile(c.BackupFile, os.O_APPEND|os.O_RDWR, 0644)
	if err != nil {
		fmt.Printf("Could not open file for writing: %s\n", err.Error())
		return
	}
	defer f.Close()
	_, err = f.Write(d)
	if err != nil {
		fmt.Printf("Could not write data to file: %s\n", err.Error())
		return
	}
}

// SetProtocol Changes the current connection's protocol (ex: 5.2).
func (c *IQC) SetProtocol(protocol string) {
	c.Write("S,SET PROTOCOL," + protocol + "\r\n")
}

// SetClientName does as the name implies and sets the client message which will also be available in stats.
func (c *IQC) SetClientName(name string) {
	c.Write("S,SET CLIENT NAME," + name + "\r\n")
}

// WatchSymbol will issue a command to start watching a symbol, this will return a fundamental and update message with the quotes.
func (c *IQC) WatchSymbol(symbol string) {
	c.Write("w" + symbol + "\r\n")
}

// WatchOptionSymbol tracks a new symbol based on contract date (for option chains), contractDate indicates the date for the option contract and isCall indicates whether it is a call / put contract.
func (c *IQC) WatchOptionSymbol(symbol string, value float64, contractDate time.Time, isCall bool) {
	// Final format should be something like: MSFT1220J30.5
	ydm := contractDate.Format("0602")
	if isCall {
		ydm = ydm + c.getCallChar(contractDate)
	} else {
		ydm = ydm + c.getPutChar(contractDate)
	}
	tSym := fmt.Sprintf("%s%s%f", symbol, ydm, value)
	c.WatchSymbol(tSym)
}

// TradeOnlyWatch Begins a trades only watch on a symbol for Level 1 updates.
func (c *IQC) TradeOnlyWatch(symbol string) {
	c.Write("t" + symbol + "\r\n")
}

// UnwatchSymbol Terminates Level 1 updates for the symbol specified.
func (c *IQC) UnwatchSymbol(symbol string) {
	c.Write("r" + symbol + "\r\n")
}

// ForceRefresh Forces a refresh from the server for the symbol specified.
func (c *IQC) ForceRefresh(symbol string) {
	c.Write("f" + symbol + "\r\n")
}

// RequestTime Requests a Time Stamp message be sent.
func (c *IQC) RequestTime() {
	c.Write("T\r\n")
}

// DisableTSUpdates Disables once per second timestamps
func (c *IQC) DisableTSUpdates() {
	c.Write("S,TIMESTAMPSOFF\r\n")
}

// EnableTSUpdates Timestamps default to on, but in the event you have stopped them manually, this will restart them into the stream.
func (c *IQC) EnableTSUpdates() {
	c.Write("S,TIMESTAMPSON\r\n")
}

// RegionWatch Begins watching a symbol for Level 1 Regional updates.
func (c *IQC) RegionWatch(symbol string) {
	c.Write("S,REGON," + symbol + "\r\n")
}

// RegionWatchOff Stops watching a symbol for Level 1 Regional updates.
func (c *IQC) RegionWatchOff(symbol string) {
	c.Write("S,REGOFF," + symbol + "\r\n")
}

// NewsOn Turns on streaming news headlines.
func (c *IQC) NewsOn() {
	c.Write("S,NEWSON\r\n")
}

// NewsOff Turns off streaming news headlines.
func (c *IQC) NewsOff() {
	c.Write("S,NEWSOFF\r\n")
}

// RequestStats Request a S,STATS message to give you information about the feed status.
func (c *IQC) RequestStats() {
	c.Write("S,REQUEST STATS\r\n")
}

// ReqFundamentalFNames Request a list of all available fundamental message field names.
func (c *IQC) ReqFundamentalFNames() {
	c.Write("S,REQUEST STATS\r\n")
}

// ReqAllUpdateFNames Request a list of all available summary/update message field names for the currently set IQFeed protocol.
func (c *IQC) ReqAllUpdateFNames() {
	c.Write("S,REQUEST ALL UPDATE FIELDNAMES\r\n")
}

// ReqCurrentUpdateFNames Request a list of field names in the current fieldset for this connection.\
// Result: You will receive a S,CURRENT UPDATE FIELDNAMES,[FIELD 1 NAME],[FIELD 2 NAME],...[FIELD N NAME],<LF> message that contains currently selected summary/update fields.
func (c *IQC) ReqCurrentUpdateFNames() {
	c.Write("S,REQUEST CURRENT UPDATE FIELDNAMES\r\n")
}

// SelectUpdateFields Change your fieldset for this connection. This fieldset applies to all summary and update messages you receive on this connection. (Comma seperated list of field names).
func (c *IQC) SelectUpdateFields(fields ...string) {
	c.Write("S,SELECT UPDATE FIELDS," + strings.Join(fields, ",") + "\r\n")
}

// SetLogLevels Change the logging levels for IQFeed. Level Docs: http://www.iqfeed.net/dev/api/docs/IQConnectLogging.cfm.
func (c *IQC) SetLogLevels(levels ...string) {
	c.Write("S,SET LOG LEVELS," + strings.Join(levels, ",") + "\r\n")
}

// RequestWatches Request a list of all symbols currently watched on this connection.
func (c *IQC) RequestWatches() {
	c.Write("S,REQUEST WATCHES\r\n")
}

// UnwatchAllSymbols Unwatch all currently watched symbols.
func (c *IQC) UnwatchAllSymbols() {
	c.Write("S,UNWATCH ALL\r\n")
}

// Connect Tells IQFeed to initiate a connection to the Level 1 server. This happens automatically upon launching the feed unless the ProductID and/or Product version have not been set. This message is ignored if the feed is already connected.
func (c *IQC) Connect() {
	c.Write("S,CONNECT\r\n")
}

// Disconnect Tells IQFeed to disconnect from the Level 1 server. This happens automatically as soon as the last client connection to IQConnect is terminated and the ClientsConnected value in the S,STATS message returns to zero (after having incremented above zero). This message is ignored if the feed is already disconnected.
func (c *IQC) Disconnect() {
	c.Write("S,DISCONNECT\r\n")
}

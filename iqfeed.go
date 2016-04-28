package iqfeed

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"
	"sync"
)

var (
	iqcLock sync.RWMutex
	iqc     *IQClient
)

// IQClient provides the main struct for the the IQ Client interface into what IQFeed will be sending us.
type IQClient struct {
	Data            chan []byte
	GenBackup       bool
	BackupFile      string
	Conn            net.Conn
	Quit            chan bool
	IPs             []string
	Fields          []string
	TrackingSymbols map[string]int
}

// WriteBackup does as the name suggests and write the []byte data directly to a file for re-use later.
func (c *IQClient) WriteBackup(d []byte) {
	if !c.GenBackup {
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

// ProcessSystemMsg handles system messages, field definitions are available here: http://www.iqfeed.net/dev/api/docs/Level1SystemMessage.cfm.
func (c *IQClient) ProcessSysMsg(d []byte) {

}

// ProcessSumMsg handles summary messages, field definitions are available here: http://www.iqfeed.net/dev/api/docs/Level1UpdateSummaryMessage.cfm.
func (c *IQClient) ProcessSumMsg(d []byte) {

}

// ProcessUpdMsg handles update messages, field definitions are available here: http://www.iqfeed.net/dev/api/docs/Level1UpdateSummaryMessage.cfm.
func (c *IQClient) ProcessUpdMsg(d []byte) {

}

// ProcessTimeMsg handles timestamp updates, field definitions are available here: http://www.iqfeed.net/dev/api/docs/TimeMessageFormat.cfm.
func (c *IQClient) ProcessTimeMsg(d []byte) {

}

// ProcessRegUpdMsg handles regional updates field definitions are available here: http://www.iqfeed.net/dev/api/docs/RegionalMessageFormat.cfm.
func (c *IQClient) ProcessRegUpdMsg(d []byte) {

}

// ProcessFndMsg handles fundamental messages, field descriptions are available here: http://www.iqfeed.net/dev/api/docs/Level1FundamentalMessage.cfm.
func (c *IQClient) ProcessFndMsg(d []byte) {

}

// ProcessNewsMsg handles summary messages, field definitions are available here: http://www.iqfeed.net/dev/api/docs/StreamingNewsMessageFormat.cfm.
func (c *IQClient) ProcessNewsMsg(d []byte) {

}

// Process404Msg handles messages indicating that a symbol was not found.
func (c *IQClient) Process404Msg(d []byte) {

}

// ProcessUpdMsg handles error messages in the form of error text.
func (c *IQClient) ProcessErrorMsg(d []byte) {

}

// ProcessReceiver is one of the main reciever functions that interprets data received by IQFeed and processes it in sub functions.
func (c *IQClient) ProcessReceiver(d []byte) {
	data := d[2:]
	switch d[0] {
	case 0x53: // Start letter is S, indicating System message (Unicode representation in integer value).
		c.ProcessSysMsg(data)
	case 0x50: // Start letter is P, indicating a summary message.
		c.ProcessSumMsg(data)
	case 0x51: // Start letter is Q, indicating an update message.
		c.ProcessUpdMsg(data)
	case 0x54: // Start letter is T, indicating Time message.
		c.ProcessTimeMsg(data)
	case 0x52: // Start letter is R, indicating regional update message
		c.ProcessRegUpdMsg(data)
	case 0x46: // Start letter is F, indicating a fundamental message
		c.ProcessFndMsg(data)
	case 0x4e: // Start letter is N, indicating a news message
		c.ProcessNewsMsg(data)
	case 0x6E: // Start letter is n, indicating Symbol not found message
		c.Process404Msg(data)
	case 0x45: // Start letter is E, error message
		c.ProcessErrorMsg(data)
	}

}

// Read function does as expected and reads data from the network stream.
func (c *IQClient) Read() {
	r := bufio.NewReader(c.Conn)

	for {

		select {
		case <-c.Quit:
			log.Println("Client quitting")
			c.Conn.Close()
			break

		default:

			line, isPrefix, err := r.ReadLine()
			for err == nil && !isPrefix {
				if c.GenBackup {
					bld := fmt.Sprintf("%s\r\n", string(line))
					c.WriteBackup([]byte(bld))
				}
				c.ProcessReceiver(line)
				line, isPrefix, err = r.ReadLine()
			}
			if isPrefix {
				log.Println("buffer size to small")
				//return Do not return and break the loop
			}
			if err != io.EOF {
				log.Println(err)
				//return Do not return and break the loop
			}
		}
	}

}

// Write performs a write on the channel data which will be picked up by the writer concurrently and written to iqfeed.
func (c *IQClient) Write(data string) {
	c.Data <- []byte(data)
}

// Writer provides a method to do writes directly to the socket in order to *talk* to iqfeed.
func (c *IQClient) Writer() {
	for {
		select {
		case data := <-c.Data:
			_, err := c.Conn.Write(data)
			if err != nil {
				log.Fatal("Error writing to socket: ", err.Error())
			}
			fmt.Printf("WRITE[%s]\n", strings.TrimSpace(string(data)))
		case <-c.Quit:
			log.Println("Client quitting")
			c.Conn.Close()
			break
		}
	}
}

// GetIQClient returns our IQClient for use in your application.
func GetIQClient() *IQClient {
	iqcLock.RLock()
	defer iqcLock.RUnlock()
	return iqc
}

// Start function willl start the concurrent functions to read and write data to the and from the network stream.
func (c *IQClient) Start() {
	go c.Writer()
	go c.Read()

}

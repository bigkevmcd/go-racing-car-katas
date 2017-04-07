package telemetry

import (
	"errors"
	"math/rand"
)

const (
	DIAGNOSTIC_MESSAGE = "AT#UD"
)

type TelemetryClient struct {
	onlineStatus              bool
	diagnosticMessageResult   string
	connectionEventsSimulator *rand.Rand
}

func NewTelemetryClient() *TelemetryClient {
	return &TelemetryClient{connectionEventsSimulator: rand.New(rand.NewSource(42))}
}

func (c *TelemetryClient) OnlineStatus() bool {
	return c.onlineStatus
}

func (c *TelemetryClient) Connect(s string) error {
	if s == "" {
		return errors.New("connection cannot be empty")
	}
	// simulate the operation on a real modem
	c.onlineStatus = c.connectionEventsSimulator.Intn(10) < 8
	return nil
}

func (c *TelemetryClient) Disconnect() {
	c.onlineStatus = false
}

func (c *TelemetryClient) Send(msg string) error {
	if msg == "" {
		return errors.New("msg cannot be empty")
	}
	if msg == DIAGNOSTIC_MESSAGE {
		c.diagnosticMessageResult =
			"LAST TX rate................ 100 MBPS\r\n" +
				"HIGHEST TX rate............. 100 MBPS\r\n" +
				"LAST RX rate................ 100 MBPS\r\n" +
				"HIGHEST RX rate............. 100 MBPS\r\n" +
				"BIT RATE.................... 100000000\r\n" +
				"WORD LEN.................... 16\r\n" +
				"WORD/FRAME.................. 511\r\n" +
				"BITS/FRAME.................. 8192\r\n" +
				"MODULATION TYPE............. PCM/FM\r\n" +
				"TX Digital Los.............. 0.75\r\n" +
				"RX Digital Los.............. 0.10\r\n" +
				"BEP Test.................... -5\r\n" +
				"Local Rtrn Count............ 00\r\n" +
				"Remote Rtrn Count........... 00"
	}
	return nil
	// here should go the real Send operation (not needed for this exercise)
}

func (c *TelemetryClient) Receive() string {
	message := ""
	if c.diagnosticMessageResult == "" {
		// simulate a received message (just for illustration - not needed for this exercise)
		messageLength := c.connectionEventsSimulator.Intn(50) + 60
		for i := messageLength; i >= 0; i-- {
			message += string(rune(c.connectionEventsSimulator.Intn(40) + 86))
		}
	} else {
		message = c.diagnosticMessageResult
		c.diagnosticMessageResult = ""
	}
	return message
}

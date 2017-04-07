package telemetry

import "errors"

const (
	DiagnosticChannelConnectionString = "*111#"
)

type TelemetryDiagnosticControls struct {
	telemetryClient *TelemetryClient
	diagnosticInfo  string
}

func NewTelemetryDiagnosticControls() *TelemetryDiagnosticControls {
	return &TelemetryDiagnosticControls{
		telemetryClient: NewTelemetryClient(),
	}
}

func (c *TelemetryDiagnosticControls) DiagnosticInfo() string {
	return c.diagnosticInfo
}

func (c *TelemetryDiagnosticControls) SetDiagnosticInfo(s string) {
	c.diagnosticInfo = s
}

func (c *TelemetryDiagnosticControls) CheckTransmission() error {
	c.diagnosticInfo = ""
	c.telemetryClient.Disconnect()

	retryLeft := 3
	for c.telemetryClient.OnlineStatus() == false && retryLeft > 0 {
		c.telemetryClient.Connect(DiagnosticChannelConnectionString)
		retryLeft--
	}

	if c.telemetryClient.OnlineStatus() == false {
		return errors.New("unable to connect")
	}
	c.telemetryClient.Send(DIAGNOSTIC_MESSAGE)
	c.diagnosticInfo = c.telemetryClient.Receive()
	return nil
}

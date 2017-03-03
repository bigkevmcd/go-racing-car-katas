package tirepressure

import (
	"math/rand"
	"time"
)

const (
	LowPressureThreshold  = 17.0
	HighPressureThreshold = 21.0
	OFFSET                = 16.0
)

type Alarm struct {
	sensor  Sensor
	AlarmOn bool
}

func NewAlarm() *Alarm {
	return &Alarm{sensor: Sensor{}}
}

func (c *Alarm) Check() {
	psiPressureValue := c.sensor.PopNextPressurePsiValue()

	if psiPressureValue < LowPressureThreshold || HighPressureThreshold < psiPressureValue {
		c.AlarmOn = true
	}
}

type Sensor struct {
}

func (s Sensor) PopNextPressurePsiValue() float64 {
	return s.samplePressure() + OFFSET
}

func (s Sensor) samplePressure() float64 {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return 6 * (r.Float64() * r.Float64())
}

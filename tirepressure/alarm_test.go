package tirepressure

import "testing"

func TestAlarm(t *testing.T) {
	alarm := NewAlarm()

	if alarm.AlarmOn {
		t.Fatal("AlarmOn is true, wanted false")
	}
}

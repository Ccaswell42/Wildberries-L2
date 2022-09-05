package main

import (
	"math"
	"testing"
	"time"
)

func TestGetCurrentTime(t *testing.T) {
	ntp, err := GetCurrentTime()
	if err != nil {
		t.Fatal(err)
	}
	local := time.Now()

	if math.Abs(float64(ntp.Second()-local.Second())) > 1 {
		t.Fatal("разница во времени больше 1 секнуды")
	}
}

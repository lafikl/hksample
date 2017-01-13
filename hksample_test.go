package hksample

import (
	"testing"

	uuid "github.com/satori/go.uuid"
)

func TestSample(t *testing.T) {
	hk := NewHKSample(.2)
	var wantedPercent float64 = 20
	var errorRate float64 = 5 // in percent
	var samples float64 = 321
	var tracked float64
	var i float64
	for i < samples {
		id := uuid.NewV4()
		if hk.Sample(id.Bytes()) {
			tracked++
		}
		i++
	}
	percentTracked := float64(tracked/samples) * 100
	msg := "Tracked (%.2f, %.2f%%) not within the the bounds (%.2f%%, %.2f%%)"

	if !(percentTracked >= wantedPercent-errorRate && percentTracked <= percentTracked+errorRate) {
		t.Fatalf(msg, tracked, percentTracked,
			wantedPercent, percentTracked+errorRate)
	}

	if hk.Sampled() != uint64(tracked) {
		t.Fatalf("Wrong samples count. Got %d, expected %d", hk.Sampled(), tracked)
	}
}

func TestSampleString(t *testing.T) {
	hk := NewHKSample(.2)
	var wantedPercent float64 = 20
	var errorRate float64 = 5 // in percent
	var samples float64 = 321
	var tracked float64
	var i float64
	for i < samples {
		id := uuid.NewV4()
		if hk.SampleString(id.String()) {
			tracked++
		}
		i++
	}
	percentTracked := float64(tracked/samples) * 100
	msg := "Tracked (%.2f, %.2f%%) not within the the bounds (%.2f%%, %.2f%%)"

	if !(percentTracked >= wantedPercent-errorRate && percentTracked <= percentTracked+errorRate) {
		t.Fatalf(msg, tracked, percentTracked,
			wantedPercent, percentTracked+errorRate)
	}

}

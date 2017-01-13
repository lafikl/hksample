// Package HKSample does sampling by hashing the key given to it,
// then it sees if the key is within the sampling buckets.
//
// This way we guarantee that if a different system uses HKSample,
// it'll always pick the same events the previous systems picked.
//
//
// ## Key Features:
//
// - No extra metadata to carry in every message you send.
// - No need for communications between systems to sample/pick the same events/sessions.
// - Tested in real world scenarios, against millions of events.
//
//
// ## Why i made this library?
//
// I needed it to build a tracing system that samples a complete session/flow,
// to gain visiblity on a data pipeline that receives events/messages
// with a single key/id identifying a complete session.
//
//
// ## Installation
//
//      go get github.com/lafikl/hksample
//
//
// ## Example
//
//      hk := NewHKSample(.20)
//      key := []byte(`5efe2993-6b3d-48e2-ac94-1afd543d9190`)
//      if hk.Sample(key) {
//          log.Printf("Sampled: %s", key)
//      }
//
//
// ## Documentation
// [GoDoc](https://godoc.org/github.com/lafikl/hksample)
//
package hksample

import (
	"hash/fnv"
	"math"
)

const (
	// prime number of slots
	slots = 23
)

type HKSample struct {
	threshold uint64
	sampled   uint64
}

// NewHKSample returns a new instance HKSample
// percent should be given as a fraction, for example:
//
//      NewHKSample(.20) // 20%
func NewHKSample(percent float64) *HKSample {
	t := math.Floor(slots * percent)
	return &HKSample{threshold: uint64(t)}
}

// Sample returns true if the supplied key falls into the buckets within
// the given threshold/percentage
func (hk *HKSample) Sample(key []byte) bool {
	f := fnv.New64()
	f.Write(key)
	h := f.Sum64()
	ok := (h % slots) <= hk.threshold
	if ok {
		hk.sampled++
	}
	return ok
}

// SampleString is the same as Sample but it takes a string
func (hk *HKSample) SampleString(key string) bool {
	return hk.Sample([]byte(key))
}

// Sampled returns the number of sampled events
func (hk *HKSample) Sampled() uint64 {
	return hk.sampled
}

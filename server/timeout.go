package main

// This file is a rewrite of github.com/cenkalti/backoff/blob/master/exponential.go which the aim 
// to simplify the logic and the behavior of this package.

import (
	"math/rand"
	"time"
)

type timeout struct {
	initial time.Duration
	current time.Duration
	max time.Duration
	factor float64
	multiplier float64
}

const (
	defaultFactor = 0.5
	defaultMultiplier = 1.5
)

func (t *timeout) next() time.Duration {
	defer t.update()
	return t.random()
}

func (t *timeout) reset() {
	t.current = t.initial
}

func (t *timeout) random() time.Duration {

	delta := t.factor * float64(t.current)
	min := float64(t.current) - delta
	max := float64(t.current) + delta
	random := rand.Float64()

	// Get a random value from the range [min, max].
	// The formula used below has a +1 because if the min is 1 and the max is 3 then
	// we want a 33% chance for selecting either 1, 2 or 3.
	return time.Duration(min + (random * (max - min + 1)))
}

func (t *timeout) update() {

	// Check for overflow, if overflow is detected set the current interval to the max interval.
	if float64(t.current) >= float64(t.max)/t.multiplier {
		t.current = t.max
	} else {
		t.current = time.Duration(float64(t.current) * t.multiplier)
	}
}

func initTimeout(t time.Duration) *timeout {
	return &timeout {
		initial: t,
		current: t,
		factor: defaultFactor,
		multiplier: defaultMultiplier,
		max: time.Duration(4 * time.Minute),
	}
}

func main() {

	t := initTimeout(time.Duration(250) * time.Millisecond)

	for i := uint16(0); i < 2000; i++ {

		println(i, t.next().String())
	}

}
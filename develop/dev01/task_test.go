package main

import (
	"testing"
	"time"
)

func TestTime(t *testing.T) {
	got := getTime()
	want := time.Now()
	difference := got.Sub(want)
	if difference < 0 {
		difference = -difference
	}
	if difference > time.Duration(time.Second * 3) {
		t.Errorf("\ngetTime():  %q\ntime.Now(): %q\nDifference: %q", got, want, difference)
	}
}

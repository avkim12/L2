package main

import (
	"testing"
	"time"
)

func TestTime(t *testing.T) {
	got := GetTime()
	want := time.Now()
	difference := got.Sub(want)
	if difference < 0 {
		difference = -difference
	}
	if difference > 0 {
		t.Errorf("\ngetTime():  %q\ntime.Now(): %q\nDifference: %q", got, want, difference)
	}
}
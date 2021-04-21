package internal

import (
	"testing"
	"time"
)

func TestGetDailyQuote(t *testing.T) {
	data := map[time.Weekday]string {
		time.Monday: "It's Monday",
		time.Tuesday: "It's Tuesday",
		time.Wednesday: "It Is Wednesday My Dudes",
		time.Thursday: "It's Thursday",
		time.Friday: "It's the start of the weekend",
		time.Saturday: "It's lördag today",
		time.Sunday: "It's Sunday",
		time.Weekday(10): "I don't know what it is today",
	}

	for d, expected := range data {
		if got := GetDailyQuote(d); got != expected {
			t.Errorf("Expected `%s`, got `%s`", expected, got)
		}
	}
}

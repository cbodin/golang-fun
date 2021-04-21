package internal

import (
	"fmt"
	"time"
)

func GetDailyQuote(d time.Weekday) string {
	switch d {
	case time.Monday, time.Tuesday, time.Thursday, time.Sunday:
		return fmt.Sprintf("It's %s", d)
	case time.Wednesday:
		return "It Is Wednesday My Dudes"
	case time.Friday:
		return "It's the start of the weekend"
	case time.Saturday:
		return "It's lördag today"
	}

	return "I don't know what it is today"
}

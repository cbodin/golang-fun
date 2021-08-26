package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/cbodin/golang-fun/internal"
)

var weekdays = map[string]time.Weekday{
	"/monday":    time.Monday,
	"/tuesday":   time.Tuesday,
	"/wednesday": time.Wednesday,
	"/thursday":  time.Thursday,
	"/friday":    time.Friday,
	"/saturday":  time.Saturday,
	"/sunday":    time.Sunday,
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var d time.Weekday

		if r.URL.Path == "/" {
			d = time.Now().Weekday()
		} else if day, ok := weekdays[r.URL.Path]; ok {
			d = day
		} else {
			http.NotFound(w, r)
			return
		}

		_, _ = fmt.Fprint(w, internal.GetDailyQuote(d))
	})

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

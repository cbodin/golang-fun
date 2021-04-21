package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/cbodin/golang-fun/internal"
)

func main() {
	http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
		d := time.Now().Weekday()
		_, _ = fmt.Fprint(w, internal.GetDailyQuote(d))
	})

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

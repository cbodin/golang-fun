package main

import (
	"fmt"
	"time"

	"github.com/cbodin/golang-fun/internal"
)

func main() {
	d := time.Now().Weekday()
	fmt.Println(internal.GetDailyQuote(d))
}

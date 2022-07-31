package main

import (
	"fmt"
	"os"
	"time"

	"github.com/beevik/ntp"
)

func main() {
	ttime, err := ntp.Time("0.beevik-ntp.pool.ntp.org")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %d", err)
		os.Exit(1)
	}
	fmt.Printf("The time is: %s\n", ttime.Format(time.RFC1123))
}

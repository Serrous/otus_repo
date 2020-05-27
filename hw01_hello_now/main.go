package main

import (
	"fmt"
	"log"
	"time"

	"github.com/beevik/ntp"
)

const ntpsrv = "0.pool.ntp.org"

func main() {
	localTime := time.Now()
	ntpTime, err := ntp.Time(ntpsrv)
	if err != nil {
		log.Fatal(err)
	}
	//The canonical way to strip a monotonic clock reading is to use t = t.Round(0).
	fmt.Printf("current time: %s\nexact time: %s\n", localTime.Round(0), ntpTime.Round(0))
}

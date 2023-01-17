package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	totalTime := flag.String("time", "", "The length of the timer | e.g. 3h, 25m, 1h30m, 45s")

	flag.Parse()

	if *totalTime == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	v, _ := time.ParseDuration(*totalTime)

	formattedTime := time.Now().Add(v).Format("2006-01-02T15:04:05") + "+00:00"

	parsedTime, _ := time.Parse(time.RFC3339, formattedTime)

	//fmt.Printf("Staring timer for %v \n", v)

	timeLeft := getTimeLeft(parsedTime)

	fmt.Print("\033[H\033[2J")
	fmt.Printf("%02d : %02d : %02d : %02d\n", timeLeft.d, timeLeft.h, timeLeft.m, timeLeft.s)
	for range time.Tick(1 * time.Second) {
		timeLeft := getTimeLeft(parsedTime)

		fmt.Print("\033[H\033[2J")
		fmt.Printf("%02d : %02d : %02d : %02d\n", timeLeft.d, timeLeft.h, timeLeft.m, timeLeft.s)
		if timeLeft.t <= 0 {
			fmt.Println(strings.Repeat("-", 30))
			fmt.Printf("Time ended: %v\n", v)
			fmt.Println(strings.Repeat("-", 30))
			os.Exit(1)
		}
	}
}

type countdown struct {
	t int
	d int
	h int
	m int
	s int
}

func getTimeLeft(t time.Time) countdown {
	timeNow := time.Now()
	left := t.Sub(timeNow)

	total := int(left.Seconds())
	days := int(total / (60 * 60 * 24))
	hours := int(total / (60 * 60) % 24)
	minutes := int(total/60) % 60
	seconds := int(total % 60)

	return countdown{
		t: total,
		d: days,
		h: hours,
		m: minutes,
		s: seconds,
	}

}

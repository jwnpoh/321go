package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

const usage = `
A simple countdown timer. 
Usage: 
321go hh mm ss
`

type countdown struct {
	hrs  string
	mins string
	secs string
    hh int
    mm int
    ss int
}

func (c *countdown) parseDuration() (time.Duration, error) {
    if c.hrs == "" {
        c.hrs = "0"
    }

    if c.mins == "" {
        c.mins = "0"
    }

    d := c.hrs + "h" + c.mins + "m" + c.secs + "s"

    duration, err := time.ParseDuration(d)
    if err != nil {
        return duration, fmt.Errorf("unable to parse duration, %v", err)
    }

    return duration, nil
}

func (c *countdown) makeInt() error {
    if c.hrs != "" {
        h, err := strconv.Atoi(c.hrs)
        if err != nil {
            return fmt.Errorf("unable to convert hrs to int")
        }
        c.hh = h
    }

    if c.mins != "" {
        m, err := strconv.Atoi(c.mins)
        if err != nil {
            return fmt.Errorf("unable to convert mins to int")
        }
        c.mm = m
    }

    s, err := strconv.Atoi(c.secs)
    if err != nil {
        return fmt.Errorf("unable to convert secs to int")
    }
    c.ss = s

    return nil
}

func main() {
	var c countdown
	args := os.Args[1:]

    if err := parseInput(args, &c); err != nil {
        fmt.Println(err)
        return
    }

    duration, err := c.parseDuration()
    if err != nil {
        fmt.Println(err)
        return
    }

    tickerTime(duration)
}

func parseInput(args []string, c *countdown) error {
	if len(args) < 1 || len(args) > 3 {
		log.Fatal(usage)
	}

	switch len(args) {
	case 3:
		c.hrs = args[0]
		c.mins = args[1]
		c.secs = args[2]
	case 2:
		c.mins = args[0]
		c.secs = args[1]
	case 1:
		c.secs = args[0]
	}

    if err := c.makeInt(); err != nil{
        return fmt.Errorf("unable to make ints out of input")
    }

    if c.mm > 59 || c.ss > 59 {
        return fmt.Errorf("are you sure you have entered the correct time? The value for minutes and seconds should not be more than 59")
    }
    return nil
}

func tickerTime(d time.Duration) {
    start := time.Now()
    end := start.Add(d)

    ticker := time.NewTicker(time.Second)
    defer ticker.Stop()
    done := make(chan bool)

    go func() {
        time.Sleep(d + time.Second)
        done <- true
    }()

    for {
        select {
        case <- done:
        endMsg()
        return
        case t := <- ticker.C:
        remaining := end.Sub(t)
        remaining += time.Second * 2
        printClock(remaining)
        }
    }
}

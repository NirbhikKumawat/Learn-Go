package main

import (
	"fmt"
	"time"
)

func main() {
	p := fmt.Print
	now := time.Now()
	p(now)
	fmt.Printf("\n")
	then := time.Date(2009, 11, 17, 20, 34, 58, 65138, time.UTC)
	p(then)
	fmt.Printf("\n")

	p(then.Year())
	fmt.Printf("\n")
	p(then.Month())
	fmt.Printf("\n")
	p(then.Day())
	fmt.Printf("\n")
	p(then.Hour())
	fmt.Printf("\n")
	p(then.Minute())
	fmt.Printf("\n")
	p(then.Second())
	fmt.Printf("\n")
	p(then.Nanosecond())
	fmt.Printf("\n")

	p(then.Location())
	fmt.Printf("\n")

	p(then.Weekday())
	fmt.Printf("\n")

	p(then.Before(now))
	fmt.Printf("\n")
	p(then.After(now))
	fmt.Printf("\n")
	p(then.Equal(now))
	fmt.Printf("\n")

	diff := now.Sub(then)
	p(diff)
	fmt.Printf("\n")
	p(diff.Hours())
	fmt.Printf("\n")
	p(diff.Minutes())
	fmt.Printf("\n")
	p(diff.Seconds())
	fmt.Printf("\n")
	p(diff.Nanoseconds())
	fmt.Printf("\n")

	p(then.Add(diff))
	fmt.Printf("\n")
	p(then.Add(-diff))
	fmt.Printf("\n")

}

package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	if len(os.Args[1:]) != 1 {
		fmt.Println("Must provide only one arg")
		os.Exit(1) // notify the parent
	}

	// get date arg from command line
	args := os.Args[1:]
	date := args[0]

	// strip brackets
	dateString := date[1 : len(date)-1]

	// this was fun. To specify custom format, you need to choose
	// the exact consts found here `https://golang.org/pkg/time/`, otherwise
	// the time package wont be able to parse your date!
	form := "2006-01-02 15:04:05 -0700"
	t, e := time.Parse(form, dateString)

	if e != nil {
		fmt.Println("Doh - unable to parse date")
		os.Exit(1) // notify the parent
	}

	// calc how long it took
	duration := time.Since(t)
	days := 0
	hours := duration.Hours()
	minutes := duration.Minutes()

	// display minutes
	if hours < 1 {
		fmt.Printf("\033[34m==>\033[0m It took %.0f minutes to complete this task\n", minutes)
		os.Exit(1)
	}

	// display hours
	if hours < 24 {
		fmt.Printf("\033[34m==>\033[0m It took %.1f hours to complete this task\n", hours)
		os.Exit(1)
	}

	// display days
	days = int(hours / 24)
	fmt.Printf("\033[34m==>\033[0m It took %d days to complete this task\n", days)
}

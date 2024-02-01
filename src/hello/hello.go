package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	if len(os.Args) > 3 {
		fmt.Println("You have provided too many arguments.")
		os.Exit(1)
	}

	if len(os.Args) <= 1 {
		fmt.Println("Please provide a name.")
		os.Exit(1)
	}

	name := os.Args[1]
	fmt.Printf("Hello, %s!\n", name)

	if len(os.Args) <= 2 {
		fmt.Println("Please provide your birtday as well. (dd-mm-yyyy)")
		os.Exit(1)
	}

	birthday := os.Args[2]
	fmt.Printf("Your birthday is %s.\n", birthday)
	calculateAge(birthday)
}

func calculateAge(birthday string) {
	parsedDate, err := time.Parse("02-01-2006", birthday)
	now := time.Now()
	today := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.UTC)

	if err != nil {
		fmt.Println("Please provide your birthday in the correct format. (dd-mm-yyyy)")
		os.Exit(1)
	} else if parsedDate.Equal(today) || parsedDate.After(today) {
		fmt.Println("We are not able to calculate your age yet. Please provide a valid date in the past.")
		os.Exit(1)
	} else {
		fmt.Println("Your birthday is in the correct format.")
	}

	duration := today.Sub(parsedDate)
	years := duration / (365 * 24 * time.Hour)
	fmt.Printf("You are %d years old.\n", years)

	months := duration / (30 * 24 * time.Hour)
	fmt.Printf("You are %d months old.\n", months)

	days := duration / (24 * time.Hour)
	fmt.Printf("You are %d days old.\n", days)

	hours := duration / time.Hour
	fmt.Printf("You are %d hours old.\n", hours)

	minutes := duration / time.Minute
	fmt.Printf("You are %d minutes old.\n", minutes)

	seconds := duration / time.Second
	fmt.Printf("You are %d seconds old.\n", seconds)

}

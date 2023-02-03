package main

import "fmt"

const (
	Day = iota
	Month
	Year
)

func main() {
	var expected, actual [3]int
	fmt.Scanf(
		"%d %d %d\n%d %d %d",
		&expected[Day],
		&expected[Month],
		&expected[Year],
		&actual[Day],
		&actual[Month],
		&actual[Year],
	)

	fine := 0
	switch {
	case actual[Year] > expected[Year]:
		// The book was returned after the expected year
		fine = 12000
	case actual[Year] == expected[Year] && actual[Month] > expected[Month]:
		// The book was returned after the expected month in the same year
		fine = 500 * (actual[Month] - expected[Month])
	case actual[Year] == expected[Year] && actual[Month] == expected[Month] && actual[Day] > expected[Day]:
		// The book was returned after the expected day in the same month and year
		fine = 15 * (actual[Day] - expected[Day])
	}

	fmt.Println(fine)
}

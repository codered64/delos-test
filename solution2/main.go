package main

import "fmt"

func main() {
	var student, candies, firstStudent int
	fmt.Scan(&student, &candies, &firstStudent)

	// The solution uses formula: (firstStudent + candies - 2) % student + 1
	// The (firstStudent + candies - 2) is to get the zero based index of the last candy.
	// The "% student" is to wrap around to the first student if the index is greater than or equal to number of students
	// by using modulo operator (the remainder of the division).
	// The "+ 1" is used to convert the zero based index back to the one based index
	fmt.Println((firstStudent+candies-2)%student + 1)
}

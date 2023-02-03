package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	arr := make([]int, n)
	for i := range arr {
		fmt.Scan(&arr[i])
	}

	// Loop from the second element to the second to last element (as 'i')
	for i := 1; i < n-1; i++ {
		// Calculate sum of the elements to the left of current element 'i'
		leftSum := 0
		for j := 0; j < i; j++ {
			leftSum += arr[j]
		}

		// Calculate sum of the elements to the right of current element 'i'
		rightSum := 0
		for j := i + 1; j < n; j++ {
			rightSum += arr[j]
		}

		// Check whether the two sums is equal
		if leftSum == rightSum {
			// Found
			fmt.Println("YES")
			return
		}
	}

	// Not found
	fmt.Println("NO")
}

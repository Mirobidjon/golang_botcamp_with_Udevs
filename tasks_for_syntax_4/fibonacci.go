package main

import "fmt"

func main() {
	//fibonacci  numbers
	fmt.Println("n ni kiriting: ")
	var n int
	fmt.Scan(&n)
	result := fib(n)
	fmt.Println("N - fibonacci soni : ", result)

	fibArr := []int{}
	fibArr = getFibArr(fibArr, n)
	fmt.Println(fibArr)
}

func fib(n int) int {
	if n < 1 {
		return 0
	} else if n <= 2 {
		return 1
	}
	return fib(n-1) + fib(n-2)
}

func getFibArr(arr []int, n int) []int {
	arr = append(arr, 0)
	if n < 1 {
		return arr
	}
	arr = append(arr, 1)
	for i := 2; i <= n; i++ {
		arr = append(arr, arr[i-1]+arr[i-2])
	}
	return arr
}

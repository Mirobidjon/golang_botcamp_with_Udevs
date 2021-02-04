package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	odd := (1 + n) / 2 * ((n + 1) / 2)
	even := (2 + n) / 2 * (n / 2)
	//1 dan n gacha toq va juft sonlar yig'indisini hisoblash
	fmt.Printf("Juft sonlar yig'indisi : %d\n", even)
	fmt.Printf("Toq sonlar yig'indisi : %d\n", odd)

	//sikl orqali hisoblash
	var odd1, even1 int = 0, 0
	for i := 1; i <= n; i++ {
		if i%2 == 0 {
			even1 += i
		} else {
			odd1 += i
		}
	}
	fmt.Println("Sikl orqali hisoblash: ")
	fmt.Printf("Juft sonlar yig'indisi : %d\n", even1)
	fmt.Printf("Toq sonlar yig'indisi : %d\n", odd1)
}

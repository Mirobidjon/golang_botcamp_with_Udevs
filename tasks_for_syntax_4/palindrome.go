package main

import (
	"fmt"
)

func main() {
	fmt.Println("Satrni kiriting : ")
	var n string

	//agar foydalanuchi son kiritishni hohlasa
	//s := strconv.Itoa(n)    n ni int qilib shu qatorni qo'shish kifoya

	fmt.Scan(&n)
	result := isPalindrome(n)
	fmt.Println(result)
}

func isPalindrome(s string) bool {
	n := len(s)
	for i := 0; i < n/2; i++ {
		if s[i] != s[n-i-1] {
			return false
		}
	}
	return true
}

package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("n ni kiriting: ")
	var n int
	fmt.Scan(&n)
	//faqat n ni o'zini tekshirish
	s := checkFizzBuzz(n)

	// 1 dan n gacha hamma sonlarni tekshirib massivga olish
	//  s := fizzBuzz(n)

	fmt.Println(s)
}

func fizzBuzz(n int) []string {
	s := []string{"1"}
	for i := 2; i <= n; i++ {
		s = append(s, checkFizzBuzz(i))
	}
	return s
}

func checkFizzBuzz(i int) string {
	var s string
	if i%15 == 0 {
		s = "FizzBuzz"
	} else if i%3 == 0 {
		s = "Fizz"
	} else if i%5 == 0 {
		s = "Buzz"
	} else {
		s = strconv.Itoa(i)
	}
	return s
}

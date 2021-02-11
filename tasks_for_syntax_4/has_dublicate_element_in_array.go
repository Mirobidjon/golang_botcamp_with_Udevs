package main

import "fmt"

func main() {
	fmt.Println("Massiv elementlari sonini kiriting: ")
	var n, x int
	arr := []int{}
	used := []bool{}
	fmt.Scan(&n)

	fmt.Println("Massiv elementlarini kiriting : ")
	for i := 0; i < n; i++ {
		fmt.Scan(&x)
		arr = append(arr, x)
		used = append(used, false)
	}
	s := duplicateFinder(arr, used)
	printDuplicateElements(s)
}

func duplicateFinder(arr []int, used []bool) []int {
	dupElements := []int{}
	for i := 0; i < len(arr); i++ {
		var k int = 0
		if !used[i] {
			for j := i; j < len(arr); j++ {
				if arr[i] == arr[j] {
					k++
					used[j] = true
				}
			}
		}
		// ikkitadan ko'p xolat uchun yozildi va hohishga qarab o'zgartirish umkin
		if k >= 2 {
			dupElements = append(dupElements, arr[i])
		}
	}
	return dupElements
}

func printDuplicateElements(s []int) {
	if len(s) == 0 {
		fmt.Println("Massivda bittadan ko'p qatnashgan element mavjud emas!")
		return
	}
	fmt.Print("Massivda bittadan ko'p qatnashgan elementlar: ")
	for _, i := range s {
		fmt.Print(i, " ")
	}

	// shu elementlardan birinchisini chiqarish :
	fmt.Println("\nMassivda bittadan ko'p qatnashgan birinchi element:", s[0])
}

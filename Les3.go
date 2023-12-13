package main

import (
	"fmt"
)

func powIteration(n int64, p int) int64 {
	var result int64 = 1
	for i := 0; i < p; i++ {
		result *= n
	}
	return result
}

func fiboR(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return fiboR(n-1) + fiboR(n-2)
}

func fiboI(n int) int {
	f0 := 0
	f1 := 1
	if n == 0 {
		return 0
	}
	for i := 2; i <= n; i++ {
		f2 := f0 + f1
		f0 = f1
		f1 = f2
	}
	return f1
}

func iPrime(i int) bool {
	for j := 2; j < i; j++ {
		if i%j == 0 {
			return false
		}
	}
	return true
}

func primeDevs(n int) int {
	count := 0
	for i := 2; i <= n; i++ {
		if iPrime(i) {
			count++
		}
	}
	return count
}

func main() {
	var n int64 = 2
	var p = 10
	var f = 9
	var prime = 20
	fmt.Printf("Результат итеративного возведения %d в степень %d = %d\n", n, p, powIteration(n, p))
	fmt.Printf("Результат рекурсивного поиска числа Фибоначчи %d = %d\n", f, fiboR(f))
	fmt.Printf("Результат итеративного поиска числа Фибоначчи %d = %d\n", f, fiboI(f))
	fmt.Printf("Результат поиска простых чисел через перебор делителей %d = %d\n", prime, primeDevs(prime))
}

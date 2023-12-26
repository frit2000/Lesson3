package main

import (
	"fmt"
	"math"
	//	"time"
)

// O(n) итерационное возведение в степень
func powIteration(n int64, p int) int64 {
	var result int64 = 1
	for i := 0; i < p; i++ {
		result *= n
	}
	return result
}

// O(log(n)+ n/2) Возведение через домножение
// квадрат квадрата, пока можно, потом просто домножаем на остаток степени
func powMult(n int64, p int) int64 {
	result := n
	tail := 1
	mult := int(math.Log2(float64(p)))
	for i := 0; i < mult; i++ {
		result *= result
		tail *= 2
	}
	for i := 0; i < (p - tail); i++ {
		result *= n
	}
	return result
}

// O(2log(n)) возведение в степень через двоичное разложение
func powBin(n int64, p int) int64 {
	if p == 0 {
		return 1
	}
	if p == 1 {
		return n
	}
	if p%2 == 0 {
		x := powBin(n, p/2)
		return x * x
	}
	return n * powBin(n, p-1)
}

// O(n^2) Рекурсивный поиск чисел Фибоначчи
func fiboR(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return fiboR(n-1) + fiboR(n-2)
}

// O(n) Итеративный поиск чисел Фибоначчи
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

// поиск чисел Фибоначчи через формулу золотого сечения
func fiboGold(n int) int64 {
	var gold float64
	n1 := float64(n)
	gold = (1 + math.Sqrt(5)) / 2
	f := (math.Pow(gold, n1) - math.Pow((1-gold), n1)) / math.Sqrt(5)
	return int64(f)
}

// Фибоначчи через матрицы
type FiboMatrix struct {
	MatrA [2][2]int
	MatrB [2][2]int
}

func mMult(m1 [2][2]int, m2 [2][2]int) (outMatr [2][2]int) {
	for k := 0; k < len(m1); k++ {
		for i := 0; i < len(m2[0]); i++ {
			for j := 0; j < len(m1[0]); j++ {
				outMatr[k][i] += m1[k][j] * m2[j][i]
			}
		}
	}
	return
}

func (f FiboMatrix) MatrixPowBin(p int) (outMatr [2][2]int) {
	if p == 0 || p == 1 {
		return f.MatrA
	}
	if p%2 == 0 {
		x := f.MatrixPowBin(p / 2)
		return mMult(x, x)
	}
	return mMult(f.MatrA, f.MatrixPowBin(p-1))
}

func main() {
	//var n int64 = 2
	//var p = 30
	var f = 20
	//var prime = 20

	mat := FiboMatrix{
		[2][2]int{{1, 1}, {1, 0}},
		[2][2]int{{1, 1}, {1, 0}},
	}
	//start := time.Now()
	//fmt.Printf("Результат итеративного возведения %d в степень %d = %d\n", n, p, powIteration(n, p))
	//duration := time.Since(start)
	//fmt.Println("время выполнения = ", duration)

	//fmt.Printf("Результат возведения %d в степень %d через домножение = %d\n", n, p, powMult(n, p))
	//fmt.Printf("Результат возведения %d в степень %d через двоичное разложение = %d\n", n, p, powBin(n, p))

	//	start := time.Now()
	//	fmt.Printf("Результат рекурсивного поиска числа Фибоначчи %d = %d\n", f, fiboR(f))
	//	duration := time.Since(start)
	//	fmt.Println("время выполнения = ", duration)

	fmt.Printf("Результат итеративного поиска числа Фибоначчи %d = %d\n", f, fiboI(f))
	//fmt.Printf("Результат поиска через формулу Бине числа Фибоначчи %d = %d\n", f, fiboGold(f))
	fmt.Printf("Результат поиска через умножение матриц числа Фибоначчи %d = %d\n", f, mat.MatrixPowBin(f - 1)[0][0])

	//fmt.Printf("Результат поиска простых чисел через перебор делителей %d = %d\n", prime, primeDevs(prime))
}

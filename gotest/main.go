package main

import (
	"fmt"
)

type deretBilang struct {
	N int
}

func (d deretBilang) prima() []int {
	primes := make([]int, 0, d.N)
	for i := 2; i < d.N; i++ {
		isPrime := true
		for _, j := range primes {
			if i%j == 0 {
				isPrime = false
				continue
			}
		}
		if isPrime == true {
			primes = append(primes, i)

		}
	}
	return primes
}

func (d deretBilang) ganjil() []int {
	odds := make([]int, 0, d.N)
	for i := 0; i < d.N; i++ {
		if i%2 != 0 {
			odds = append(odds, i)
		} else {
			continue
		}
	}
	return odds
}

func (d deretBilang) genap() []int {
	evens := make([]int, 0, d.N)
	for i := 0; i < d.N; i++ {
		if i%2 == 0 && i != 0 {
			evens = append(evens, i)
		} else {
			continue
		}
	}
	return evens
}

func (d deretBilang) fibonacci() []int {
	fibos := make([]int, d.N+1, d.N+2)
	if d.N < 2 {
		fibos = fibos[0:2]
	}
	fibos[0] = 0
	fibos[1] = 1
	for i := 2; i <= d.N; i++ {
		fibos[i] = fibos[i-1] + fibos[i-2]
	}
	return fibos
}

func main() {
	deret := deretBilang{5}
	fmt.Println("Deret Bilangan Prima : ", deret.prima())
	fmt.Println("Deret Bilangan Ganjil : ", deret.ganjil())
	fmt.Println("Deret Bilangan Genap : ", deret.genap())
	fmt.Println("Deret Bilangan Fibonacci : ", deret.fibonacci())
}

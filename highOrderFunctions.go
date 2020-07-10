package main

import (
	"fmt"
	"time"
)

func GerarFibonacci(n int) func() {
	return func() {
		a,b := 0,1

		fib := func() int {
			a,b = b,a+b

			return a
		}
		for i:=0;i<n;i++ {
			fmt.Printf("%d ",fib())
		}
	}
}

func Cronometrar(funcao func()) {
	inicio := time.Now()

	funcao()

	fmt.Printf("\nTempo de Execução: %s\n\n",time.Since(inicio))
}

func main(){
	Cronometrar(GerarFibonacci(10))
	Cronometrar(GerarFibonacci(20))
	Cronometrar(GerarFibonacci(6))
}
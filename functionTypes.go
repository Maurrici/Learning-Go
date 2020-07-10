package main

import "fmt"

type Agregadora func(n,m int) int

func Agregar(valores []int, inicial int, fn Agregadora) int{
	resultado := inicial

	for _, v := range valores{
		resultado = fn(v,resultado)
	}

	return resultado
}

func CalcularSoma(valores []int) int{
	soma := func(n, m int) int {
		return n + m
	}

	return Agregar(valores,0,soma)
}

func CalcularProduto(valores []int) int{
	produto := func(n, m int) int {
		return m * n
	}
	
	return Agregar(valores,1,produto)
}

func main()  {
	valores := []int{-3,1,-8,3,-7,7,8,-1}

	fmt.Println(CalcularSoma(valores))
	fmt.Println(CalcularProduto(valores))
}
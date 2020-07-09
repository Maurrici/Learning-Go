package main

import (
	"fmt"
	"os"
	"strconv"
)

func main(){
	if len(os.Args) != 3 {
		fmt.Println("Uso: <valore> <moeda>")
	}

	moedaOrigem := os.Args[2]
	valorOrigem, err := strconv.ParseFloat(os.Args[1], 64)
	if err != nil{
		fmt.Printf("%s é um valor inválido!", os.Args[1])
		os.Exit(1)
	}

	var valorFinal float64
	var moedaFinal string

	if moedaOrigem == "real"{
		valorFinal = valorOrigem/5.35
		moedaFinal = "dólar"
	}else if moedaOrigem == "dolar"{
		valorFinal = valorOrigem * 5.35
		moedaFinal = "real"
	}else{
		fmt.Printf("A moeda %s não está disponível!", moedaOrigem)
		os.Exit(1)
	}

	fmt.Printf("%.2f %s = %.2f %s", valorOrigem, moedaOrigem, valorFinal, moedaFinal)
}

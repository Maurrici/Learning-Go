package main

import (
	"fmt"
	"os"
	"strconv"
)

func main(){
	if len(os.Args) < 3{
		fmt.Println("Uso: conversor<valores> <unidade>")
		os.Exit(1)
	}

	unidadeOrigem := os.Args[len(os.Args)-1]
	valoresOrigem := os.Args[1:len(os.Args)-1]

	var unidadeDestino string

	if unidadeOrigem == "celsius"{
		unidadeDestino = "fahrenheit"
	}else if unidadeOrigem == "quilometros"{
		unidadeDestino = "milhas"
	}else{
		fmt.Printf("%s não é uma unidade válida.", unidadeOrigem)
		os.Exit(1)
	}

	for i, v := range valoresOrigem{
		valorOrigem, err := strconv.ParseFloat(v, 64)
		if err != nil{
			fmt.Printf("O valor %s na posição %d é inválido", v, i)
			os.Exit(1)
		}

		var valorConvertido float64
		if unidadeDestino == "celsius"{
			valorConvertido = valorOrigem*1.8 + 32
		}else{
			valorConvertido = valorOrigem/1.60934
		}

		fmt.Printf("%.2f %s = %.2f %s\n", valorOrigem, unidadeOrigem, valorConvertido, unidadeDestino)
	}
}



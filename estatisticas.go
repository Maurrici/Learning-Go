package main

import(
	"fmt"
	"os"
	"strings"
)

func main(){
	palavras := os.Args[1:]

	estatisticas := colherEstatisticas(palavras)

	imprimir(estatisticas)
}

func colherEstatisticas(palavras []string) map[string]int {
	estatisticas := make(map[string]int)

	for _ , palavra := range palavras{
		inicial := string(palavra[0])
		inicial = strings.ToUpper(inicial)

		valor, encontrado := estatisticas[inicial]
		if encontrado {
			estatisticas[inicial] = valor + 1
		}else{
			estatisticas[inicial] = 1
		}
	}

	return estatisticas
}

func imprimir(estatisticas map[string]int){
	fmt.Println("Contagem de Palavras iniciadas com cada letra:")

	for inicial, contador := range estatisticas {
		fmt.Printf("%s = %d repetições\n", inicial, contador)
	}
}
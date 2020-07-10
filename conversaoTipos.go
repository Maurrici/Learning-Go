package main

import (
	"fmt"
)

type ListaCompras []string

func main(){
	lista := ListaCompras{"Alface", "Repolho", "Sal"}
	slice := []string{"Alface", "Repolho", "Sal"}

	exibirLista(ListaCompras(slice))
	exibirSlice([]string(lista))
}

func exibirLista(lista ListaCompras){
	fmt.Println(lista)
}

func exibirSlice(slice []string){
	fmt.Println(slice)
}

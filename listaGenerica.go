package main

import (
	"fmt"
)

type ListaGenerica []interface{}

func (lista *ListaGenerica) RemoverIndice(indice int) interface{} {
	l := *lista
	removido := l[indice]
	*lista = append(l[:indice],l[indice+1:]...)
	return removido
}

func (lista *ListaGenerica) RemoverInicio() interface{} {
	return lista.RemoverIndice(0)
}

func (lista *ListaGenerica) RemoverFim() interface{} {
	return lista.RemoverIndice(len(*lista)-1)
}

func main(){
	lista := ListaGenerica{1, "Café", 3.14, "Maurício", false}

	fmt.Println("Lista Original:",lista)

	fmt.Printf("\nRemovendo do inicio: %v foi removido\nLista:%v", lista.RemoverInicio(), lista)
	fmt.Printf("\nRemovendo do indice 2: %v foi removido\nLista:%v", lista.RemoverIndice(2), lista)
	fmt.Printf("\nRemovendo do fim: %v foi removido\nLista:%v", lista.RemoverFim(), lista)
}

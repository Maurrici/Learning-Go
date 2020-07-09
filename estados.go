package main

import (
	"fmt"
	"sort"
)

func main(){
	estados := make(map[string]Estado, 6)

	estados["CE"] = Estado{"Ceará", 100000, "Fortaleza"}
	estados["PI"] = Estado{"Piaui", 123000, "Teresina"}
	estados["MA"] = Estado{"Maranhão", 100987, "São Luís"}
	estados["BA"] = Estado{"Bahia", 875155, "Salvador"}
	estados["GO"] = Estado{"Goiás", 284825, "Goiânia"}
	estados["SP"] = Estado{"São Paulo", 8000000, "São Paulo"}

	exibirInfo(estados)
}

func exibirInfo(estados map[string]Estado){
	indicesOrdenados := make([]string, 0, len(estados))

	for indice := range estados{
		indicesOrdenados = append(indicesOrdenados, indice)

	}

	sort.Strings(indicesOrdenados)

	for _, sigla := range indicesOrdenados{
		fmt.Printf("\n%s :\n\tNome: %s\n\tPopulação:%d\n\tCapital:%s",
			sigla,estados[sigla].nome, estados[sigla].populacao,estados[sigla].capital)
	}
}

type Estado struct {
	nome string
	populacao int
	capital string
}


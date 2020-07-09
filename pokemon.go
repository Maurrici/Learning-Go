package main

import (
	"fmt"
	"sort"
)

func main(){
	//Gerando uma lista com insformações pré salvas
	listTreinadores := make(map[string]Treinador)
	listTreinadores["Ash"] = Treinador{"Ash", []string{"pikachu", "squirtle", "bulbasaur"} }
	listTreinadores["Misty"] = Treinador{"Misty", []string{"psyduck", "togepi"} }
	listTreinadores["Brock"] = Treinador{"Brock", []string{"croagunk", "onix"} }

	//Inicio
	var op int
	loop:
	for{
		renderTreinadores(listTreinadores)
		fmt.Println("Digite uma das Operações a baixo:\n1 - Ver Pokemons de um treinador\n" +
			"2 - Adicionar um Pokemon\n3 - Adicionar um treinador\n0 - Sair")
		fmt.Scan(&op)
		switch op {
		case 1:
			renderPokemons(listTreinadores)
			break
		case 2:
			addPokemon(listTreinadores)
		case 3:
			addTreinador(listTreinadores)
		case 0:
			break loop
		default:
			fmt.Println("Entrada Inválida!")
		}
	}
}
//Funções da aplicação
func renderTreinadores(listTreinadores map[string]Treinador){
	indicesOrdenados := make([]string,0,len(listTreinadores))

	for indice := range listTreinadores{
		indicesOrdenados = append(indicesOrdenados, indice)
	}

	sort.Strings(indicesOrdenados)

	fmt.Println("---------------------------------")
	fmt.Println("Lista de Treinadores")
	fmt.Println("---------------------------------")
	fmt.Printf("Nome\t\tQuantidade de Pokemons\n")
	for _,nome := range indicesOrdenados{
		fmt.Printf("%s\t\t%d\n",nome,listTreinadores[nome].QuantidadePokemon())
	}
}

func renderPokemons(listTreinadores map[string]Treinador){
	var nome string
	fmt.Println("De qual treinador?")
	fmt.Scan(&nome)
	fmt.Println(nome + " tem os seguintes pokemons",listTreinadores[nome].pokemons)
	fmt.Scanln()
}

func addPokemon(listTreinadores map[string]Treinador){
	var nome string
	var newPokemon string
	fmt.Println("A qual treinador?")
	fmt.Scan(&nome)
	fmt.Println("Digite o nome do pokémon:")
	fmt.Scan(&newPokemon)

	treinador := listTreinadores[nome]
	treinador.AddPokemon(newPokemon)
	listTreinadores[treinador.Nome()] = treinador
}

func addTreinador(listTreinadores map[string]Treinador)  {
	var nome string
	var newPokemon string
	pokemons := []string{}
	fmt.Println("A qual treinador?")
	fmt.Scan(&nome)
	for{
		fmt.Println("Digite o nome do pokémon ou 0 para finalizar:")
		fmt.Scan(&newPokemon)
		if newPokemon == "0"{
			break
		}
		pokemons = append(pokemons,newPokemon)
	}

	listTreinadores[nome] = Treinador{nome,pokemons}
}

//Tipo e seus métodos
type Treinador struct {
	nome string
	pokemons []string
}

func (treinador Treinador) Nome() string{
	return treinador.nome
}

func (treinador Treinador) QuantidadePokemon() int{
	return len(treinador.pokemons)
}

func (treinador *Treinador) AddPokemon(newPokemon string){
	treinador.pokemons = append(treinador.pokemons, newPokemon)
}


package main

import (
	"fmt"
	"time"
)

type Operacao interface {
	Calcular() int
	String() string
}

type Soma struct {
	operador1 int
	operador2 int
}

func (s Soma) Calcular() int {
	return s.operador1 + s.operador2
}

func (s Soma) String() string{
	return fmt.Sprintf("%d + %d",s.operador1,s.operador2)
}

type Subtracao struct {
	operador1 int
	operador2 int
}

func (s Subtracao) Calcular() int{
	return s.operador1 - s.operador2
}

func (s Subtracao) String() string {
	return fmt.Sprintf("%d - %d",s.operador1,s.operador2)
}

type Idade struct {
	anoNascimento int
}

func (idade Idade) Calcular() int{
	return time.Now().Year() - idade.anoNascimento
}

func (idade Idade) String() string{
	return fmt.Sprintf("Idade desde %d",idade.anoNascimento)
}

func main(){
	operacoes := make([]Operacao, 4)

	operacoes[0] = Soma{10, 20}
	operacoes[1] = Soma{30, 40}
	operacoes[2] = Subtracao{100,120}
	operacoes[3] = Subtracao{28,45}

	fmt.Printf("Valor Final: %d\n", calcularOperacoes(operacoes))

	idades := make([]Operacao,3)

	idades[0] = Idade{2001}
	idades[1] = Idade{2000}
	idades[2] = Idade{1999}

	fmt.Printf("Valor Final: %d", calcularOperacoes(idades))
}

func calcularOperacoes(operacoes []Operacao) int {
	resultado := 0
	for _, op := range operacoes{
		valor := op.Calcular()
		fmt.Printf("%v = %d\n", op.String(), valor)
		resultado += valor
	}
	return resultado
}
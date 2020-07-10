package main

import (
	"fmt"
	"io"
)

type LeitorStrings struct {}

func (leitor LeitorStrings) Read(p []byte) (int, error){
	p[0] = 'A'
	p[1] = 'B'
	p[2] = 'C'
	p[3] = 'D'

	return len(p), nil
}

func lerString(r io.Reader) string{
	p := make([]byte, 4)
	r.Read(p)
	return string(p)
}

func main(){
	leitor := LeitorStrings{}
	fmt.Println(lerString(leitor))
}
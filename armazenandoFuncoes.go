package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main(){
	expr := regexp.MustCompile("\\b\\w")
	texto := "maurício de moura dos santos"

	transformadora := func(s string) string{

		return strings.ToUpper(s)
	}

	fmt.Println(transformadora(texto))
	fmt.Printf(expr.ReplaceAllStringFunc(texto,transformadora))
}

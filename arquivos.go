package main

import(
	"fmt"
	"os"
)

func CriarArquivos(dirBase string, arquivos ...string){
	for _, nome := range arquivos{
		caminhoArquivo := fmt.Sprintf("%s\\%s.%s",dirBase,nome,"txt")

		arq, err := os.Create(caminhoArquivo)

		defer arq.Close()

		if err != nil{
			fmt.Printf("Não foi possível criar o arquivo %s: %v\n",nome,err)
			os.Exit(1)
		}

		fmt.Printf("O arquivo %s foi criado com sucesso!\n", arq.Name())
	}
}

func main(){
	dirBase := "C:\\Users\\Maurício de Moura\\go\\src\\awesomeProject"

	CriarArquivos(dirBase)
	CriarArquivos(dirBase,"teste1")
	CriarArquivos(dirBase,"teste2","teste3","teste4")
}
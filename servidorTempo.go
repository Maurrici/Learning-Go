package main

import (
	"fmt"
	"net/http"
	"time"
)

func main(){
	http.HandleFunc(//Registra um serviço
		"/tempo",
		func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w,"%s",time.Now().Format("02/01/2006 15:04:05"))
		})

	http.ListenAndServe(":8080",nil)
}
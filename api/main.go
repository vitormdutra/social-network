package main

import (
	"api/src/router"
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Rodando API!")

	r := router.Gerar()

	http.ListenAndServe(":1234", r)
}

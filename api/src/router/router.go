package router

import (
	"api/src/router/rotas"
	"github.com/gorilla/mux"
)

// Gerar vau retornar uma rota
func Gerar() *mux.Router {
	r := mux.NewRouter()
	return rotas.Configurar(r)
}

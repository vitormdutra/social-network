package router

import "github.com/gorilla/mux"

// Gerar vau retornar uma rota
func Gerar() *mux.Router {
	return mux.NewRouter()
}

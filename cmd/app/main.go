package main

import (
	"github.com/gorilla/mux"
	"member_role/cmd/app/handler"
	"member_role/internal/controller"
	"member_role/internal/repository"
	"member_role/internal/repository/infla"
	"member_role/internal/service"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	h := getHandler()

	r.PathPrefix("/member-roles").Handler(h)
	http.ListenAndServe(":8087", r)
}

func getHandler() http.Handler {
	return handler.NewHandler(controller.NewController(service.NewService(repository.NewRepository(infla.NewDB()))))
}

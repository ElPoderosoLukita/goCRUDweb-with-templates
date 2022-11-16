package main

import (
	"net/http"

	"github.com/ElPoderosoLukita/goCRUDweb/router"
	"github.com/gorilla/mux"
)

func main() {
	api := router.API{}
	r := mux.NewRouter()

	api.CreateRouter(r)

	http.ListenAndServe(":8081", r)

}

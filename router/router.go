package router

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (a API) CreateRouter(mux *mux.Router) {
	mux.HandleFunc("/get/users", a.GetUsers).Methods(http.MethodGet)
	mux.HandleFunc("/post/user", a.PostUser).Methods(http.MethodPost)
	mux.HandleFunc("/update/user/{id}", a.UpdateUser).Methods(http.MethodPut)
	mux.HandleFunc("/delete/user/{id}", a.DeleteUser).Methods(http.MethodDelete)
}

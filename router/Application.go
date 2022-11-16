package router

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type API struct {
}

type Users struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Lastname string `json:"lastname"`
	Age      int    `json:"age"`
}

var (
	users []*Users = []*Users{
		{
			ID:       1,
			Name:     "Lucas",
			Lastname: "Da Rosa",
			Age:      15,
		},
	}
)

//HandlerFunc
func (a API) GetUsers(w http.ResponseWriter, r *http.Request) {
	temp, err := template.New("getUsers.html").ParseFiles("templates/getUsers.html")
	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusOK)
	temp.Execute(w, users)
}

func (a API) PostUser(w http.ResponseWriter, r *http.Request) {
	u := Users{}
	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		panic(err)
	}

	users = append(users, &u)
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "User created correctly!")
}

func (a API) UpdateUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	idInt, err := strconv.Atoi(id)
	user := Users{}
	verification := false

	if err != nil {
		http.NotFound(w, r)
		fmt.Fprintf(w, "El id que has proporcionado es incorrecto.")
	}

	json.NewDecoder(r.Body).Decode(&user)

	for _, u := range users {
		if u.ID == idInt {
			u.Name = user.Name
			u.Lastname = user.Lastname
			u.Age = user.Age
			verification = true
			break
		}
		verification = false
	}

	if verification {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "User updated correctly!")
	} else {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "The id that you sent was bad")
	}
}

func (a API) DeleteUser(w http.ResponseWriter, r *http.Request) {
	index := 0
	verification := false
	params := mux.Vars(r)
	id := params["id"]
	idInt, err := strconv.Atoi(id)

	if err != nil {
		http.NotFound(w, r)
		fmt.Fprintf(w, "El id que has proporcionado es incorrecto.")
	}

	for i, v := range users {
		if v.ID == idInt {
			index = i
			verification = true
			break
		}

		verification = false
	}

	if verification {
		users = append(users[:index], users[index+1:]...)
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "User deleted correctly!")
	} else {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "The id that you sent was bad")
	}
}

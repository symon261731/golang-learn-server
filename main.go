package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"test-server/package/instances"
	"test-server/package/mockDB"
	"test-server/package/transport/rest"
)

var PORT = ":4000"

func main() {
	r := mux.NewRouter()

	var fakeDb = mockDB.MockDB{
		List: map[int]*instances.User{},
	}

	r.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		rest.GetAllUsers(fakeDb.List)
	})
	r.HandleFunc("/create", func(writer http.ResponseWriter, request *http.Request) {
		rest.CreateUser(writer, request, &fakeDb)
		fmt.Println(fakeDb.List)
	})
	r.HandleFunc("/friends/{id}", func(writer http.ResponseWriter, request *http.Request) {
		rest.ShowFriends(writer, request, &fakeDb)
	})
	r.HandleFunc("/user", func(writer http.ResponseWriter, request *http.Request) {
		rest.DeleteUserById(writer, request, &fakeDb)
	})
	r.HandleFunc("/make_friends", func(writer http.ResponseWriter, request *http.Request) {
		rest.MakeFriends(writer, request, &fakeDb)
	})
	r.HandleFunc("/{user_id}", func(writer http.ResponseWriter, request *http.Request) {
		rest.ChangeAgeOfUser(writer, request, &fakeDb)
	})

	log.Printf("Веб-сервер запущен на http://127.0.0.1%s", PORT)
	err := http.ListenAndServe(PORT, r)

	log.Fatal(err)

}

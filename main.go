package main

import (
	"fmt"
	"log"
	"net/http"
	"test-server/package/mockDB"
	"test-server/package/transport/rest"
)

var PORT = ":4000"

func main() {
	mux := http.NewServeMux()

	var fakeDb mockDB.MockDB

	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		rest.GetAllUsers(fakeDb.List)
	})
	mux.HandleFunc("/create", func(writer http.ResponseWriter, request *http.Request) {
		rest.CreateUser(writer, request, fakeDb)
		fmt.Println(fakeDb.List)
	})

	log.Printf("Веб-сервер запущен на http://127.0.0.1%s", PORT)
	err := http.ListenAndServe(PORT, mux)

	log.Fatal(err)

}

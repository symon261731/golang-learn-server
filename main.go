package main

import (
	"log"
	"net/http"
	"test-server/package/mockDB"
	"test-server/package/transport/rest"
)

func main() {
	PORT := ":4000"
	var users []mockDB.User
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		rest.GetAllUsers(users)
	})
	mux.HandleFunc("/create", func(writer http.ResponseWriter, request *http.Request) {
		rest.CreateUser(request, users)
	})

	log.Printf("Веб-сервер запущен на http://127.0.0.1%s", PORT)
	err := http.ListenAndServe(PORT, mux)

	log.Fatal(err)

}

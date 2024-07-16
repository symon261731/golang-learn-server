package rest

import (
	"bytes"
	"log"
	"net/http"
	"test-server/package/mockDB"
)

// Разбить на
// create
// make_friends
// deleteUser
// getFriendsOfUser
// UpdateAgeOfUser

func GetAllUsers(Users []mockDB.User) {
	log.Println(Users)
}

func CreateUser(request *http.Request, Users []mockDB.User) {
	if request.Method == "POST" {
		buf := new(bytes.Buffer)
		buf.ReadFrom(request.Body)
		log.Println(buf.String())
	}
}

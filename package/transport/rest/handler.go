package rest

import (
	"encoding/json"
	"log"
	"net/http"
	"test-server/package/instances"
	"test-server/package/mockDB"
)

// Разбить на
// create
// make_friends
// deleteUser
// getFriendsOfUser
// UpdateAgeOfUser

func GetAllUsers(Users []instances.User) {
	log.Println("List of users")
	log.Println(Users)
}

func CreateUser(writer http.ResponseWriter, request *http.Request, db mockDB.MockDB) {
	if request.Method == "POST" {
		var jsonData instances.CreateUserFormData
		err := json.NewDecoder(request.Body).Decode(&jsonData)

		if err != nil {
			http.Error(writer, err.Error(), http.StatusBadRequest)
			log.Println(err.Error())
			return
		}

		newUser := instances.User{
			Id:      len(db.List) + 1,
			Name:    jsonData.Name,
			Age:     jsonData.Age,
			Friends: make([]string, 0),
		}

		db.AddNewUser(newUser)

		log.Println("New user created", newUser)
	} else {
		http.Error(writer, "Invalid HTTP verb.", http.StatusBadRequest)
	}
}

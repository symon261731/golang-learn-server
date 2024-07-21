package rest

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
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

// CreateUser эндпоинт для создания пользователя
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
			Friends: make([]instances.FriendsOfUser, 0),
		}

		db.AddNewUser(newUser)

		log.Println("New user created", newUser)
	} else {
		http.Error(writer, "Invalid HTTP verb.", http.StatusBadRequest)
	}
}

// ShowFriends эндпоинт для показа друзей пользователя
func ShowFriends(writer http.ResponseWriter, request *http.Request, db mockDB.MockDB) {
	if request.Method == "GET" {

		vars := mux.Vars(request)

		id := vars["id"]

		IntId, err := strconv.Atoi(id)

		if err != nil {
			log.Println("invalid syntax of id", err)
			http.Error(writer, "invalid syntax of id", http.StatusBadRequest)
		}

		user, err2 := db.ShowAllFriendsOfUser(IntId)
		if err2 != nil {
			log.Println(err2, id)
			http.Error(writer, "not found friends of this user", http.StatusNotFound)
			return
		}
		buf := &bytes.Buffer{}
		gob.NewDecoder(buf).Decode(user)

		writer.Write(buf.Bytes())
	} else {
		http.Error(writer, "Invalid HTTP verb.", http.StatusBadRequest)
	}
}

// DeleteUserById эндпоинт для удаления пользователя по id
func DeleteUserById(writer http.ResponseWriter, request *http.Request, db mockDB.MockDB) {
	if request.Method == "DELETE" {
		params := request.URL.Query()
		//TODO разрабраться почему надо тянуть нулевой индекс
		id := params["target_id"][0]
		fmt.Println(id)
		IntId, err := strconv.Atoi(id)

		if err != nil {
			log.Println("invalid syntax target_id", IntId)
			http.Error(writer, "invalid syntax of target_id", http.StatusBadRequest)
		}

		db.DeleteUser(1)

		buf := &bytes.Buffer{}
		successResponseText := "user delete success"
		err2 := gob.NewDecoder(buf).Decode(&successResponseText)
		if err2 != nil {
			return
		}

		writer.Write(buf.Bytes())
	}
}

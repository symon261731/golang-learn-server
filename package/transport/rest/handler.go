package rest

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
	"test-server/package/instances"
	"test-server/package/mockDB"
)

func GetAllUsers(Users map[int]*instances.User) {
	log.Println("List of users")
	for _, user := range Users {
		log.Println(*user)
	}
}

// CreateUser эндпоинт для создания пользователя
func CreateUser(writer http.ResponseWriter, request *http.Request, db *mockDB.MockDB) {
	if request.Method != "POST" {
		http.Error(writer, "Invalid HTTP verb.", http.StatusBadRequest)
		return
	}

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

	jData, errJson := json.Marshal(newUser.Id)

	if errJson != nil {
		log.Println(errJson)
		http.Error(writer, "server error", http.StatusInternalServerError)
		return
	}

	log.Println("New user created", newUser)

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusCreated)
	writer.Write(jData)

}

// ShowFriends эндпоинт для показа друзей пользователя
func ShowFriends(writer http.ResponseWriter, request *http.Request, db *mockDB.MockDB) {
	if request.Method != "GET" {
		http.Error(writer, "Invalid HTTP verb.", http.StatusBadRequest)
		return
	}

	vars := mux.Vars(request)

	id := vars["id"]

	IntId, err := strconv.Atoi(id)

	if err != nil {
		log.Println("invalid syntax of id", err)
		http.Error(writer, "invalid syntax of id", http.StatusBadRequest)
	}

	friends, err2 := db.ShowAllFriendsOfUser(IntId)

	if err2 != nil {
		log.Println(err2, id)
		http.Error(writer, err2.Error(), http.StatusBadRequest)
		return
	}

	log.Println(friends)
	data, errJson := json.Marshal(friends)

	if errJson != nil {
		log.Println(errJson)
		http.Error(writer, "not found friends of this user", http.StatusInternalServerError)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.Write(data)
}

// DeleteUserById эндпоинт для удаления пользователя по id
func DeleteUserById(writer http.ResponseWriter, request *http.Request, db *mockDB.MockDB) {
	if request.Method != "DELETE" {
		http.Error(writer, "need delete method", http.StatusNotFound)
		return
	}

	params := request.URL.Query()
	id := params["target_id"][0]
	IntId, err := strconv.Atoi(id)

	if err != nil {
		log.Println("invalid syntax target_id", IntId)
		http.Error(writer, "invalid syntax of target_id", http.StatusBadRequest)
		return
	}

	nameOfDeleteUser, errByDelete := db.DeleteUser(IntId)

	if errByDelete != nil {
		http.Error(writer, fmt.Sprintf("not found user with id %s", id), http.StatusBadRequest)
		return
	}

	data, errJson := json.Marshal(fmt.Sprintf("user %s delete success", nameOfDeleteUser))
	if errJson != nil {
		log.Println(errJson)
		http.Error(writer, "not found friends of this user", http.StatusInternalServerError)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.Write(data)
}

// MakeFriends эндпоинт для связывания двух пользователей
func MakeFriends(writer http.ResponseWriter, request *http.Request, db *mockDB.MockDB) {
	if request.Method != "POST" {
		http.Error(writer, "invalid request", http.StatusNotFound)
		return
	}

	var jsonData instances.PostIdsFriends
	err := json.NewDecoder(request.Body).Decode(&jsonData)

	if err != nil {
		log.Println("invalid DTO", err)
		http.Error(writer, "invalid DTO", http.StatusBadRequest)
		return
	}

	resultString, err := db.MakeNewFriend(jsonData.Source_id, jsonData.Target_id)

	if err != nil {
		log.Println("error", err)
		http.Error(writer, "error", http.StatusInternalServerError)
		return
	}

	log.Println(resultString)

	data, jsonErr := json.Marshal(resultString)

	if jsonErr != nil {
		log.Println("error", jsonErr)
		http.Error(writer, "error by server", http.StatusInternalServerError)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.Write(data)
}

func ChangeAgeOfUser(writer http.ResponseWriter, request *http.Request, db *mockDB.MockDB) {
	if request.Method != "PUT" {
		http.Error(writer, "invalid request", http.StatusNotFound)
		return
	}
	vars := mux.Vars(request)

	id := vars["user_id"]

	intUserId, errFormatId := strconv.Atoi(id)

	if errFormatId != nil {
		log.Println("invalid syntax of dynamic url", errFormatId)
		http.Error(writer, "invalid syntax of user_id", http.StatusBadRequest)
		return
	}

	var jsonData instances.PutNewAgeJson
	err := json.NewDecoder(request.Body).Decode(&jsonData)

	if err != nil {
		log.Println("invalid DTO", err)
		http.Error(writer, "invalid DTO", http.StatusBadRequest)
		return
	}

	intNewAge, errFormatAge := strconv.Atoi(jsonData.NewAge)

	if errFormatAge != nil {
		log.Println("invalid syntax of newAge", errFormatAge)
		http.Error(writer, "invalid syntax of `new age`", http.StatusBadRequest)
		return
	}

	successMessage, errChangeAge := db.ChangeAgeOfUser(intUserId, intNewAge)

	if errChangeAge != nil {
		http.Error(writer, "server error", http.StatusInternalServerError)
		return
	}

	data, jsonErr := json.Marshal(successMessage)

	if jsonErr != nil {
		log.Println("error", jsonErr)
		http.Error(writer, "error by server", http.StatusInternalServerError)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.Write(data)
}

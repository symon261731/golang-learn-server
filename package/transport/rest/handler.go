package rest

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
	"test-server/package/instances"
	"test-server/package/mockDB"
)

func GetAllUsers(Users *[]instances.User) {
	log.Println("List of users")
	log.Println(Users)
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

	jData, _ := json.Marshal(newUser.Id)

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
		http.Error(writer, "not found friends of this user", http.StatusNotFound)
		return
	}

	log.Println(friends)
	buf := &bytes.Buffer{}
	// не могу отправить данные обратно на клиент
	err3 := gob.NewDecoder(buf).Decode(&friends)
	if err3 != nil {
		http.Error(writer, "invalid syntax of id", http.StatusInternalServerError)
		log.Println(err3)
		return
	}

	writer.Write(buf.Bytes())
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

	db.DeleteUser(1)

	buf := &bytes.Buffer{}
	successResponseText := "user delete success"
	err2 := gob.NewDecoder(buf).Decode(&successResponseText)
	if err2 != nil {
		log.Println("error", err2)
		http.Error(writer, "error by server", http.StatusInternalServerError)
		return
	}

	writer.Write(buf.Bytes())
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

	err = db.MakeNewFriend(jsonData.Source_id, jsonData.Target_id)

	if err != nil {
		log.Println("error", err)
		http.Error(writer, "error", http.StatusInternalServerError)
		return
	}
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

	err = db.ChangeAgeOfUser(intUserId, intNewAge)

	if err != nil {
		http.Error(writer, "server error", http.StatusInternalServerError)
		return
	}

}

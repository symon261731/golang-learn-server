package mockDB

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"test-server/package/instances"
	"test-server/package/utils"
)

type MockDB struct {
	List map[int]*instances.User
}

func (db *MockDB) AddNewUser(newUser instances.User) {
	db.List[newUser.Id] = &newUser
}

func (db *MockDB) MakeNewFriend(sendingUserId string, receivingUserId string) (string, error) {
	intSendingUserId, formatToIntErr := strconv.Atoi(sendingUserId)
	if formatToIntErr != nil {
		return "", formatToIntErr
	}
	intReceivingUser, formatToIntErr2 := strconv.Atoi(receivingUserId)
	if formatToIntErr2 != nil {
		return "", formatToIntErr2
	}

	sendingUser := db.List[intSendingUserId]
	receivingUser := db.List[intReceivingUser]

	if utils.CheckUserInFriendList(sendingUser.Friends, receivingUser.Id) {
		var err = errors.New("this user is friend already")
		return "", err
	}

	if utils.CheckUserInFriendList(receivingUser.Friends, sendingUser.Id) {
		var err = errors.New("this user is friend already")
		return "", err
	}

	db.List[intSendingUserId].Friends = append(db.List[intSendingUserId].Friends, instances.FriendsOfUser{Id: receivingUser.Id, Name: receivingUser.Name})
	db.List[intReceivingUser].Friends = append(db.List[intReceivingUser].Friends, instances.FriendsOfUser{Id: sendingUser.Id, Name: sendingUser.Name})

	resultString := fmt.Sprintf("user %s and %s starts be a friends", db.List[intSendingUserId].Name, db.List[intReceivingUser].Name)

	return resultString, nil
}

func (db *MockDB) DeleteUser(userId int) (string, error) {

	neededUser := db.List[userId]

	if neededUser != nil {
		for _, friend := range neededUser.Friends {
			var newListOfFriend = utils.FilterFriendsOfUser(db.List[friend.Id].Friends, neededUser.Id)

			db.List[friend.Id].Friends = newListOfFriend
		}

		nameOfUser := neededUser.Name
		log.Println(neededUser.Id, neededUser.Name, "deleteSucess")
		delete(db.List, userId)

		return nameOfUser, nil
	}

	return "", errors.New("пользователя с таким id не существует")
}

func (db *MockDB) ShowAllFriendsOfUser(userId int) ([]instances.FriendsOfUser, error) {

	neededUser := db.List[userId]

	if neededUser != nil {
		return neededUser.Friends, nil
	}

	return []instances.FriendsOfUser{}, errors.New(fmt.Sprintf("пользователя с таким id не существует"))

}

func (db *MockDB) ChangeAgeOfUser(userId int, newAge int) (string, error) {

	neededUser := db.List[userId]

	if neededUser != nil {
		log.Println("user age changed", db.List[userId].Name, db.List[userId].Age)
		neededUser.Age = newAge
		return "возраст пользователя успешно обновлен", nil
	}

	return "", errors.New("пользователь не найден")
}

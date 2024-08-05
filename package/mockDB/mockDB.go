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
	List []instances.User
}

func (db *MockDB) AddNewUser(newUser instances.User) {
	db.List = append(db.List, newUser)

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

	indexOfSendingUser := utils.FindIndexOfUserById(db.List, intSendingUserId)
	indexOfReceivingUser := utils.FindIndexOfUserById(db.List, intReceivingUser)

	sendingUser := db.List[indexOfSendingUser]
	receivingUser := db.List[indexOfReceivingUser]

	if utils.CheckUserInFriendList(sendingUser.Friends, receivingUser.Id) {
		var err = errors.New("this user is friend already")
		return "", err
	}

	if utils.CheckUserInFriendList(receivingUser.Friends, sendingUser.Id) {
		var err = errors.New("this user is friend already")
		return "", err
	}

	db.List[indexOfSendingUser].Friends = append(db.List[indexOfSendingUser].Friends, instances.FriendsOfUser{Id: receivingUser.Id, Name: receivingUser.Name})
	db.List[indexOfReceivingUser].Friends = append(db.List[indexOfReceivingUser].Friends, instances.FriendsOfUser{Id: sendingUser.Id, Name: sendingUser.Name})

	resultString := fmt.Sprintf("user %s and %s started to be a friend", db.List[indexOfSendingUser].Name, db.List[indexOfReceivingUser].Name)

	return resultString, nil
}

func (db *MockDB) DeleteUser(userId int) {
	var newUserList []instances.User

	for _, user := range db.List {
		if user.Id == userId {
			continue
		}

		filteredUserList := utils.FilterFriendsOfUser(user.Friends, userId)
		user.Friends = filteredUserList
		newUserList = append(newUserList, user)
	}

	// Вот тут кажется достачный хардкод, но и тут вроде ДБ по хроошему должна быть
	db.List = newUserList

}
func (db *MockDB) ShowAllFriendsOfUser(userId int) ([]instances.FriendsOfUser, error) {
	indexOfNeededUser := utils.FindIndexOfUserById(db.List, userId)
	if indexOfNeededUser == -1 {
		return []instances.FriendsOfUser{}, errors.New("not found friends by this id")
	}

	return db.List[indexOfNeededUser].Friends, nil
}

func (db *MockDB) ChangeAgeOfUser(userId int, newAge int) error {
	indexOfNeededUser := utils.FindIndexOfUserById(db.List, userId)
	if indexOfNeededUser == -1 {
		return errors.New("not found user")
	}

	db.List[indexOfNeededUser].Age = newAge

	log.Println("user age changed", db.List[indexOfNeededUser].Name, db.List[indexOfNeededUser].Age)
	return nil
}

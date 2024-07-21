package mockDB

import (
	"errors"
	"test-server/package/instances"
	"test-server/package/utils"
)

type MockDB struct {
	List []instances.User
}

func (db *MockDB) AddNewUser(newUser instances.User) {
	db.List = append(db.List, newUser)
}

func (db *MockDB) MakeNewFriend(sendingUserId int, receivingUserId int) {

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

func (db *MockDB) ChangeAgeOfUser(userId int, newAge int) {

}

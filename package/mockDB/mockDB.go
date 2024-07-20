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

}
func (db *MockDB) ShowAllFriendsOfUser(userId int) ([]string, error) {
	indexOfNeededUser := utils.FindIndexOfUserById(db.List, userId)

	if indexOfNeededUser == -1 {
		return []string{}, errors.New("not found friends by this id")
	}

	return db.List[indexOfNeededUser].Friends, nil
}

func (db *MockDB) ChangeAgeOfUser(userId int, newAge int) {

}

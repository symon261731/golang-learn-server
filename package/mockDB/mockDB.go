package mockDB

import (
	"test-server/package/instances"
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
func (db *MockDB) ShowAllFriendsOfUser(userId int) {

}

func (db *MockDB) ChangeAgeOfUser(userId int, newAge int) {

}

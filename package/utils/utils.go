package utils

import "test-server/package/instances"

func FindIndexOfUserById(list []instances.User, id int) int {

	for i, user := range list {
		if id == user.Id {
			return i
		}
	}

	return -1

}

func FilterFriendsOfUser(listOfUser []instances.FriendsOfUser, filterById int) []instances.FriendsOfUser {
	var filteredSliceOfFriends []instances.FriendsOfUser
	for _, userFriend := range listOfUser {

		if userFriend.Id != filterById {
			filteredSliceOfFriends = append(filteredSliceOfFriends, userFriend)
		}

	}

	return filteredSliceOfFriends
}

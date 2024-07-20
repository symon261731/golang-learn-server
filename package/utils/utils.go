package utils

import "test-server/package/mockDB"

func FindIndexOfUserById(list mockDB.MockDB, id int) int {

	for i, user := range list.List {
		if id == user.Id {
			return i
		}
	}

	return -1

}

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

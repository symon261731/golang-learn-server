package instances

type CreateUserFormData struct {
	Name string
	Age  int
}

type FriendsOfUser struct {
	Id   int
	Name string
}

type User struct {
	Id      int
	Name    string
	Age     int
	Friends []FriendsOfUser
}

type PostIdsFriends struct {
	Source_id string
	Target_id string
}

type PutNewAgeJson struct {
	NewAge string `json:"new age"`
}

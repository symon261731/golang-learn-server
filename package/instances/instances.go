package instances

type CreateUserFormData struct {
	Name string
	Age  int
}

type User struct {
	Id      int
	Name    string
	Age     int
	Friends []string
}

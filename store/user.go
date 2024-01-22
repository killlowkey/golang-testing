package store

type User struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type UserStore interface {
	List() ([]*User, error)
	GetUserById(int) (*User, error)
}

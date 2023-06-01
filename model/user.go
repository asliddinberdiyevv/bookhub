package model

type User struct {
	FirstName string `json:"firstname" db:"firstname"`
	LastName string `json:"lastname" db:"lastname"`
	Username string `json:"username" db:"username"`
	Password string `json:"password" db:"hashed_password"`
}
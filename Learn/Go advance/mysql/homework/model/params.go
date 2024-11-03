package model

type User struct {
	Id       int64  `json:"id"`
	Name     string `json:"name"`
	Phone    int    `json:"phone"`
	Password string `json:"password"`
}

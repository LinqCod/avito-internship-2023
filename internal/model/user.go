package model

type User struct {
	Id       int64  `json:"id"`
	Username string `json:"username"`
}

type CreateUserDTO struct {
	Username string `json:"username"`
}

type CreateUserResponse struct {
	Id int64 `json:"id"`
}

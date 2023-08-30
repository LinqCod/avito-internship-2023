package dto

type CreateUserDTO struct {
	Username string `json:"username"`
}

type CreateUserResponse struct {
	Id int64 `json:"id"`
}

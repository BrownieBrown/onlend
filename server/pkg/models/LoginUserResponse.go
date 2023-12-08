package models

type LoginUserResponse struct {
	AccessToken string
	Id          string `json:"id"`
	Username    string `json:"username"`
}

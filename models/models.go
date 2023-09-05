package models

type User struct {
	User_ID  int    `json:"user_id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Mood     string `json:"mood"`
}

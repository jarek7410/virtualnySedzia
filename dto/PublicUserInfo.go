package dto

type PublicUserInfo struct {
	Username string `json:"username"`
	Name     string `json:"name"`
	Surname  string `json:"surname"`
	Email    string `json:"email"`
	PID      string `json:"pid"`
	ID       uint
}

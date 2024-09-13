package dto

type UserDataMe struct {
	ID       uint   `json:"ID"`
	Username string `json:"username"`
	Email    string `json:"email"`
	PID      string `json:"pid"`
	Name     string `json:"name"`
	Surname  string `json:"surname"`
}

type UserChangeDataMe struct {
	//Username string `json:"username,omitempty"`
	Email   string `json:"email,omitempty"`
	PID     string `json:"pid,omitempty"`
	Name    string `json:"name,omitempty"`
	Surname string `json:"surname,omitempty"`
}

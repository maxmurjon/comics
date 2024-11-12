package models

// type Register struct {
// 	Name     string `json:"name"`
// 	Login    string `json:"login"`
// 	Password string `json:"password"`
// }

type Login struct {
	PhoneNumber    string `json:"phone_number"`
	Password string `json:"password"`
}

type LoginResponse struct {
	UserData *User  `json:"user_data"`
	Token    string `json:"token"`
}
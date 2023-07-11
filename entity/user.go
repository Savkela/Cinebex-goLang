package entity

type User struct {
	Name                string `json:"name"`
	DateOfBirth         string `json:"dateOfBirth"`
	Email               string `json:"email"`
	Username            string `json:"username"`
	Password            string `json:"password"`
	AccountConformation bool   `json:"accountConformation"`
	Role                string `json:"role"`
}

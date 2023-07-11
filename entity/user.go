package entity

import "time"

type Role string

const (
	RoleBuyer         Role = "Buyer"
	RoleAdministrator Role = "Administrator"
)

type User struct {
	ID                  int64         `json:"id"`
	Name                string        `json:"name"`
	DateOfBirth         time.Time     `json:"dateOfBirth"`
	Email               string        `json:"email"`
	Username            string        `json:"username"`
	Password            string        `json:"password"`
	AccountConformation bool          `json:"accountConformation"`
	Role                Role          `json:"role"`
	Reservations        []Reservation `json:"reservations"`
}

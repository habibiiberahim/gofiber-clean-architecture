package schemas

import "time"

type SchemaAuth struct {
	ID        string    `json:"id"`
	Fullname  string    `json:"fullname" `
	Email     string    `json:"email"`
	Token     string    `json:"token"`
	Password  string    `json:"password"`
	Cpassword string    `json:"confirm_password" `
	Active    bool      `json:"active"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
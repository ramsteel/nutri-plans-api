package dto

import "time"

type RegisterRequest struct {
	Email     string    `json:"email" validate:"required,email"`
	Password  string    `json:"password" validate:"required,min=8"`
	Username  string    `json:"username" validate:"required"`
	FirstName string    `json:"first_name" validate:"required"`
	LastName  string    `json:"last_name" validate:"required"`
	Dob       time.Time `json:"dob" validate:"required"`
	Gender    string    `json:"gender" validate:"required,oneof=M F"`
	CountryID uint      `json:"country_id" validate:"required,min=1"`
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

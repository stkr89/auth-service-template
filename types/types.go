package types

import "github.com/google/uuid"

type CreateUserRequest struct {
	FirstName string `json:"firstName" validate:"required" conform:"name"`
	LastName  string `json:"lastName" validate:"required" conform:"name"`
	Email     string `json:"email" validate:"required" conform:"email"`
	Password  string `json:"password" validate:"required,min=8"`
}

type CreateUserResponse struct {
	ID        uuid.UUID `json:"id" validate:"required"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Email     string    `json:"email"`
}

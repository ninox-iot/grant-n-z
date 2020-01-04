package model

import "github.com/google/uuid"

// User data in jwt
type AuthUser struct {
	UserUuid  uuid.UUID `json:"user_uuid"`
	UserId    int       `json:"user_id"`
	UserEmail string    `json:"user_email"`
	ServiceId int       `json:"service_id"`
	Expires   string    `json:"expires"`
	RoleId    int       `json:"role_id"`
}

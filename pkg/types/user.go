package types

import (
	"github.com/google/uuid"
	_ "github.com/google/uuid"
)

type User struct {
	Guid       uuid.UUID
	Login      string
	Password   string
	Name       string
	Email      string
	VerifyCode string
	CreatedAt  string
	UpdatedAt  string
}

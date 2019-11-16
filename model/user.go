package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// User :
type User struct {
	ID                primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name              string             `json:"name" bson:"name"`
	Email             string             `json:"email" bson:"email"`
	Password          string             `json:"-" bson:"password"`
	CreatedAt         time.Time          `json:"createAt,omitempty" bson:"createdAt,omitempty"`
	UpdatedAt         time.Time          `json:"updateAt,omitempty" bson:"updatedAt,omitempty"`
	PasswordUpdatedAt time.Time          `json:"passwordUpdatedAt,omitempty" bson:"passwordUpdatedAt,omitempty"`
}

// EmployeeRank :
type EmployeeRank struct {
	Employee User
	Rank     UserRank
}

// Role :
type Role struct {
	Employee User     `json:"user"`
	Role     UserRole `json:"role"`
}

// UserRole :
type UserRole struct {
	Code     int    `json:"code,omitempty"`
	CodeName string `json:"codeName,omitempty"`
}

// UserRank :
type UserRank struct {
	Name string `json:"name,omitempty"`
	Code string `json:"code,omitempty"`
}

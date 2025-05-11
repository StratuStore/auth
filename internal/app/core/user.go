package core

import (
	"github.com/google/uuid"
)

type User struct {
	ID       uuid.UUID `json:"id" bson:"_id,omitempty"`
	Sub      string    `json:"sub"`
	Email    string    `json:"email"`
	Name     string    `json:"name"`
	Picture  string    `json:"picturePath"`
	Sessions []Session `json:"sessions"`
}

func (u *User) GetClaims() map[string]interface{} {
	return map[string]interface{}{
		"id":      u.ID.String(),
		"sub":     u.Sub,
		"email":   u.Email,
		"name":    u.Name,
		"picture": u.Picture,
	}
}

func NewUserFromClaims(claims map[string]interface{}) (user User, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = r.(error)
		}
	}()

	sub := claims["sub"].(string)
	email := claims["email"].(string)
	name := claims["name"].(string)
	picture := claims["picture"].(string)

	return User{
		Sub:     sub,
		Email:   email,
		Name:    name,
		Picture: picture,
	}, err
}

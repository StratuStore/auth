package core

import (
	"errors"
	"github.com/google/uuid"
)

type User struct {
	ID      uuid.UUID `json:"id" bson:"_id,omitempty"`
	Sub     string    `json:"sub" bson:"sub"`
	Email   string    `json:"email" bson:"email"`
	Name    string    `json:"name" bson:"name"`
	Picture string    `json:"picturePath" bson:"picturePath"`
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
	sub, ok := claims["sub"].(string)
	if !ok {
		return User{}, errors.New("unable to parse sub")
	}
	email, ok := claims["email"].(string)
	if !ok {
		return User{}, errors.New("unable to parse email")
	}
	name, ok := claims["name"].(string)
	if !ok {
		return User{}, errors.New("unable to parse name")
	}
	picture, ok := claims["picture"].(string)
	if !ok {
		return User{}, errors.New("unable to parse picture")
	}

	return User{
		Sub:     sub,
		Email:   email,
		Name:    name,
		Picture: picture,
	}, err
}

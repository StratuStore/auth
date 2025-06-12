package core

import (
	"errors"
	"github.com/mbretter/go-mongodb/types"
)

type User struct {
	ID      types.ObjectId `json:"id" bson:"_id,omitempty"`
	Email   string         `json:"email" bson:"email"`
	Name    string         `json:"name" bson:"name"`
	Picture string         `json:"picturePath" bson:"picturePath"`
}

func (u *User) GetClaims() map[string]interface{} {
	return map[string]interface{}{
		"id":      string(u.ID),
		"email":   u.Email,
		"name":    u.Name,
		"picture": u.Picture,
	}
}

func NewUserFromClaims(claims map[string]interface{}) (user User, err error) {
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
		Email:   email,
		Name:    name,
		Picture: picture,
	}, err
}

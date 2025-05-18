package core

import (
	"fmt"
	"github.com/google/uuid"
)

type Session struct {
	ID         string    `json:"id" bson:"_id,omitempty"`
	UserSub    string    `json:"userSub"`
	Salt       uuid.UUID `json:"salt" bson:"salt"`
	DeviceData string    `json:"deviceData"`
}

func (session *Session) GetClaims() map[string]interface{} {
	return map[string]interface{}{
		"id":         session.ID,
		"userSub":    session.UserSub,
		"deviceData": session.DeviceData,
		"salt":       session.Salt.String(),
	}
}

func NewSessionFromClaims(claims map[string]interface{}) (session Session, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = r.(error)
		}
	}()

	idStr, _ := claims["id"].(string)
	userSub := claims["userSub"].(string)
	deviceData := claims["deviceData"].(string)
	saltString, ok := claims["salt"].(string)
	if !ok {
		return Session{}, fmt.Errorf("unable to process salt")
	}
	salt, err := uuid.Parse(saltString)
	if err != nil {
		return Session{}, fmt.Errorf("salt must be uuid")
	}

	return Session{
		ID:         idStr,
		UserSub:    userSub,
		DeviceData: deviceData,
		Salt:       salt,
	}, err
}

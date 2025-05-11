package core

import (
	"fmt"
	"github.com/google/uuid"
)

type Session struct {
	ID           uuid.UUID `json:"id" bson:"_id,omitempty"`
	UserSub      string    `json:"userSub"`
	RefreshToken string    `json:"-" bson:"refreshToken"`
	DeviceData   string    `json:"deviceData"`
}

func (session *Session) GetClaims() map[string]interface{} {
	return map[string]interface{}{
		"id":         session.ID.String(),
		"userSub":    session.UserSub,
		"deviceData": session.DeviceData,
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

	id, err := uuid.Parse(idStr)
	if err != nil {
		return Session{}, fmt.Errorf("failed to parse session id: %w", err)
	}

	return Session{
		ID:         id,
		UserSub:    userSub,
		DeviceData: deviceData,
	}, err
}

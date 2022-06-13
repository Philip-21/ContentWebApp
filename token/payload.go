package token

import (
	"errors"
	"time"

	"github.com/gofrs/uuid"
)

type Payload struct {
	ID        uuid.UUID `json:"id"`
	Username  string    `json:"username"` //this identifs the token owner
	IssuedAt  time.Time `json:"issued_at"`
	ExpiredAt time.Time `json:"expired_at"`
}

//NewPayload() creates a new token payload with a specific username and duration.
func NewPayload(username string, duration time.Duration) (*Payload, error) {
	tokenID, err := uuid.DefaultGenerator.NewV4() // generate a unique token ID.
	if err != nil {
		return nil, err
	}

	payload := &Payload{
		ID:        tokenID,
		Username:  username,
		IssuedAt:  time.Now(),
		ExpiredAt: time.Now().Add(duration),
	}
	return payload, nil
}

var ErrExpiredToken = errors.New("token has expired")

//expiration ofthe payload token
func (payload *Payload) Valid() error {
	//If time.Now() is after the payload.ExpiredAt,
	//then it means that the token has expired
	if time.Now().After(payload.ExpiredAt) {
		return ErrExpiredToken
	}
	return nil
}

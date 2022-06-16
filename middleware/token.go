package middleware

import (
	"errors"
	"fmt"
	"time"

	"github.com/aead/chacha20poly1305"
	"github.com/gofrs/uuid"
	"github.com/o1egl/paseto"
)

type Payload struct {
	ID        uuid.UUID `json:"id"`
	Email     string    `json:"email"` //this identifs the token owner
	IssuedAt  time.Time `json:"issued_at"`
	ExpiredAt time.Time `json:"expired_at"`
}

//NewPayload() creates a new token payload with a specific username and duration.
func NewPayload(email string, duration time.Duration) (*Payload, error) {
	tokenID, err := uuid.DefaultGenerator.NewV4() // generate a unique token ID.
	if err != nil {
		return nil, err
	}

	payload := &Payload{
		ID:        tokenID,
		Email:     email,
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

type PasetoMaker struct {
	paseto       *paseto.V2
	symmetricKey []byte
}

// NewPasetoMaker(), which takes a symmetricKey string as input,
//and returns a token.Maker interface or an error.
//this function creates a new PasetoMaker instance
func NewPasetoMaker(symmetricKey string) (Maker, error) {
	if len(symmetricKey) != chacha20poly1305.KeySize { //Paseto version 2 uses Chacha20 Poly1305 algorithm to encrypt the payload
		return nil, fmt.Errorf("invalid key size: must be exactly %d characters", chacha20poly1305.KeySize)
	} //If the key length is not correct then we just return a nil object and an error saying invalid key size

	maker := &PasetoMaker{
		paseto:       paseto.NewV2(),
		symmetricKey: []byte(symmetricKey),
	}

	return maker, nil
}

//create new paseto token
func (maker *PasetoMaker) CreateToken(email string, duration time.Duration) (string, error) {
	payload, err := NewPayload(email, duration)
	if err != nil {
		return "", err
	}

	return maker.paseto.Encrypt(maker.symmetricKey, payload, nil)
}

//PASETO VerifyToken method
func (maker *PasetoMaker) VerifyToken(token string) (*Payload, error) {
	//declaring an empty payload object to store the decrypted data
	payload := &Payload{}

	err := maker.paseto.Decrypt(token, maker.symmetricKey, payload, nil)
	if err != nil {
		return nil, err
	}

	err = payload.Valid()
	if err != nil {
		return nil, err
	}

	return payload, nil
}

// a token Maker interface to manage the creation and verification of the Paseto tokens
type Maker interface {
	CreateToken(email string, duration time.Duration) (string, error) //create and sign a new token for a specific username and valid duration.
	VerifyToken(token string) (*Payload, error)                       //checks if the input token is valid or not.
}

//makes routes secure by only allowing logged in users to have access to certain parts,pages(routes) of the application
// func Auth (next  http.Handler)http.Handler{

// }

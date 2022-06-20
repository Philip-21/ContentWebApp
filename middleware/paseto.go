package middleware

import (
	"fmt"
	"time"

	"github.com/aead/chacha20poly1305"
	"github.com/o1egl/paseto"
)

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

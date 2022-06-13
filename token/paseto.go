package token

import (
	"fmt"
	"time"

	"github.com/aead/chacha20poly1305"
	"github.com/o1egl/paseto"
)

// a token Maker interface to manage the creation and verification of tokens
type Maker interface {
	CreateToken(username string, duration time.Duration) (string, error) //create and sign a new token for a specific username and valid duration.
	VerifyToken(token string) (*Payload, error)                          //checks if the input token is valid or not.
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
func (maker *PasetoMaker) CreateToken(username string, duration time.Duration) (string, error) {
	payload, err := NewPayload(username, duration)
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

// //the function to create the response object
// func (server *Server) Createuser(ctx *gin.Context) {

// 	//user, err := server.store.
// 	user, err := server.store.

// 	rsp := models.NewUserResponse()
// 	ctx.JSON(http.StatusOK, rsp)
// }

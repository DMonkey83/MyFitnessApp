package token

import (
	"crypto/ed25519"
	"encoding/hex"
	"log"
	"time"

	"github.com/o1egl/paseto"
)

// PasetoMaker is a PASETO token maker
type PasetoMaker struct {
	paseto    *paseto.V2
	publicKey ed25519.PublicKey
}

// NewPasetoMaker creates a new PasetoMaker
func NewPasetoMaker(key []byte) (Maker, error) {
	key, _ = hex.DecodeString("1eb9dbbbbc047c03fd70604e0071f0987e16b28b757225c11f00415d0e20b1a2")
	publicKey := ed25519.PublicKey(key)

	maker := &PasetoMaker{
		paseto:    paseto.NewV2(),
		publicKey: ed25519.PublicKey(publicKey),
	}

	return maker, nil
}

// CreateToken creates a new token for a specific username and duration
func (maker *PasetoMaker) CreateToken(username string, duration time.Duration) (string, *Payload, error) {
	payload, err := NewPayload(username, duration)
	if err != nil {
		return "", payload, err
	}
	b, _ := hex.DecodeString("b4cbfb43df4ce210727d953e4a713307fa19bb7d9f85041438d9e11b942a37741eb9dbbbbc047c03fd70604e0071f0987e16b28b757225c11f00415d0e20b1a2")
	privateKey := ed25519.PrivateKey(b)

	footer := "new footer"

	token, err := maker.paseto.Sign(privateKey, payload, footer)
	if err != nil {
		return "", payload, err

	}
	return token, payload, err
}

// VerifyToken checks if the token is valid or not
func (maker *PasetoMaker) VerifyToken(token string) (*Payload, error) {
	payload := &Payload{}

	var newFooter string

	log.Print("maker", maker.publicKey, payload)

	err := maker.paseto.Verify(token, maker.publicKey, &payload, &newFooter)
	if err != nil {
		return nil, ErrInvalidToken
	}

	err = payload.Valid()
	if err != nil {
		return nil, err
	}

	return payload, nil
}

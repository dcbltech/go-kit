package jwt

import (
	"time"

	"github.com/square/go-jose/v3"
	"github.com/square/go-jose/v3/jwt"
)

func Sign(key, id string) (string, error) {
	expiry := time.Now().AddDate(0, 1, 0)

	claims := &jwt.Claims{
		ID:     id,
		Expiry: jwt.NewNumericDate(expiry),
	}

	return SignWith(key, claims)
}

func SignWith(key string, claims *jwt.Claims) (string, error) {
	sig, err := jose.NewSigner(jose.SigningKey{Algorithm: jose.HS256, Key: []byte(key)}, (&jose.SignerOptions{}).WithType("JWT"))
	if err != nil {
		return "", err
	}

	token, err := jwt.Signed(sig).Claims(claims).CompactSerialize()
	if err != nil {
		return "", err
	}

	return token, nil
}

func Parse(key, raw string) (claims *jwt.Claims, err error) {
	k := []byte(key)

	token, err := jwt.ParseSigned(raw)
	if err != nil {
		return nil, err
	}

	claims = &jwt.Claims{}
	if err = token.Claims(k, claims); err != nil {
		return
	}

	return
}

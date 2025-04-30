package jwt

import (
	"testing"
	"time"

	"github.com/square/go-jose/v3/jwt"
	"github.com/stretchr/testify/assert"
)

func TestSign(t *testing.T) {
	key := "test-secret-key"
	id := "test-id"

	token, err := Sign(key, id)

	assert.NoError(t, err, "Sign() returned an error")
	assert.NotEmpty(t, token, "Sign() returned an empty token")
}

func TestSignWith(t *testing.T) {
	key := "test-secret-key"
	expiry := time.Now().Add(24 * time.Hour)
	claims := &jwt.Claims{
		ID:     "test-id",
		Expiry: jwt.NewNumericDate(expiry),
	}

	token, err := SignWith(key, claims)

	assert.NoError(t, err, "SignWith() returned an error")
	assert.NotEmpty(t, token, "SignWith() returned an empty token")
}

func TestParse(t *testing.T) {
	key := "test-secret-key"
	id := "test-id"

	token, err := Sign(key, id)
	assert.NoError(t, err, "Sign() returned an error")

	claims, err := Parse(key, token)
	assert.NoError(t, err, "Parse() returned an error")
	assert.Equal(t, id, claims.ID, "Parse() returned claims with incorrect ID")
}

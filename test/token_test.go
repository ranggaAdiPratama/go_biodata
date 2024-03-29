package test

import (
	"errors"
	"testing"
	"time"

	"github.com/ranggaAdiPratama/go_biodata/token"
	"github.com/ranggaAdiPratama/go_biodata/util"
	"github.com/stretchr/testify/require"
)

var (
	ErrExpiredToken = errors.New("token has expired")
	ErrInvalidToken = errors.New("token is invalid")
)

func TestPasetoMaker(t *testing.T) {
	maker, err := token.NewPasetoMaker(util.RandomString(32))
	require.NoError(t, err)

	username := util.RandomUsername()
	duration := time.Minute

	issuedAt := time.Now()
	expiredAt := issuedAt.Add(duration)

	token, payload, err := maker.CreateToken(username, 1, duration)
	require.NoError(t, err)
	require.NotEmpty(t, token)
	require.NotEmpty(t, payload)

	payload, err = maker.VerifyToken(token)
	require.NoError(t, err)
	require.NotEmpty(t, token)

	require.NotZero(t, payload.ID)
	require.Equal(t, username, payload.Username)
	require.WithinDuration(t, issuedAt, payload.IssuedAt, time.Second)
	require.WithinDuration(t, expiredAt, payload.ExpiredAt, time.Second)
}

func TestExpiredPasetoToken(t *testing.T) {
	maker, err := token.NewPasetoMaker(util.RandomString(32))

	require.NoError(t, err)

	token, payload, err := maker.CreateToken(util.RandomUsername(), 1, -time.Minute)

	require.NoError(t, err)
	require.NotEmpty(t, token)

	payload, err = maker.VerifyToken(token)

	require.Error(t, err)
	require.EqualError(t, err, ErrExpiredToken.Error())
	require.Nil(t, payload)
}

func TestInvalidPasetoTokenAlgNone(t *testing.T) {
	maker, err := token.NewPasetoMaker(util.RandomString(32))

	require.NoError(t, err)

	token1, payload1, err := maker.CreateToken(util.RandomUsername(), 1, time.Minute)

	require.NoError(t, err)
	require.NotEmpty(t, token1)
	require.NotEmpty(t, payload1)

	maker1, err := token.NewPasetoMaker(util.RandomString(32))

	require.NoError(t, err)

	token2, payload2, err := maker1.CreateToken(util.RandomUsername(), 1, time.Minute)

	require.NoError(t, err)
	require.NotEmpty(t, payload2)

	payload, err := maker.VerifyToken(token2)

	require.Error(t, err)
	require.EqualError(t, err, ErrInvalidToken.Error())
	require.Nil(t, payload)
}

package token

import (
	"errors"
	"strconv"
	"time"

	"github.com/hako/branca"
)

var (
	DefaultToken *Token
)

type Token struct {
	b *branca.Branca
}

func New(key string, expire time.Duration) (*Token, error) {
	if len(key) != 32 {
		return nil, errors.New("invalid key")
	}
	b := branca.NewBranca(key)
	b.SetTTL(uint32(expire.Milliseconds()))
	return &Token{
		b: b,
	}, nil
}

func (t *Token) Encode(userId int64) (string, error) {
	return t.b.EncodeToString(strconv.FormatInt(userId, 10))
}

func (t *Token) Decode(str string) (int64, error) {
	s, err := t.b.DecodeToString(str)
	if err != nil {
		return 0, err
	}
	return strconv.ParseInt(s, 10, 64)
}

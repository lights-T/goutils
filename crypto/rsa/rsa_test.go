package rsa

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	privKey = `MIICdwIBADANBgkqhkiG9w0BAQEFAASCAmEwggJdAgEAAoGBAMqoAuswLprgCxW5YEUDEdqk3PudIuV/Wpqcju7f9gcvNJH+s9wOC2Zs+GwZ7QEm26LCtwUGX10U9TUCIlLQpDhWHM40+77mhuWA4uEofg6zTOCUsaFyEZm3xzTuxjbVzlGXNAzuLNRE8EqzEU4I8AakPvyY6P2lmdmgRvRQDVAfAgMBAAECgYBICKZY8NxwApkOFMFqZmfvPtCpwzYHO1h6QpHvyL3L2fSmvFE0M+3Lb4px6lk7IpPJa8rgR16YWH28ZNDMfQsRtIdo5PZtYbeFJs5h7zGuLh9Asr43a4tPaLihHJKieQeyku3R2mlkMf8jRU40a9npoKTAGSrYwU/vQNb7w2mgoQJBAPyaBz88DWY5EsZ1mqhbc6DXfVvCHjpHtQN5WC96rE/fPqFJmaHYhOYINd+aqwusGc33m8RH7c1tVCuIlDvtydECQQDNYfgLvwe77Er+WNOqQxj4RA7KTi9nXXXuE0/qMJgqMRptM8vtAP36dYNXppNn1nFR80dQXD8hY6Y6qWD24wbvAkEAkqdWZ8cUvHGMTf5/YRlfQ1V4qWpFJG73T+IGaeJd4i1pbjiN4qITXn4L0Rs6DRfJD4SfQdDE5ox/3pp3/Wcr0QJAJFpC3VFivRCF9Z8jV++oa8kgFQ7htRoF1a31Zy5SwKUQWGPipICYc8x5Avqo/KgoRqkY5lBtnCtXMOrqDskAYQJBAJJGxH0Mue7HNel917/7goBsFXqcedzM2v2GLzlro/XegptkeRYfryiAXb1OKelYc6lTx+cGDl6nWr91G+LfZro=`
	pubKey  = `MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQDKqALrMC6a4AsVuWBFAxHapNz7nSLlf1qanI7u3/YHLzSR/rPcDgtmbPhsGe0BJtuiwrcFBl9dFPU1AiJS0KQ4VhzONPu+5oblgOLhKH4Os0zglLGhchGZt8c07sY21c5RlzQM7izURPBKsxFOCPAGpD78mOj9pZnZoEb0UA1QHwIDAQAB`

	r *Rsa

	text        = "加餐聊顿觉，密意未全消，解榻情何限，密疏坐寂寥。"
	cryptedText = "UotCVgUVuMR4+nYVF7/tF0/i1DcfZyYn43FTzqUDzcJZyvKhuH7bNEBuoYt29P16/MsVXJQzhVBmJYGvvaVGih0FULf2KzzuaLfg2CO+CxsDFvX5y3t8VS/qsZcx0gJIBy/WbH52WCq5jevg3BArVClZEeTWz8yrEQvFNkClc3A="
	sign        = "ADxvkSqKPrrjloyiwRUna4zRcY9AmfaAyWla1K7SrS20XT3V4nTxP9z4JD5WGDq0EGley/9loHvN1nW9E72CJ7ZaCTongImkzO3k+GzxhmIjCCF6y1/H3DWjPpXPL81Qtzg+S4B4h89kEcuTpHK9uzx3OAuB2TYzmiKhRGcZg7s="
)

func init() {
	var err error
	if r, err = New(pubKey, privKey); err != nil {
		log.Fatal(err)
	}
}

func TestGenerateKey(t *testing.T) {
	err := GenerateKey(1024)
	assert.NoError(t, err)
}

func TestNew(t *testing.T) {
	r, err := New(pubKey, privKey)
	assert.NoError(t, err)
	assert.NotNil(t, r)
}

func TestRsa_Encrypt(t *testing.T) {
	result, err := r.Encrypt(text)
	assert.NoError(t, err)
	assert.NotEmpty(t, result)
}

func BenchmarkRsa_Encrypt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := r.Encrypt(text)
		assert.NoError(b, err)
	}
}

func TestRsa_Decrypt(t *testing.T) {
	result, err := r.Decrypt(cryptedText)
	assert.NoError(t, err)
	assert.EqualValues(t, text, result)
}

func BenchmarkRsa_Decrypt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := r.Decrypt(cryptedText)
		assert.NoError(b, err)
	}
}

func TestRsa_Sign(t *testing.T) {
	sign, err := r.Sign(text)
	assert.NoError(t, err)
	assert.NotEmpty(t, sign)
}

func BenchmarkRsa_Sign(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := r.Sign(text)
		assert.NoError(b, err)
	}
}

func TestRsa_Verify(t *testing.T) {
	err := r.Verify(text, sign)
	assert.NoError(t, err)
}

func BenchmarkRsa_Verify(b *testing.B) {
	for i := 0; i < b.N; i++ {
		err := r.Verify(text, sign)
		assert.NoError(b, err)
	}
}

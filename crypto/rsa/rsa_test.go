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

//正确流程：签名 + 传输 + 验证
//发送方：
//1. 原始数据 → 计算哈希 → 用私钥签名 → 得到签名
//2. 发送：原始数据 + Base64 编码的签名
//
//接收方：
//1. 收到：原始数据 + 签名（Base64 解码）
//2. 对收到的数据计算哈希
//3. 用公钥验证：这个哈希 + 签名 是否匹配
//4. 匹配 → 数据可信；不匹配 → 被篡改或来源非法
func Test_Auth(t *testing.T) {
	privateKey := "MIICdwIBADANBgkqhkiG9w0BAQEFAASCAmEwggJdAgEAAoGBAML8rhqz2tTzANW+UGXsdkTJNhS8jpLqk+SyqrsD/bkqiCbSG6tbeHbJcLKecxQlhSMdCAydmU50JK2THtzb648S82MAXz/JKBGC1feRLstBFqdZcS4a6TxFaJnpSaNTB4jFBGgGcM/91mzDz0v58GINfGrhmsa5fRldwRmPYnNRAgMBAAECgYBYuJiP1d5wntF2cE4s0ldOHS/aZ6GH/+yjVxiQV9SO+GdTIq8sXUaG5km9PJOoSxo1S/RpqRwksnwt7o9Qd1DK0D3iNEjBni6ywzm6P33swcwSrx5dRuYwpeR7L3KJQZkS2ugrX99ZoABwemWGFj6U3rDnV0mFPI6WMMzQ3KCvMQJBAMWbVqKgFfa8mNnd2v1s/risrBeEz4F2NY/rm9D+KG+BXxVhwgpc5XwFC8zympOs1AHBhudskid840ptEaM7rqsCQQD8mykoh5YABRAu43YFUNrJPVYOyHJ4l1OpqnLZcGOJbpFbVy9pFPe7uw3WtAfvyY4AJHd2lapWJhkqwZf2tvXzAkAKsR786aCGmynCEAj7UVxu7ZjaJOt9W8IGKX9izX2umtdkNsfi+6fHEBbVXgMTHnTSK4B7IRq/XDiIHGKp7F7FAkEAs+l63f/7qNXyWcLtqwmUWiISagL/7L2y+7OHizCN5DNY2dp1zPz/GLk4OQQOZw2B0r4mS9J6+FK4OAicSD61WwJBAKEPBBXNOtciODI4T4ncpsX2XlR/cSp2jWH+e25qkTcy8XGDyiKn3WQCWj1rU4nLJYfx25TlrqbPwbDcSkEEgiU="
	publicKey := "MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQDC/K4as9rU8wDVvlBl7HZEyTYUvI6S6pPksqq7A/25Kogm0hurW3h2yXCynnMUJYUjHQgMnZlOdCStkx7c2+uPEvNjAF8/ySgRgtX3kS7LQRanWXEuGuk8RWiZ6UmjUweIxQRoBnDP/dZsw89L+fBiDXxq4ZrGuX0ZXcEZj2JzUQIDAQAB"

	originalData := []byte("用户ID: 12345, 操作: 转账100元, 时间: 2025-04-05")
	t.Run("Signature", func(t *testing.T) {
		// ====================
		// 1. 模拟：生成密钥对
		// ====================
		//privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
		//if err != nil {
		//	log.Fatal("生成密钥失败:", err)
		//}
		//publicKey := &privateKey.PublicKey

		// ====================
		// 2. 发送方：签名
		// ====================
		//originalData := []byte("用户ID: 12345, 操作: 转账100元, 时间: 2025-04-05")

		// 私钥签名
		encodedSignature, err := SignPKCS1v15(privateKey, originalData)
		if err != nil {
			t.Fatal("签名失败: ", err)
		}
		t.Logf("编码后的签名: %s", encodedSignature)
		t.Logf("编码后的数据: %s", originalData)
	})

	t.Run("Verify", func(t *testing.T) {
		encodedSignature := "dkXsf9FiXIVHJGiwtcFSCj9hqQQkcQpK/iCQl/1D974TzwoQ1wTjw2GEFIlyqziSKORCztIHbTGS4kwbXjMOkIcUFdD8ZNBZlPWqdz6ilj/2OwDPnZtqFvMqjDCfD4cnyEmntPUimE84+1u1/KpuNari0q4sLlBWU3EFpdN4iJo="

		// 3.3 使用公钥验证签名
		err := VerifyPKCS1v15(publicKey, encodedSignature, originalData)
		if err != nil {
			t.Fatal("❌ 签名验证失败：数据可能被篡改或来源非法")
			return
		}
		t.Log("验签成功")
	})
}

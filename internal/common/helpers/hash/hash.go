package hash

import (
	"crypto/hmac"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/sha512"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"os"

	"go.uber.org/zap"
)

// Encode рассчет хеша SHA256
func Encode(data *[]byte, key string) string {
	h := hmac.New(sha256.New, []byte(key))
	h.Write(*data)
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}

type RSADecrypt struct {
	private *rsa.PrivateKey
}

func (re RSADecrypt) Decrypt(ciphertext []byte) ([]byte, error) {
	hash := sha512.New()

	decode, err := base64.StdEncoding.DecodeString(string(ciphertext))
	if err != nil {
		return nil, err
	}

	return rsa.DecryptOAEP(hash, rand.Reader, re.private, decode, nil)
}

func NewRSARSADecrypt(privatePath string, log *zap.SugaredLogger) *RSADecrypt {
	bytes, err := os.ReadFile(privatePath)
	if err != nil {
		log.Fatal("failed to read private key", zap.Error(err))
	}

	block, _ := pem.Decode(bytes)
	private, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		log.Fatal("failed to parse private key", zap.Error(err))
	}
	rsaKey, ok := private.(*rsa.PrivateKey)
	if !ok {
		log.Fatal("failed to parse private key")
	}

	return &RSADecrypt{
		rsaKey,
	}
}

type RSAEncrypt struct {
	public *rsa.PublicKey
}

func (re RSAEncrypt) Encrypt(msg []byte) ([]byte, error) {
	hash := sha512.New()
	data, err := rsa.EncryptOAEP(hash, rand.Reader, re.public, msg, nil)
	if err != nil {
		return nil, err
	}
	encode := base64.StdEncoding.EncodeToString(data)
	return []byte(encode), nil
}

func NewRSAEncrypt(publicPath string, log *zap.SugaredLogger) *RSAEncrypt {
	bytes, err := os.ReadFile(publicPath)
	if err != nil {
		log.Fatal("failed to read public key", zap.Error(err))
	}

	block, _ := pem.Decode(bytes)
	public, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		log.Fatal("failed to parse public key", zap.Error(err))
	}
	rsaKey, ok := public.(*rsa.PublicKey)
	if !ok {
		log.Fatal("failed to parse public key")
	}

	return &RSAEncrypt{
		rsaKey,
	}
}

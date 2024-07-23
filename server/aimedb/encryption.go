package aimedb

import (
	"crypto/aes"
	"crypto/cipher"
)

const aesKey = "Copyright(C)SEGA"

// ecbEncrypter implements ECB encryption
type ecbEncrypter struct {
	b         cipher.Block
	blockSize int
}

// NewECBEncrypter creates a new ECB encrypter
func NewECBEncrypter(b cipher.Block) cipher.BlockMode {
	return &ecbEncrypter{
		b:         b,
		blockSize: b.BlockSize(),
	}
}

func (x *ecbEncrypter) BlockSize() int { return x.blockSize }
func (x *ecbEncrypter) CryptBlocks(dst, src []byte) {
	if len(src)%x.blockSize != 0 {
		panic("ecbEncrypter: input not full blocks")
	}
	if len(dst) < len(src) {
		panic("ecbEncrypter: output smaller than input")
	}
	for len(src) > 0 {
		x.b.Encrypt(dst, src[:x.blockSize])
		src = src[x.blockSize:]
		dst = dst[x.blockSize:]
	}
}

// ecbDecrypter implements ECB decryption
type ecbDecrypter struct {
	b         cipher.Block
	blockSize int
}

// NewECBDecrypter creates a new ECB decrypter
func NewECBDecrypter(b cipher.Block) cipher.BlockMode {
	return &ecbDecrypter{
		b:         b,
		blockSize: b.BlockSize(),
	}
}

func (x *ecbDecrypter) BlockSize() int { return x.blockSize }
func (x *ecbDecrypter) CryptBlocks(dst, src []byte) {
	if len(src)%x.blockSize != 0 {
		panic("ecbDecrypter: input not full blocks")
	}
	if len(dst) < len(src) {
		panic("ecbDecrypter: output smaller than input")
	}
	for len(src) > 0 {
		x.b.Decrypt(dst, src[:x.blockSize])
		src = src[x.blockSize:]
		dst = dst[x.blockSize:]
	}
}

// Encrypt encrypts the data using AES encryption in ECB mode
func Encrypt(data []byte) ([]byte, error) {
	block, err := aes.NewCipher([]byte(aesKey))
	if err != nil {
		return nil, err
	}
	ecb := NewECBEncrypter(block)
	ciphertext := make([]byte, len(data))
	ecb.CryptBlocks(ciphertext, data)
	return ciphertext, nil
}

// Decrypt decrypts the data using AES encryption in ECB mode
func Decrypt(data []byte) ([]byte, error) {
	block, err := aes.NewCipher([]byte(aesKey))
	if err != nil {
		return nil, err
	}
	ecb := NewECBDecrypter(block)
	plaintext := make([]byte, len(data))
	ecb.CryptBlocks(plaintext, data)
	return plaintext, nil
}

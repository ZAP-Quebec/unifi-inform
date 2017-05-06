package inform

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"
)

type IV Key

type Key []byte

func (k Key) IsValid() bool {
	return len(k) == blockSize
}

func (k Key) IsDefault() bool {
	return bytes.Equal(k, DEFAULT_KEY)
}

func (k Key) String() string {
	return hex.EncodeToString(k)
}

var (
	DEFAULT_KEY = Key([]byte{
		0xba, 0x86, 0xf2, 0xbb,
		0xe1, 0x07, 0xc7, 0xc5,
		0x7e, 0xb5, 0xf2, 0x69,
		0x07, 0x75, 0xc7, 0x12,
	})
)

const (
	blockSize int = 16 // bytes
)

func GenerateIV() (IV, error) {
	iv := make([]byte, blockSize)
	if n, err := rand.Read(iv); err != nil {
		return nil, err
	} else if n != blockSize {
		return nil, errors.New("Could not get enough randomness to generate IV")
	}
	return iv, nil
}

func Encrypt(iv IV, key Key, data []byte) (result []byte, err error) {
	if err = assertCryptoParams(iv, key); err != nil {
		return
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return
	}

	srcLen := len(data)
	padLen := blockSize - (srcLen % blockSize)
	srcBuf := make([]byte, len(data)+padLen)
	dstBuf := make([]byte, len(data)+padLen)

	copy(srcBuf, data)
	padding := bytes.Repeat([]byte{byte(padLen)}, padLen)
	copy(srcBuf[srcLen:], padding)

	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(dstBuf, srcBuf)

	return dstBuf, nil
}

func Decrypt(iv, key, data []byte) (result []byte, err error) {
	if err = assertCryptoParams(iv, key); err != nil {
		return
	}

	if len(data)%blockSize != 0 {
		return nil, fmt.Errorf("encrypted data must be a multiple of %d bytes", blockSize)
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return
	}

	dataLen := len(data)
	result = make([]byte, dataLen)
	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(result, data)

	padLen := int(result[dataLen-1])
	if padLen > blockSize {
		return nil, fmt.Errorf("Invalid padding: %d > %d (blockSize)", padLen, blockSize)
	}

	return result[:dataLen-padLen], nil
}

func assertCryptoParams(iv IV, key Key) error {
	if !Key(iv).IsValid() {
		return fmt.Errorf("iv length must be %d bytes [len(iv)==%d]", blockSize, len(iv))
	}
	if !key.IsValid() {
		return fmt.Errorf("key length must be %d bytes [len(key)==%d]", blockSize, len(key))
	}
	return nil
}

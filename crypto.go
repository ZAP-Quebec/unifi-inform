package inform

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"errors"
	"fmt"
	"github.com/ZAP-Quebec/unifi-inform/data"
)

type IV data.Key

const (
	blockSize int = data.KEY_SIZE // bytes
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

func Encrypt(iv IV, key data.Key, data []byte) (result []byte, err error) {
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

func Decrypt(iv IV, key data.Key, data []byte) (result []byte, err error) {
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

func assertCryptoParams(iv IV, key data.Key) error {
	if !data.Key(iv).IsValid() {
		return fmt.Errorf("iv length must be %d bytes [len(iv)==%d]", blockSize, len(iv))
	}
	if !key.IsValid() {
		return fmt.Errorf("key length must be %d bytes [len(key)==%d]", blockSize, len(key))
	}
	return nil
}

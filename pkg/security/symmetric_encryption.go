package security

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
)

// AesEncrypt encrypts plain text string into cipher text string
func AesEncrypt(str string) (string, error) {
	key := []byte("0123456789abcdef0123456789abcdef")
	plainText := []byte(str)
	plainText, err := pkcs7Pad(plainText, aes.BlockSize)
	if err != nil {
		return "", fmt.Errorf(`PlainText: "%s" has error`, plainText)
	}
	
	if len(plainText)%aes.BlockSize != 0 {
		err := fmt.Errorf(`PlainText: "%s" has the wrong block size`, plainText)
		return "", err
	}
	
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	
	cipherText := make([]byte, aes.BlockSize+len(plainText))
	iv := cipherText[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}
	
	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(cipherText[aes.BlockSize:], plainText)
	
	return fmt.Sprintf("%x", cipherText), nil
}

// AesDecrypt decrypts cipher text string into plain text string
func AesDecrypt(str string) (string, error) {
	key := []byte("0123456789abcdef0123456789abcdef")
	cipherText, err := hex.DecodeString(str)
	if err != nil {
		return "", err
	}
	
	if len(cipherText) < aes.BlockSize {
		return "", fmt.Errorf("CipherText too short")
	}
	
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	
	// iv := []byte(_config.Cfg.AesIv)[:aes.BlockSize]
	// cipherText = cipherText[:aes.BlockSize]
	iv := []byte("0123456789ABCDEF")[:aes.BlockSize]
	
	if len(cipherText)%aes.BlockSize != 0 {
		return "", fmt.Errorf("CipherText is not a multiple of the block size")
	}
	
	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(cipherText, cipherText)
	
	cipherText, err = pkcs7Unpad(cipherText, aes.BlockSize)
	if err != nil {
		return "", err
	}
	
	return fmt.Sprintf("%s", cipherText), nil
}

func pkcs7Pad(buf []byte, size int) ([]byte, error) {
	bufLen := len(buf)
	padLen := size - bufLen%size
	padded := make([]byte, bufLen+padLen)
	copy(padded, buf)
	
	for i := 0; i < padLen; i++ {
		padded[bufLen+i] = byte(padLen)
	}
	
	return padded, nil
}

func pkcs7Unpad(padded []byte, size int) ([]byte, error) {
	if len(padded)%size != 0 {
		return nil, fmt.Errorf("pkcs7: Padded value wasn't in correct size")
	}
	
	bufLen := len(padded) - int(padded[len(padded)-1])
	
	if bufLen < 1 {
		return nil, fmt.Errorf("pkcs7: Padded value wasn't in correct length")
	}
	
	buf := make([]byte, bufLen)
	copy(buf, padded[:bufLen])
	
	return buf, nil
}

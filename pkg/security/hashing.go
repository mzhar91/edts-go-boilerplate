package security

import (
	"crypto/sha512"
	"encoding/base64"
	"math/rand"
	"strconv"
	"time"
)

// Define the size of the salt
const saltSize = 16

// GenerateRandomSalt Generate a random 16 bytes securely using the
// Cryptographically secure pseudorandom number generator (CSPRNG)
// int the crypto.rand package
func GenerateRandomSalt() []byte {
	var salt = make([]byte, saltSize)
	
	_, err := rand.Read(salt[:])
	if err != nil {
		panic(err)
	}
	
	return salt
}

// GenerateSalt Generate a random 16 bytes securely based on email
func GenerateSalt(str string) []byte {
	b := make([]byte, saltSize)
	for i := range b {
		b[i] = str[rand.Intn(len(str))]
	}
	
	return b
}

// Hash Combine our password and salt and hash them using the SHA-512
// hashing algorithm and then return our hashed password
// as a base64 encoded string
func Hash(str string, salt []byte) string {
	// Convert password string to byte slice
	var strBytes = []byte(str)
	
	// Create sha-512 hasher
	var sha512Hasher = sha512.New()
	
	// Append salt to password
	strBytes = append(strBytes, salt...)
	
	// Write password bytes to the hasher
	sha512Hasher.Write(strBytes)
	
	// Get the sha512 hashed password
	var hashedPasswordBytes = sha512Hasher.Sum(nil)
	
	// Convert the hashed password to a base64 encoded string
	var base64EncodedPasswordHash = base64.URLEncoding.EncodeToString(hashedPasswordBytes)
	
	return base64EncodedPasswordHash
}

// HashWithTime Combine our password and salt and hash them using the SHA-512
// hashing algorithm and then return our hashed password
// as a base64 encoded string
func HashWithTime(str string, salt []byte) string {
	// Add time now into string to make more unique salt
	str = str + strconv.Itoa(time.Now().Hour()) + strconv.Itoa(time.Now().YearDay())
	
	// Convert password string to byte slice
	var strBytes = []byte(str)
	
	// Create sha-512 hasher
	var sha512Hasher = sha512.New()
	
	// Append salt to password
	strBytes = append(strBytes, salt...)
	
	// Write password bytes to the hasher
	sha512Hasher.Write(strBytes)
	
	// Get the sha512 hashed password
	var hashedPasswordBytes = sha512Hasher.Sum(nil)
	
	// Convert the hashed password to a base64 encoded string
	var base64EncodedPasswordHash = base64.URLEncoding.EncodeToString(hashedPasswordBytes)
	
	return base64EncodedPasswordHash
}

// DoMatch Check if two passwords match
func DoMatch(hashed, str string, salt []byte) bool {
	var strHash = Hash(str, salt)
	
	return hashed == strHash
}

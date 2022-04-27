package hashpaww

import (
	"crypto/rand"

	"crypto/sha512"
	"encoding/base64"
)

func generateRandomSalt(saltSize int) []byte {
	var salt = make([]byte, saltSize)
	_, err := rand.Read(salt[:])
	if err != nil {
		panic(err)

	}
	return salt
}
func hashPassword(password string, salt []byte) string {
	var passwordBytes = []byte(password)
	var sha512Hasher = sha512.New()
	passwordBytes = append(passwordBytes, salt...)
	sha512Hasher.Write(passwordBytes)
	var hashedPasswordBytes = sha512Hasher.Sum(nil)
	var base64EncodedPasswordHash = base64.URLEncoding.EncodeToString(hashedPasswordBytes)
	return base64EncodedPasswordHash

}

func Mhash(passw string) string {

	var HashedPassword = hashPassword(passw, generateRandomSalt(16))
	return HashedPassword

}

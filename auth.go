package ambition

import (
	crand "crypto/rand"
	"encoding/base64"
	"golang.org/x/crypto/bcrypt"
	"math/rand"
	"strconv"
	"time"
)

// CreateSaltAndHashedPassword takes a password as a byte array
// and returns a hashed version of the password and a password salt
// that was generated and used in the hashing process
func CreateSaltAndHashedPassword(password []byte) ([]byte, []byte, error) {
	// Create password salt from the current time and a random integer
	// Salts just need to be unique in case of a database breach.
	// Current nanosecond + random 31bit int have a very small chance
	// of ever having a collision
	nano := time.Now().UnixNano()
	rand.Seed(nano)
	random := rand.Int31()
	salt := strconv.Itoa(int(nano)) + strconv.Itoa(int(random))

	// Create password hash using bcrypt
	hash, err := bcrypt.GenerateFromPassword([]byte(string(password)+salt), 10)

	return []byte(salt), hash, err
}

// CompareSaltAndPasswordToHash takes a salt, a hashedPassword, and a submitted password and
// returns true if the submitted password was the original password saved
func CompareSaltAndPasswordToHash(salt, hashedPassword, password []byte) bool {
	// Combine sumbited password and salt
	saltedPassword := []byte(string(password) + string(salt))

	// Use bcrypt to compare hashedPassword and saltedPassword
	err := bcrypt.CompareHashAndPassword(hashedPassword, saltedPassword)
	return err == nil
}

// GenerateRandomBytes returns securely generated random bytes.
// It will return an error if the system's secure random
// number generator fails to function correctly, in which
// case the caller should not continue.
func GenerateRandomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := crand.Read(b)
	// Note that err == nil only if we read len(b) bytes.
	if err != nil {
		return nil, err
	}

	return b, nil
}

// GenerateRandomString returns a URL-safe, base64 encoded
// securely generated random string.
// It will return an error if the system's secure random
// number generator fails to function correctly, in which
// case the caller should not continue.
func GenerateRandomString(s int) (string, error) {
	b, err := GenerateRandomBytes(s)
	return base64.URLEncoding.EncodeToString(b), err
}

func HashToken(token string) string {
	hash, _ := bcrypt.GenerateFromPassword([]byte(token), 10)
	return string(hash)
}

func CompareHashAndToken(hash string, token string) bool {

	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(token))
	return err == nil
}

func LoginUser(username string, password string) {}

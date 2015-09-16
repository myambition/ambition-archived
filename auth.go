package ambition

import (
	crand "crypto/rand"
	"encoding/base64"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"math/rand"
	"strconv"
	"time"
)

func CreateSaltAndHashedPassword(password []byte) (passwordSalt []byte, hashedPassword []byte) {
	nano := time.Now().UnixNano()
	rand.Seed(nano)
	random := rand.Int31()
	salt := strconv.Itoa(int(nano)) + strconv.Itoa(int(random))
	fmt.Println(string(salt))
	hash, _ := bcrypt.GenerateFromPassword([]byte(string(password)+salt), 10)
	fmt.Println(string(password) + salt)
	fmt.Println(hash)
	return []byte(salt), hash
}

func CompareSaltAndPasswordToHash(salt, password, hashedPassword []byte) bool {
	saltedPassword := []byte(string(password) + string(salt))
	fmt.Println(string(saltedPassword))
	fmt.Println("Why are the two hashes different but still pass?")
	fmt.Println(string(hashedPassword))
	hash, _ := bcrypt.GenerateFromPassword(saltedPassword, 10)
	fmt.Println(string(hash))
	err := bcrypt.CompareHashAndPassword(hashedPassword, saltedPassword)
	check(err)
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
	check(err)
	return err == nil
}

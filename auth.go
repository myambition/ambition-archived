package ambition

import (
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

package intermediate

import (
	"crypto/rand"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/base64"
	"fmt"
	"io"
)


func main() {

	password := "password123"

	hash := sha256.Sum256([]byte(password))
	hash512 := sha512.Sum512([]byte(password))
	
	fmt.Println(password)
	fmt.Println(hash)
	fmt.Println(hash512)
	fmt.Printf("sha-256 hash hex val: %x\n", hash)
	fmt.Printf("sha-256 hash hex val: %x\n", hash512)

	salt, err := generateSalt()
	fmt.Println("Original salt:", salt)
	fmt.Printf("Original salt: %x\n", salt)
	if err != nil {
		fmt.Println("Error generating salt:", err)
		return
	}

	// hash the password with salt
	signUpHash := hashPassword(password, salt)

	// Store the salt and password in database, just printing as of now
	saltStr := base64.StdEncoding.EncodeToString(salt)
	fmt.Println("salt:", saltStr) // simulate as storing in databse
	fmt.Println("hash:",signUpHash) // simulate as storing in databse 
	hashOriginalPassword := sha256.Sum256([]byte(password))
	fmt.Println("hash of just the password string whitout salt:", base64.StdEncoding.EncodeToString(hashOriginalPassword[:]))

	// verify
	// retrieve the saltStr and decode it
	decodedSalt, err := base64.StdEncoding.DecodeString(saltStr)
	if err != nil {
		fmt.Println("Error decoding salt:", err)
		return
	}
	loginHash := hashPassword(password, decodedSalt)

	// compared the stored signUpHash with loginHash
	if signUpHash == loginHash {
		fmt.Println("Password verified successfully!")
	} else {
		fmt.Println("Login failed. Please check user credentials.")
	}
}

func generateSalt() ([]byte, error) {
	salt := make([]byte, 16)
	_, err := io.ReadFull(rand.Reader, salt)
	if err != nil {
		return nil, err
	}
	return salt, nil
}

func hashPassword(password string, salt []byte) string {
	saltedPassword := append(salt, []byte(password)...)
	hash := sha256.Sum256(saltedPassword)
	return base64.StdEncoding.EncodeToString(hash[:])
}
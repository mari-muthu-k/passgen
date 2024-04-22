package hashpass

import (
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"
)

// HashPassword is used to encrypt the password
func HashPassword(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		log.Panic(err)
	}

	return string(bytes)
}

func ComparePassword(userPassword string, providedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(providedPassword), []byte(userPassword))
	fmt.Println(err, len([]byte(providedPassword)))
	return err == nil
}

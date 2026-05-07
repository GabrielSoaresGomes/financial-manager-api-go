package crypto

import "golang.org/x/crypto/bcrypt"

func CashPassword(password string) (string, error) {
    hashedPassword, generateHashError := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    if generateHashError != nil {
        return "", generateHashError
    }
    return string(hashedPassword), nil
}

func CompareHashAndPassword(hashedPassword string, comparingPassword string) bool {
    hashedPasswordByte := []byte(hashedPassword)
    comparingPasswordByte := []byte(comparingPassword)

    comparePasswordError := bcrypt.CompareHashAndPassword(hashedPasswordByte, comparingPasswordByte)
    return comparePasswordError == nil // Não seu erro então a senha é igual
}

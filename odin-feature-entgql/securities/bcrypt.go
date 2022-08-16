package securities

import "golang.org/x/crypto/bcrypt"

func HashPassword(s string) string {
	hashed, _ := bcrypt.GenerateFromPassword([]byte(s), 10)
	return string(hashed)
}

func ComparePassword(hashed, normal string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashed), []byte(normal))
}

package auth

import (
	"math/rand"
	"time"

	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

func hashAndSalt(password string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		logrus.Info(err)
	}
	return string(hash)
}

func comparePasswords(hashedPwd string, plainPwd string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPwd), []byte(plainPwd))
	return err == nil
}
func GeneratePassword() string {

	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	const length = 12

	seededRand := rand.New(rand.NewSource(int64(time.Now().UnixNano())))

	b := make([]byte, length)

	for i := range b {

		b[i] = charset[seededRand.Intn(len(charset))]

	}

	return string(b)

}

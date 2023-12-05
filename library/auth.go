package library

import (
	"crypto/sha256"
	"encoding/hex"
	"go-project/model"
	"io"
	"os"

	"github.com/golang-jwt/jwt"
)

type claims struct {
	Email       string
	DisplayName string
	jwt.StandardClaims
}

func GetJwtKey() []byte {
	file, err := os.Open("crypt/key.pem")
	if err != nil {
		return nil
	}
	defer file.Close()
	fileByte, _ := io.ReadAll(file)
	return fileByte
}

func CreateJwt(model model.User) (string, error) {
	//expirationTime := time.Now().Add(15 * time.Minute)
	claims := &claims{
		Email:          model.UserName,
		DisplayName:    model.DisplayName,
		StandardClaims: jwt.StandardClaims{
			//ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(GetJwtKey())
}

func HashString(s string) string {
	h := sha256.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

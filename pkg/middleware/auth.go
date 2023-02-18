package custommiddleware

import (
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"log"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"
)

var SecretKey []byte

const (
	SECRET  string = "SECRET"
	NO_AUTH string = "unauthorized access"
)

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		setSecret()
		if len(SecretKey) <= 0 {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Missing Secret"))
			return
		}

		authHeader := strings.Split(r.Header.Get("Authorization"), "Bearer ")
		if len(authHeader) != 2 {
			log.Println("malformed token")
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte(NO_AUTH))
			return
		}

		jwtToken := authHeader[1]
		token, err := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
			_, ok := token.Method.(*jwt.SigningMethodHMAC)
			if !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return SecretKey, nil
		})
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte(NO_AUTH))
			return
		}

		if claims, ok := token.Claims.(jwt.Claims); ok && token.Valid {
			fmt.Println(claims)
		} else {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte(NO_AUTH))
			return
		}

		next.ServeHTTP(w, r)
	})
}

func GenerateToken(jwtPayload JWTPayload) (string, error) {
	setSecret()

	token := jwt.New(jwt.SigningMethodHS256)
	/* Create a map to store our claims */
	claims := token.Claims.(jwt.MapClaims)

	/* Set token claims */
	claims["id"] = jwtPayload.ID
	claims["role"] = jwtPayload.Role
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	claims["iat"] = time.Now().Unix()

	tokenString, err := token.SignedString(SecretKey)
	if err != nil {
		log.Fatal("Error in Generating key")
		return "", err
	}

	return tokenString, nil
}

// Set Secret If Secret is Not Initialized
func setSecret() {
	var m sync.Mutex
	if len(SecretKey) <= 0 {
		m.Lock()
		defer m.Unlock()
		SecretKey = []byte(os.Getenv(SECRET))
	}
}

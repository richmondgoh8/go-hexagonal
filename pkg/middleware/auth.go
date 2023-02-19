package custommiddleware

import (
	"context"
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

type Claims struct {
	ID   int    `json:"id"`
	Role string `json:"role"`
	jwt.RegisteredClaims
}

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
		token, err := jwt.ParseWithClaims(jwtToken, &Claims{}, func(token *jwt.Token) (interface{}, error) {
			// verify the signing method
			_, ok := token.Method.(*jwt.SigningMethodHMAC)
			if !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			// return secret key for validating the token
			return SecretKey, nil
		})
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte(NO_AUTH))
			return
		}

		claims, ok := token.Claims.(*Claims)
		if !ok || !token.Valid {
			log.Println("Token is invalid")
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte(NO_AUTH))
			return
		}

		// set to context map
		ctx := context.WithValue(r.Context(), "jwt_claims", claims)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func GenerateToken(jwtPayload JWTPayload) (string, error) {
	setSecret()

	// Create a map to set & store our token claims
	claims := jwt.MapClaims{
		"id":   jwtPayload.ID,
		"role": jwtPayload.Role,
		"iat":  time.Now().Unix(),
		"exp":  time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

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

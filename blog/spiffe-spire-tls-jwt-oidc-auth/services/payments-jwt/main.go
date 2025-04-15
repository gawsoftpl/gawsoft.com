package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/MicahParks/keyfunc/v2"
)

var (
	jwksURL   string
	audience  string
	sub    string
	jwks      *keyfunc.JWKS
)

func main() {
	log.Println("Starting...")

	// ENV variables
	jwksURL = os.Getenv("JWKS_URI")
	audience = os.Getenv("JWT_AUDIENCE")
	sub = os.Getenv("JWT_SUB")


	if jwksURL == "" || audience == "" || sub == "" {
		log.Fatal("JWKS_URI, JWT_AUDIENCE, and JWT_SUB must be set")
	}

	// Create JWKS client
	var err error
	jwks, err = keyfunc.Get(jwksURL, keyfunc.Options{
		RefreshInterval: time.Minute * 5,
		RefreshTimeout:  time.Second * 5,
		RefreshErrorHandler: func(err error) {
			log.Printf("Error refreshing JWKS: %v", err)
		},
	})
	if err != nil {
		log.Fatalf("Error getting JWKS: %v", err)
	}

	// Route
	http.HandleFunc("/", jwtHandler)

	log.Println("✅ Server running on :3000")
	log.Fatal(http.ListenAndServe(":3000", nil))
}

func jwtHandler(w http.ResponseWriter, r *http.Request) {
	authHeader := r.Header.Get("Authorization")
	if !strings.HasPrefix(authHeader, "Bearer ") {
		http.Error(w, "Missing or invalid Authorization header", http.StatusUnauthorized)
		return
	}

	tokenStr := strings.TrimPrefix(authHeader, "Bearer ")

	// Parse and validate token
	token, err := jwt.Parse(tokenStr, jwks.Keyfunc,
		jwt.WithAudience(audience),
		jwt.WithSubject(sub),
	)
	if err != nil || !token.Valid {
		http.Error(w, "Unauthorized: "+err.Error(), http.StatusUnauthorized)
		return
	}

	claims := token.Claims.(jwt.MapClaims)
	fmt.Fprintf(w, "🎉 Token is valid!\n\nClaims:\n")
	for k, v := range claims {
		fmt.Fprintf(w, "%s: %v\n", k, v)
	}
}

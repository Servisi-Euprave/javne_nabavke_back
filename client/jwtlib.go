package client

import (
	"crypto/rsa"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func CheckAuthWithPublicKey(publicKey *rsa.PublicKey) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		bearerToken := ctx.GetHeader("Authorization")
		if bearerToken == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"errors": "Authorization header not provided"})
			return
		}
		bearer, tokenStr, found := strings.Cut(bearerToken, " ")
		if !found || bearer != "Bearer" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"errors": "invalid format for bearer token"})
			return
		}
		token, err := jwt.ParseWithClaims(tokenStr, &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
				log.Println("Invalid signing method")
				return nil, errors.New("unexpected signing method")
			}
			return publicKey, nil
		})
		if err != nil || !token.Valid {
			log.Printf("Public key: %v\n", publicKey)
			log.Printf("errors parsing token: %s", err.Error())
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"errors": "invalid token"})
			return
		}
		// Token is valid, extract the claims and set them in the context
		claims, ok := token.Claims.(*jwt.RegisteredClaims)
		if !ok {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"errors": "invalid claims"})
			return
		}
		log.Println(claims.Subject)
		ctx.Set("claims", claims.Subject)
		ctx.Next()
	}
}
func ReadRSAPublicKeyFromFile(filePath string) (*rsa.PublicKey, error) {
	var key *rsa.PublicKey
	keyFile, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("No key provided, errors generating key: %w", err)
	}
	defer keyFile.Close()
	keyData, err := ioutil.ReadAll(keyFile)
	if err != nil {
		return nil, fmt.Errorf("Error reading from key file: %w", err)
	}
	key, err = jwt.ParseRSAPublicKeyFromPEM(keyData)
	if err != nil {
		return key, fmt.Errorf("Error parsing PMA encoded key: %w", err)
	}
	return key, nil
}

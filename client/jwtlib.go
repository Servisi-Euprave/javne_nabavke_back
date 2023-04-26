package client

import (
	"crypto/rsa"
	"errors"
	"fmt"
	"io/ioutil"
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
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization header not provided"})
			return
		}
		bearer, tokenStr, found := strings.Cut(bearerToken, " ")
		if !found || bearer != "Bearer" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid format for bearer token"})
			return
		}
		token, err := jwt.ParseWithClaims(tokenStr, &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
				return nil, errors.New("unexpected signing method")
			}
			return publicKey, nil
		})
		if err != nil || !token.Valid {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
			return
		}
		// Token is valid, extract the claims and set them in the context
		claims, ok := token.Claims.(*jwt.RegisteredClaims)
		if !ok {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid claims"})
			return
		}
		ctx.Set("claims", claims)
		ctx.Next()
	}

}
func ReadRSAPublicKeyFromFile(filePath string) (*rsa.PublicKey, error) {
	var key *rsa.PublicKey
	keyFile, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("No key provided, error generating key: %w", err)
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
